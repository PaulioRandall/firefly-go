package runner

import (
	"github.com/PaulioRandall/firefly-go/pkg/ast"
)

func computeInfix(n ast.Node, compute infixComputer) (ast.NumberNode, error) {

	ien, ok := n.(ast.InfixExprNode)
	if !ok {
		return zero, newBug("ast.InfixExprNode node expected")
	}

	left, right, e := computeInfixExpr(ien)
	if e != nil {
		return zero, e
	}

	result := compute(left, right)
	return result, nil
}

func computeInfixExpr(n ast.InfixExprNode) (left, right ast.NumberNode, e error) {

	left, e = computeNode(n.Left)
	if e != nil {
		return zero, zero, e
	}

	right, e = computeNode(n.Right)
	if e != nil {
		return zero, zero, e
	}

	return left, right, nil
}

type infixComputer func(left, right ast.NumberNode) ast.NumberNode

func addNumbers(left, right ast.NumberNode) ast.NumberNode {
	return newNumber(left.Value + right.Value)
}

func subNumbers(left, right ast.NumberNode) ast.NumberNode {
	return newNumber(left.Value - right.Value)
}

func mulNumbers(left, right ast.NumberNode) ast.NumberNode {
	return newNumber(left.Value * right.Value)
}

func divNumbers(left, right ast.NumberNode) ast.NumberNode {
	return newNumber(left.Value / right.Value)
}
