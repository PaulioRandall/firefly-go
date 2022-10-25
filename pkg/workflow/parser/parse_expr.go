package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// EXPR := TERM | OPERATION
// TERM := VARIABLE | LITERAL

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

func acceptExpressions(a auditor) []ast.Expr {
	var nodes []ast.Expr

	for a.More() {
		v := acceptExpression(a)
		if v == nil {
			break
		}

		nodes = append(nodes, v)
	}

	return nodes
}

func acceptExpression(a auditor) ast.Expr {
	switch {
	case a.is(token.ParenOpen):
		n := parseParenExpr(a)
		return operation(a, n, 0)
	case a.is(token.BracketOpen):
		return parseList(a)
	case a.is(token.BraceOpen):
		return parseMap(a)
	}

	if n := acceptOperand(a); n != nil {
		return operation(a, n, 0)
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

func expectExpression(a auditor) ast.Expr {
	switch {
	case a.is(token.ParenOpen):
		return parseParenExpr(a)
	case a.is(token.BracketOpen):
		return parseList(a)
	case a.is(token.BraceOpen):
		return parseMap(a)
	}

	left := expectOperand(a)
	return operation(a, left, 0)
}

func parseParenExpr(a auditor) ast.Expr {
	a.expect(token.ParenOpen)
	n := expectExpression(a)
	a.expect(token.ParenClose)
	return n
}
