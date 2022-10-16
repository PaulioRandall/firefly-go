package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func expectAssignment(a tokenAuditor) ast.Assign {
	return ast.Assign{
		Left:     expectVariables(a),
		Operator: expect(a, token.Assign),
		Right: ast.ExprSet{
			Exprs: expectExpressions(a),
		},
	}

	/*
		// TODO: Move specific parameter checks to the validator
		if len(n.Left) > len(n.Right) {
			panic(MissingExpr)
		} else if len(n.Left) < len(n.Right) {
			panic(MissingVar)
		}
	*/
}
