package runner

import (
	"github.com/PaulioRandall/firefly-go/firefly/ast"
)

type infixComputer func(left, right ast.NumberTree) (ast.NumberTree, error)

func computeInfix(tr ast.Tree, compute infixComputer) (ast.NumberTree, error) {

	ien, ok := tr.(ast.InfixTree)
	if !ok {
		return zero, newBug("ast.InfixTree node expected")
	}

	left, right, e := computeInfixExpr(ien)
	if e != nil {
		return zero, e
	}

	return compute(left, right)
}

func computeInfixExpr(tr ast.InfixTree) (left, right ast.NumberTree, e error) {

	left, e = computeTree(tr.Left)
	if e != nil {
		return zero, zero, e
	}

	right, e = computeTree(tr.Right)
	if e != nil {
		return zero, zero, e
	}

	return left, right, nil
}
