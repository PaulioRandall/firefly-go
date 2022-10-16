package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast/asttest"
	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func Test_parseAssign_1(t *testing.T) {
	// a = 0

	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(token.Identifier, "a"),
		gen(token.Assign, "="),
		gen(token.Number, "0"),
		gen(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		asttest.Assign(
			asttest.Variables(given[0]),
			given[1],
			asttest.ExprSet(asttest.Expressions(given[2])...),
		),
	}

	assert(t, given, exp)
}

func Test_parseAssign_2(t *testing.T) {
	// a, b = 0, 1

	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(token.Identifier, "a"),
		gen(token.Comma, ","),
		gen(token.Identifier, "b"),
		gen(token.Assign, "="),
		gen(token.Number, "0"),
		gen(token.Comma, ","),
		gen(token.Number, "1"),
		gen(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		asttest.Assign(
			asttest.Variables(given[0], given[2]),
			given[3],
			asttest.ExprSet(asttest.Expressions(given[4], given[6])...),
		),
	}

	assert(t, given, exp)
}

func Test_parseAssign_3(t *testing.T) {
	// a b = 0, 1

	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(token.Identifier, "a"),
		gen(token.Identifier, "b"),
		gen(token.Assign, "="),
		gen(token.Number, "0"),
		gen(token.Comma, ","),
		gen(token.Number, "1"),
		gen(token.Terminator, "\n"),
	}

	assertError(t, given, UnexpectedToken)
}

/* TODO: Move to validator
func Test_parseAssign_4(t *testing.T) {
	// a, b = 0 1

	given := []token.Token{
		tok(token.Identifier, "a"),
		tok(token.Comma, ","),
		tok(token.Identifier, "b"),
		tok(token.Assign, "="),
		tok(token.Number, "0"),
		tok(token.Number, "1"),
		tok(token.Terminator, "\n"),
	}

	assertError(t, given, MissingExpr)
}

func Test_parseAssign_5(t *testing.T) {
	// a, b = 0

	given := []token.Token{
		tok(token.Identifier, "a"),
		tok(token.Comma, ","),
		tok(token.Identifier, "b"),
		tok(token.Assign, "="),
		tok(token.Number, "0"),
		tok(token.Terminator, "\n"),
	}

	assertError(t, given, MissingExpr)
}

func Test_parseAssign_6(t *testing.T) {
	// a = 0, 1

	given := []token.Token{
		tok(token.Identifier, "a"),
		tok(token.Assign, "="),
		tok(token.Number, "0"),
		tok(token.Comma, ","),
		tok(token.Number, "1"),
		tok(token.Terminator, "\n"),
	}

	assertError(t, given, MissingVar)
}
*/

func Test_parseAssign_7(t *testing.T) {
	// a, b 0, 1

	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(token.Identifier, "a"),
		gen(token.Comma, ","),
		gen(token.Identifier, "b"),
		gen(token.Number, "0"),
		gen(token.Comma, ","),
		gen(token.Number, "1"),
		gen(token.Terminator, "\n"),
	}

	assertError(t, given, UnexpectedToken)
}

func Test_parseAssign_8(t *testing.T) {
	// a, b, c = false, 0, ""

	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(token.Identifier, "a"),
		gen(token.Comma, ","),
		gen(token.Identifier, "b"),
		gen(token.Comma, ","),
		gen(token.Identifier, "c"),
		gen(token.Assign, "="), // 5
		gen(token.False, "false"),
		gen(token.Comma, ","),
		gen(token.Number, "0"),
		gen(token.Comma, ","),
		gen(token.String, `""`), // 10
		gen(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		asttest.Assign(
			asttest.Variables(given[0], given[2], given[4]),
			given[5],
			asttest.ExprSet(
				asttest.Expressions(given[6], given[8], given[10])...,
			),
		),
	}

	assert(t, given, exp)
}
