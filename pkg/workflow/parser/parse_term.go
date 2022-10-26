package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var MissingLiteral = err.Trackable("Missing literal")
var MissingIdentifier = err.Trackable("Missing identifier")

// TERM := VAR | LITERAL
func parseTerm(a auditor) ast.Expr {
	if !a.More() {
		panic(a.unexpectedEOF("operand"))
	}

	if expr := acceptTerm(a); expr != nil {
		return expr
	}

	panic(a.unexpectedToken("operand", a.Peek()))
}

func acceptTerm(a auditor) ast.Expr {
	switch {
	case !a.More():
		return nil
	case a.is(token.Identifier):
		return parseVariable(a)
	case a.match(token.IsLiteral):
		return expectLiteral(a)
	default:
		return nil
	}
}

// VARS := VAR { Comma VAR }
func expectVariables(a auditor) []ast.Variable {
	var nodes []ast.Variable

	v := parseVariable(a)
	nodes = append(nodes, v)

	for a.accept(token.Comma) {
		v := parseVariable(a)
		nodes = append(nodes, v)
	}

	return nodes
}

// VAR := Identifier
func parseVariable(a auditor) ast.Variable {
	n, e := a.expect_new(token.Identifier)

	if e != nil {
		panic(MissingIdentifier.Wrapf(
			e, "Expected identifier but got %s", a.Peek().String(),
		))
	}

	return ast.Variable{
		Identifier: n,
	}
}

// LITERAL := True | False | Number | String
func expectLiteral(a auditor) ast.Expr {
	if n := acceptLiteral(a); n != nil {
		return n
	}
	panic(MissingLiteral.Trackf("Expected literal but got %s", a.Peek().String()))
}

func acceptLiteral(a auditor) ast.Expr {

	switch a.Peek() {
	case token.True:
	case token.False:
	case token.Number:
	case token.String:
	default:
		return nil
	}

	return ast.Literal{
		Token: a.Read(),
	}
}
