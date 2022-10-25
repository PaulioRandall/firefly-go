package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var MissingIdentifier = err.Trackable("Missing identifier")

// VARIABLES := Identifier { Comma Identifier }
// VARIABLE  := Identifier

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
	n, e := a.expect_new(token.Identifier)

	if e != nil {
		panic(MissingIdentifier.Wrapf(
			e, "Expected identifier but got %s", a.Peek().String(),
		))
	}

	return ast.Variable{
		Identifier: n,
	}
}
