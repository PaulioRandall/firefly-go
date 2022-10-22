package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

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
	if expr := acceptOperand(a); expr != nil {
		return operation(a, expr, 0)
	}
	return nil
}

func acceptOperand(a auditor) ast.Expr {
	switch {
	case !a.More():
		return nil
	case a.match(token.IsLiteral):
		return expectLiteral(a)
	default:
		return nil
	}
}

func acceptLiteral(a auditor) ast.Expr {
	tk, ok := a.acquireIf(token.IsLiteral)
	if !ok {
		return nil
	}

	return ast.Literal{
		Token: tk,
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

func expectLiteral(a auditor) ast.Expr {
	return ast.Literal{
		Token: a.expectFor("literal", token.IsLiteral),
	}
}

func operation(a auditor, left ast.Expr, leftOperatorPriorty int) ast.Expr {
	if !a.notMatch(token.IsBinaryOperator) {
		return left
	}

	if leftOperatorPriorty >= a.Peek().Precedence() {
		return left
	}

	op := a.Next()

	right := expectOperand(a)
	right = operation(a, right, op.Precedence())

	left = ast.BinaryOperation{
		Left:     left,
		Operator: op,
		Right:    right,
	}

	return operation(a, left, leftOperatorPriorty)
}
