package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func expectVariables(a tokenAuditor) []ast.Variable {
	var nodes []ast.Variable

	v := expectVariable(a)
	nodes = append(nodes, v)

	for accept(a, token.Comma) {
		v := expectVariable(a)
		nodes = append(nodes, v)
	}

	return nodes
}

func expectVariable(a tokenAuditor) ast.Variable {
	return ast.Variable{
		Identifier: expect(a, token.Identifier),
	}
}
