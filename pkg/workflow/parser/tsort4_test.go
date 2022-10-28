package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func Test_parseIf_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// if true
	// end
	given := []token.Token{
		gen(token.If, "if"),
		gen(token.True, "true"),
		gen(token.Terminator, "\n"),
		gen(token.End, "end"),
		gen(token.Terminator, "\n"),
	}

	exp := ifStmt(
		given[0],
		lit(given[1]),
		nil,
		given[3],
	)

	doParseTest(t, given, exp)
}

func Test_parseIf_2(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// if true
	//   a = 0
	// end
	given := []token.Token{
		gen(token.If, "if"),
		gen(token.True, "true"),
		gen(token.Terminator, "\n"),
		gen(token.Identifier, "a"),
		gen(token.Assign, "="), // 4
		gen(token.Number, "0"),
		gen(token.Terminator, "\n"),
		gen(token.End, "end"),
		gen(token.Terminator, "\n"), // 8
	}

	body := stmts(
		assStmt(
			vars(given[3]),
			given[4],
			lits(given[5]),
		),
	)

	exp := ifStmt(
		given[0],
		lit(given[1]),
		body,
		given[7],
	)

	doParseTest(t, given, exp)
}

func Test_parseIf_3(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// if true
	//   a = 0
	given := []token.Token{
		gen(token.If, "if"),
		gen(token.True, "true"),
		gen(token.Terminator, "\n"),
		gen(token.Identifier, "a"),
		gen(token.Assign, "="),
		gen(token.Number, "0"),
		gen(token.Terminator, "\n"),
	}

	doErrorTest(t, given,
		//ErrUnexpectedToken,
		//ErrMissingTerminator,
		ErrBadIfStmt,
		ErrBadStmt,
		ErrParsing,
	)
}
