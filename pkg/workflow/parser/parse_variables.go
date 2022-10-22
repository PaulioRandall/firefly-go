package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func expectVariables(a auditor) []ast.Variable {
	var nodes []ast.Variable

	v := expectVariable(a)
	nodes = append(nodes, v)

	for a.accept(token.Comma) {
		v := expectVariable(a)
		nodes = append(nodes, v)
	}

	return nodes
}

func expectVariable(a auditor) ast.Variable {
	return ast.Variable{
		Identifier: a.expect(token.Identifier),
	}
}
