package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// ASSIGN := VARS Assign EXPRS
func expectAssignment(a auditor) ast.Assign {
	return ast.Assign{
		Left:     acceptVariables(a),
		Operator: a.expect(token.Assign),
		Right: ast.SeriesOfExpr{
			Exprs: acceptExpressions(a),
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
