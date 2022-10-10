package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func parseVariables(a *auditor) []ast.Variable {
	var nodes []ast.Variable

	v := parseVariable(a)
	nodes = append(nodes, v)

	for a.accept(token.Comma) {
		v := parseVariable(a)
		nodes = append(nodes, v)
	}

	return nodes
}

func parseVariable(a *auditor) ast.Variable {
	return ast.Variable{
		Token: a.expect(token.Var),
	}
}
