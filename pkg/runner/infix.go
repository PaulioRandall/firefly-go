package runner

import (
	"github.com/PaulioRandall/firefly-go/pkg/ast"
)

func computeInfix(n ast.Node, compute infixComputer) (ast.NumberNode, error) {

	ien, ok := n.(ast.InfixNode)
	if !ok {
		return zero, newBug("ast.InfixNode node expected")
	}

	left, right, e := computeInfixExpr(ien)
	if e != nil {
		return zero, e
	}

	return compute(left, right)
}

func computeInfixExpr(n ast.InfixNode) (left, right ast.NumberNode, e error) {

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

type infixComputer func(left, right ast.NumberNode) (ast.NumberNode, error)

func addNumbers(left, right ast.NumberNode) (ast.NumberNode, error) {
	return newNumber(left.Value + right.Value), nil
}

func subNumbers(left, right ast.NumberNode) (ast.NumberNode, error) {
	return newNumber(left.Value - right.Value), nil
}

func mulNumbers(left, right ast.NumberNode) (ast.NumberNode, error) {
	return newNumber(left.Value * right.Value), nil
}

func divNumbers(left, right ast.NumberNode) (ast.NumberNode, error) {
	if right.Value == 0 {
		return zero, newError("Can't divide by zero")
	}
	return newNumber(left.Value / right.Value), nil
}
