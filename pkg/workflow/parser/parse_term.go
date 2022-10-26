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
func expectTerm(a auditor) (ast.Expr, error) {
	if term, ok := acceptTerm(a); ok {
		return term, nil
	}

	if !a.More() {
		return nil, MissingTerm.Wrap(
			UnexpectedEOF.Trackf("Expected term but got EOF"),
			"Unable to parse term",
		)
	}

	return nil, MissingTerm.Wrap(
		UnexpectedToken.Trackf("Expected term but got %s", a.Peek().String()),
		"Unable to parse term",
	)
}

func acceptTerm(a auditor) (ast.Expr, bool) {
	if !a.More() {
		return nil, false
	}

	if term, ok := acceptVariable(a); ok {
		return term, true
	}

	if term, ok := acceptLiteral(a); ok {
		return term, true
	}

	if term, ok := acceptList(a); ok {
		return term, true
	}

	if term, ok := acceptMap(a); ok {
		return term, true
	}

	return nil, false
}

// VAR := Identifier
func acceptVariable(a auditor) (ast.Variable, bool) {
	if !a.is(token.Identifier) {
		return ast.Variable{}, false
	}

	n := ast.Variable{
		Identifier: a.Read(),
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

// ***** OLD ******

// VARS := VAR {Comma VAR}
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

// VAR := Identifier
func expectVariable(a auditor) ast.Variable {
	if n, ok := acceptVariable(a); ok {
		return n
	}
	panic(badNextToken(a, "variable", "identifier"))
}

// LITERAL := True | False | Number | String
func expectLiteral(a auditor) ast.Expr {
	if n, ok := acceptLiteral(a); ok {
		return n
	}
	panic(badNextToken(a, "literal", "literal"))
}
