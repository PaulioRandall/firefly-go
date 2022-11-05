package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	ErrMissingExpr       = err.Trackable("Missing expression")
	ErrMissingParenClose = err.Trackable("Missing closing parenthesis")

	ErrBadExpr      = err.Trackable("Failed to parse expression")
	ErrBadParenExpr = err.Trackable("Failed to parse parenthesized expression")
)

func acceptExprsUntil(a auditor, closer token.TokenType) []ast.Expr {
	var nodes []ast.Expr

	for a.isNot(closer) {
		v, ok := acceptExpression(a)
		if !ok {
			break
		}

		nodes = append(nodes, v)

		if !a.accept(token.Comma) {
			break
		}
	}

	return nodes
}

// EXPRS := [EXPR {Comma EXPR}]
func parseSeriesOfExpr(a auditor) ast.SeriesOfExpr {
	var nodes []ast.Expr

	for more := true; more; more = a.accept(token.Comma) {
		n, ok := acceptExpression(a)

		if !ok {
			break
		}

		nodes = append(nodes, n)
	}

	return ast.SeriesOfExpr{
		Nodes: nodes,
	}
}

func expectExpressions(a auditor) []ast.Expr {
	var nodes []ast.Expr

	v := expectExpression(a)
	nodes = append(nodes, v)

	for a.accept(token.Comma) {
		v := expectExpression(a)
		nodes = append(nodes, v)
	}

	return nodes
}

// EXPR := PAREN_EXPR | TERM | OPERATION
func acceptExpression(a auditor) (ast.Expr, bool) {
	defer wrapPanic(func(e error) error {
		return ErrBadExpr.Wrap(e, "Bad expression syntax")
	})

	if left, ok := acceptParenExpr(a); ok {
		return operation(a, left, 0), true
	}

	if left, ok := acceptTerm(a); ok {
		return operation(a, left, 0), true
	}

	return nil, false
}

func expectExpression(a auditor) ast.Expr {
	defer wrapPanic(func(e error) error {
		return ErrBadExpr.Wrap(e, "Bad expression syntax")
	})

	if left, ok := acceptExpression(a); ok {
		return operation(a, left, 0)
	}

	panic(unableToParse(a, ErrMissingExpr, "any in [PAREN_EXPR | TERM | LIST | MAP | OPERATION]"))
}

// PAREN_EXPR := ParenOpen EXPR ParenClose
func acceptParenExpr(a auditor) (ast.Expr, bool) {
	defer wrapPanic(func(e error) error {
		return ErrBadParenExpr.Wrap(e, "Bad parenthesized expression")
	})

	if a.isNot(token.ParenOpen) {
		return nil, false
	}

	a.Read()
	n := expectExpression(a)

	if a.isNot(token.ParenClose) {
		panic(ErrMissingParenClose.Track("Parenthesized expression was left open"))
	}

	a.Read()
	return n, true
}

func isNextExprOpener(a auditor) bool {
	return a.isAny(token.ParenOpen, token.BracketOpen, token.BraceOpen)
}

func isNextLiteral(a auditor) bool {
	return a.isAny(token.Number, token.String, token.True, token.False)
}
