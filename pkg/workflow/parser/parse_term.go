package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var MissingTerm = err.Trackable("Missing term")
var MissingLiteral = err.Trackable("Missing literal")
var MissingIdentifier = err.Trackable("Missing identifier")

// TERM := VAR | LITERAL | LIST | MAP
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

func expectTerm(a auditor) (ast.Expr, error) {
	if n, ok := acceptTerm(a); ok {
		return n, nil
	}

	return nil, unableToParse(a, MissingTerm, "term")
}

// VARIABLES := [VARIABLE {Comma VARIABLE}]
func acceptVariables(a auditor) []ast.Variable {
	var nodes []ast.Variable

	for more := true; more; more = a.accept(token.Comma) {
		n, ok := acceptVariable(a)

		if !ok {
			break
		}

		nodes = append(nodes, n)
	}

	return nodes
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

// LITERAL := True | False | Number | String
func acceptLiteral(a auditor) (ast.Expr, bool) {
	if a.isNotAny(token.True, token.False, token.Number, token.String) {
		return nil, false
	}

	n := ast.Literal{
		Token: a.Read(),
	}

	return n, true
}
