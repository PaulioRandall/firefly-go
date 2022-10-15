package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/auditor"
)

func expectVariables(a *auditor.Auditor) []ast.Variable {
	var nodes []ast.Variable

	v := expectVariable(a)
	nodes = append(nodes, v)

	for a.Accept(token.Comma) {
		v := expectVariable(a)
		nodes = append(nodes, v)
	}

	return nodes
}

func expectVariable(a *auditor.Auditor) ast.Variable {
	return ast.Variable{
		Identifier: a.Expect(token.Identifier),
	}
}
