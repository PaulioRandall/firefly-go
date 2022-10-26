package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func acceptExprsUntil(a auditor, closer token.TokenType) []ast.Expr {
	var nodes []ast.Expr

	for a.isNot(closer) {
		v := acceptExpression(a)
		if v == nil {
			break
		}

		nodes = append(nodes, v)

		if !a.accept(token.Comma) {
			break
		}
	}

	return nodes
}

func acceptExpression(a auditor) ast.Expr {
	switch {
	case a.is(token.ParenOpen):
		n := expectParenExpr(a)
		return operation(a, n, 0)
	case a.is(token.BraceOpen):
		return expectMap(a)
	}

	if term, ok := acceptTerm(a); ok {
		return operation(a, term, 0)
	}

	return nil
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

// EXPR := PAREN_EXPR | TERM | LIST | MAP  | OPERATION
func expectExpression(a auditor) ast.Expr {
	switch {
	case a.is(token.ParenOpen):
		return expectParenExpr(a)
	case a.is(token.BraceOpen):
		return expectMap(a)
	}

	left := expectOperand(a)
	return operation(a, left, 0)
}

// PAREN_EXPR := ParenOpen EXPR ParenClose
func expectParenExpr(a auditor) ast.Expr {
	a.expect(token.ParenOpen)
	n := expectExpression(a)
	a.expect(token.ParenClose)
	return n
}
