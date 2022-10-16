package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func expectAssignment(r BufReaderOfTokens) ast.Assign {
	return ast.Assign{
		Left:     expectVariables(r),
		Operator: expect(r, token.Assign),
		Right: ast.ExprSet{
			Exprs: expectExpressions(r),
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
