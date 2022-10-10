package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func expectAssignment(a *auditor) ast.Assign {
	n := ast.Assign{}

	n.Left = expectVariables(a)
	n.Token = a.expect(token.Assign)
	n.Right = expectExpressions(a)

	// TODO: Move specific parameter checks to the validator
	if len(n.Left) > len(n.Right) {
		panic(MissingExpr)
	} else if len(n.Left) < len(n.Right) {
		panic(MissingVar)
	}

	return n
}
