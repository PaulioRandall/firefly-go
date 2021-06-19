package runner

import (
	"github.com/PaulioRandall/firefly-go/pkg/ast"
)

func computeNode(n ast.Node) (ast.NumberNode, error) {

	var result ast.NumberNode
	var e error

	switch n.Type() {
	case ast.AstNumber:
		result, e = computeNumber(n)

	case ast.AstAdd:
		result, e = computeInfix(n, addNumbers)

	case ast.AstSub:
		result, e = computeInfix(n, subNumbers)

	case ast.AstMul:
		result, e = computeInfix(n, mulNumbers)

	case ast.AstDiv:
		result, e = computeInfix(n, divNumbers)

	default:
		e = newBug("Unknown AST node")
	}

	if e != nil {
		return zero, e
	}
	return result, nil
}
