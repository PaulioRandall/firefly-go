package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	ErrMissingVariable = err.Trackable("Failed to parse variable")
)

// TERM := VARIABLE | LITERAL | LIST | MAP
func acceptTerm(a auditor) (ast.Expr, bool) {
	if !a.More() {
		return nil, false
	}

	if n, ok := acceptVariable(a); ok {
		return n, true
	}

	if n, ok := acceptLiteral(a); ok {
		return n, true
	}

	if n, ok := acceptList(a); ok {
		return n, true
	}

	if n, ok := acceptMap(a); ok {
		return n, true
	}

	return nil, false
}

// VARIABLES := [VARIABLE {Comma VARIABLE}]
func parseSeriesOfVar(a auditor) ast.SeriesOfVar {
	var nodes []ast.Variable

	for more := true; more; more = a.accept(token.Comma) {
		n, ok := acceptVariable(a)

		if !ok {
			break
		}

		nodes = append(nodes, n)
	}

	return ast.SeriesOfVar{
		Nodes: nodes,
	}
}

// VARIABLE := Identifier
func acceptVariable(a auditor) (ast.Variable, bool) {
	if !a.accept(token.Identifier) {
		return ast.Variable{}, false
	}

	n := ast.Variable{
		Identifier: a.Prev(),
	}

	return n, true
}

func expectVariable(a auditor) ast.Variable {
	if n, ok := acceptVariable(a); ok {
		return n
	}

	panic(unableToParse(a, ErrMissingVariable, "variable"))
}

// LITERAL := True | False | Number | String
func acceptLiteral(a auditor) (ast.Expr, bool) {
	if a.isNotAny(token.Bool, token.Number, token.String) {
		return nil, false
	}

	n := ast.Literal{
		Token: a.Read(),
	}

	return n, true
}
