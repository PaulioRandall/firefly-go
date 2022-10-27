package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func isAssignment(a auditor) bool {
	if !a.accept(token.Identifier) {
		return false
	}

	ident := a.Prev()
	is := a.isAny(token.Comma, token.Assign)

	a.Putback(ident)
	return is
}

// ASSIGNMENT := VARIABLES Assign EXPRESSIONS
func acceptAssignment(a auditor) (ast.Assign, bool) {
	if !isAssignment(a) {
		return ast.Assign{}, false
	}

	n := ast.Assign{
		Left:     parseSeriesOfVar(a),
		Operator: a.expect(token.Assign),
		Right:    parseSeriesOfExpr(a),
	}

	return n, true
}

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
