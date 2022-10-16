package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func expectVariables(r PosReaderOfTokens) []ast.Variable {
	var nodes []ast.Variable

	v := expectVariable(r)
	nodes = append(nodes, v)

	for accept(r, token.Comma) {
		v := expectVariable(r)
		nodes = append(nodes, v)
	}

	return nodes
}

func expectVariable(r PosReaderOfTokens) ast.Variable {
	return ast.Variable{
		Identifier: expect(r, token.Identifier),
	}
}
