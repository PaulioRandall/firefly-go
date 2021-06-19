package runner

import (
	"github.com/PaulioRandall/firefly-go/pkg/ast"
)

var zero ast.NumberNode

func computeNumber(n ast.Node) (ast.NumberNode, error) {
	num, ok := n.(ast.NumberNode)
	if !ok {
		return ast.NumberNode{}, newBug("ast.NumberNode node expected")
	}
	return num, nil
}

func newNumber(n int64) ast.NumberNode {
	return ast.NumberNode{
		Value: n,
	}
}
