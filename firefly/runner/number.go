package runner

import (
	"github.com/PaulioRandall/firefly-go/firefly/ast"
)

var zero ast.NumberTree

func computeNumber(tr ast.Tree) (ast.NumberTree, error) {
	num, ok := tr.(ast.NumberTree)
	if !ok {
		return ast.NumberTree{}, newBug("ast.NumberTree node expected")
	}
	return num, nil
}

func newNumber(n int64) ast.NumberTree {
	return ast.NumberTree{
		Value: n,
	}
}
