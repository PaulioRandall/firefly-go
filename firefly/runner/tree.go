package runner

import (
	"github.com/PaulioRandall/firefly-go/firefly/ast"
)

func computeTree(tr ast.Tree) (ast.NumberTree, error) {

	var result ast.NumberTree
	var e error

	switch tr.Type() {
	case ast.NODE_NUM:
		result, e = computeNumber(tr)

	case ast.NODE_ADD:
		result, e = computeInfix(tr, addNumbers)

	case ast.NODE_SUB:
		result, e = computeInfix(tr, subNumbers)

	case ast.NODE_MUL:
		result, e = computeInfix(tr, mulNumbers)

	case ast.NODE_DIV:
		result, e = computeInfix(tr, divNumbers)

	default:
		e = newBug("Unknown AST node")
	}

	if e != nil {
		return zero, e
	}
	return result, nil
}

func addNumbers(left, right ast.NumberTree) (ast.NumberTree, error) {
	return newNumber(left.Value + right.Value), nil
}

func subNumbers(left, right ast.NumberTree) (ast.NumberTree, error) {
	return newNumber(left.Value - right.Value), nil
}

func mulNumbers(left, right ast.NumberTree) (ast.NumberTree, error) {
	return newNumber(left.Value * right.Value), nil
}

func divNumbers(left, right ast.NumberTree) (ast.NumberTree, error) {
	if right.Value == 0 {
		return zero, newError("Can't divide by zero")
	}
	return newNumber(left.Value / right.Value), nil
}
