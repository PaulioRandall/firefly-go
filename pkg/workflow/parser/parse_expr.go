package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var MissingExpr = err.Trackable("Missing expression")

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
func acceptExpressions(a auditor) []ast.Expr {
	var nodes []ast.Expr

	for more := true; more; more = a.accept(token.Comma) {
		n, ok := acceptExpression(a)

		if !ok {
			break
		}

		nodes = append(nodes, n)
	}

	return nodes
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
	if left, ok := acceptParenExpr(a); ok {
		return operation(a, left, 0), true
	}

	if left, ok := acceptTerm(a); ok {
		return operation(a, left, 0), true
	}

	return nil, false
}

func expectExpression(a auditor) ast.Expr {
	if left, ok := acceptExpression(a); ok {
		return operation(a, left, 0)
	}

	panic(unableToParse(a, MissingExpr, "any in [PAREN_EXPR | TERM | LIST | MAP | OPERATION]"))
}

// PAREN_EXPR := ParenOpen EXPR ParenClose
func acceptParenExpr(a auditor) (ast.Expr, bool) {
	if a.isNot(token.ParenOpen) {
		return nil, false
	}

	a.Read()
	n := expectExpression(a)
	a.expect(token.ParenClose)

	return n, true
}
