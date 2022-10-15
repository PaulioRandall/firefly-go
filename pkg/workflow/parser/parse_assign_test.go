package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/auditor"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast/asttest"
)

func Test_parseAssign_1(t *testing.T) {
	// a = 0

	given := []token.Token{
		tok(token.Identifier, "a"),
		tok(token.Assign, "="),
		tok(token.Number, "0"),
		tok(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		asttest.Assign(
			asttest.Variables(given[0]),
			given[1],
			asttest.Expressions(given[2]),
		),
	}

	assert(t, given, exp)
}

func Test_parseAssign_2(t *testing.T) {
	// a, b = 0, 1

	given := []token.Token{
		tok(token.Identifier, "a"),
		tok(token.Comma, ","),
		tok(token.Identifier, "b"),
		tok(token.Assign, "="),
		tok(token.Number, "0"),
		tok(token.Comma, ","),
		tok(token.Number, "1"),
		tok(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		asttest.Assign(
			asttest.Variables(given[0], given[2]),
			given[3],
			asttest.Expressions(given[4], given[6]),
		),
	}

	assert(t, given, exp)
}

func Test_parseAssign_3(t *testing.T) {
	// a b = 0, 1

	given := []token.Token{
		tok(token.Identifier, "a"),
		tok(token.Identifier, "b"),
		tok(token.Assign, "="),
		tok(token.Number, "0"),
		tok(token.Comma, ","),
		tok(token.Number, "1"),
		tok(token.Terminator, "\n"),
	}

	assertError(t, given, auditor.UnexpectedToken)
}

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

func Test_parseAssign_7(t *testing.T) {
	// a, b 0, 1

	given := []token.Token{
		tok(token.Identifier, "a"),
		tok(token.Comma, ","),
		tok(token.Identifier, "b"),
		tok(token.Number, "0"),
		tok(token.Comma, ","),
		tok(token.Number, "1"),
		tok(token.Terminator, "\n"),
	}

	assertError(t, given, auditor.UnexpectedToken)
}

func Test_parseAssign_8(t *testing.T) {
	// a, b, c = false, 0, ""

	given := []token.Token{
		tok(token.Identifier, "a"),
		tok(token.Comma, ","),
		tok(token.Identifier, "b"),
		tok(token.Comma, ","),
		tok(token.Identifier, "c"),
		tok(token.Assign, "="), // 5
		tok(token.False, "false"),
		tok(token.Comma, ","),
		tok(token.Number, "0"),
		tok(token.Comma, ","),
		tok(token.String, `""`), // 10
		tok(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		asttest.Assign(
			asttest.Variables(given[0], given[2], given[4]),
			given[5],
			asttest.Expressions(given[6], given[8], given[10]),
		),
	}

	assert(t, given, exp)
}
