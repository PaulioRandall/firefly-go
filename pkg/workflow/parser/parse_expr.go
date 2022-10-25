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

func acceptOperand(a auditor) ast.Expr {
	switch {
	case !a.More():
		return nil
	case a.is(token.Identifier):
		return expectIdentifier(a)
	case a.match(token.IsLiteral):
		return expectLiteral(a)
	default:
		return nil
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

func expectOperand(a auditor) ast.Expr {
	if !a.More() {
		panic(a.unexpectedEOF("operand"))
	}

	if expr := acceptOperand(a); expr != nil {
		return expr
	}

	panic(a.unexpected("operand", a.Peek()))
}

func operation(a auditor, left ast.Expr, leftOperatorPriorty int) ast.Expr {
	if !a.notMatch(token.IsBinaryOperator) {
		return left
	}

	if leftOperatorPriorty >= a.Peek().Precedence() {
		return left
	}

	op := a.Read()

	var right ast.Expr
	if a.is(token.ParenOpen) {
		right = parseParenExpr(a)
	} else {
		right = expectOperand(a)
	}

	right = operation(a, right, op.Precedence())

	left = ast.BinaryOperation{
		Left:     left,
		Operator: op,
		Right:    right,
	}

	return operation(a, left, leftOperatorPriorty)
}

func parseParenExpr(a auditor) ast.Expr {
	a.expect(token.ParenOpen)
	n := expectExpression(a)
	a.expect(token.ParenClose)
	return n
}
