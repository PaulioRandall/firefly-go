package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func Test_parseAssign_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// a = 0
	given := []token.Token{
		gen(token.Identifier, "a"),
		gen(token.Assign, "="),
		gen(token.Number, "0"),
		gen(token.Terminator, "\n"),
	}

	exp := assStmt(
		vars(given[0]),
		given[1],
		lits(given[2]),
	)

	doParseTest(t, given, exp)
}

func Test_parseAssign_2(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// a, b = 0, 1
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

	exp := assStmt(
		vars(given[0], given[2]),
		given[3],
		lits(given[4], given[6]),
	)

	doParseTest(t, given, exp)
}

func Test_parseAssign_3(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// a b = 0, 1
	given := []token.Token{
		gen(token.Identifier, "a"),
		gen(token.Identifier, "b"),
		gen(token.Assign, "="),
		gen(token.Number, "0"),
		gen(token.Comma, ","),
		gen(token.Number, "1"),
		gen(token.Terminator, "\n"),
	}

	doErrorTest(t, given, ErrUnexpectedToken)
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

	doErrorTest(t, given, MissingExpr)
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

	doErrorTest(t, given, MissingExpr)
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

	doErrorTest(t, given, MissingVar)
}
*/

func Test_parseAssign_7(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// a, b 0, 1
	given := []token.Token{
		gen(token.Identifier, "a"),
		gen(token.Comma, ","),
		gen(token.Identifier, "b"),
		gen(token.Number, "0"),
		gen(token.Comma, ","),
		gen(token.Number, "1"),
		gen(token.Terminator, "\n"),
	}

	doErrorTest(t, given, ErrUnexpectedToken)
}

func Test_parseAssign_8(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// a, b, c = false, 0, ""
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

	exp := assStmt(
		vars(given[0], given[2], given[4]),
		given[5],
		lits(given[6], given[8], given[10]),
	)

	doParseTest(t, given, exp)
}
