package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// ASSIGNMENT := VARIABLES Assign EXPRESSIONS
func expectAssignment(a auditor) ast.Assign {
	return ast.Assign{
		Left:     parseSeriesOfVar(a),
		Operator: a.expect(token.Assign),
		Right:    parseSeriesOfExpr(a),
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
