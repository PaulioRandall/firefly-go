package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	ErrBadAssign = err.Trackable("Failed to parse assignment")
)

// ASSIGNMENT := VARIABLES Assign EXPRESSIONS
func acceptAssignment(a auditor) (ast.Assign, bool) {
	if !isAssignment(a) {
		return ast.Assign{}, false
	}

	return expectAssignment(a), true
}

func expectAssignment(a auditor) ast.Assign {
	defer wrapPanic(func(e error) error {
		return ErrBadAssign.Wrap(e, "Expected assignment or encountered bad assignment")
	})

	return ast.Assign{
		Left:     parseSeriesOfVar(a),
		Operator: a.expect(token.Assign),
		Right:    parseSeriesOfExpr(a),
	}
}

func isAssignment(a auditor) bool {
	if !a.accept(token.Identifier) {
		return false
	}

	ident := a.Prev()
	is := a.isAny(token.Comma, token.Assign)

	a.Putback(ident)
	return is
}
