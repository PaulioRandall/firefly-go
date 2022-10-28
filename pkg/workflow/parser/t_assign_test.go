package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func Test_assign_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// a = 0
	given := []token.Token{
		gen(token.Identifier, "a"),
		gen(token.Assign, "="),
		gen(token.Number, "0"),
		gen(token.Newline, "\n"),
	}

	exp := assStmt(
		vars(given[0]),
		given[1],
		lits(given[2]),
	)

	doParseTest(t, given, exp)
}

func Test_assign_2(t *testing.T) {
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
		gen(token.Newline, "\n"),
	}

	exp := assStmt(
		vars(given[0], given[2]),
		given[3],
		lits(given[4], given[6]),
	)

	doParseTest(t, given, exp)
}

func Test_assign_3(t *testing.T) {
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
		gen(token.Newline, "\n"),
	}

	exp := assStmt(
		vars(given[0], given[2], given[4]),
		given[5],
		lits(given[6], given[8], given[10]),
	)

	doParseTest(t, given, exp)
}

func Test_assign_4(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// a b = 0, 1
	//
	// Missing comma, not an assignment error
	given := []token.Token{
		gen(token.Identifier, "a"),
		gen(token.Identifier, "b"),
		gen(token.Assign, "="),
		gen(token.Number, "0"),
		gen(token.Comma, ","),
		gen(token.Number, "1"),
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrMissingTerminator,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_assign_5(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// a, b 0, 1
	//
	// Missing assignment operator
	given := []token.Token{
		gen(token.Identifier, "a"),
		gen(token.Comma, ","),
		gen(token.Identifier, "b"),
		gen(token.Number, "0"),
		gen(token.Comma, ","),
		gen(token.Number, "1"),
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrBadAssign,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_assign_6(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// a, b = 0 1
	//
	// Missing comma so looked for terminator, not an assignment error
	given := []token.Token{
		gen(token.Identifier, "a"),
		gen(token.Comma, ","),
		gen(token.Identifier, "b"),
		gen(token.Assign, "="),
		gen(token.Number, "0"),
		gen(token.Number, "1"),
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrMissingTerminator,
		ErrBadStmt,
		ErrParsing,
	)
}
