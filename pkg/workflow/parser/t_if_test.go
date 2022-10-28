package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func Test_if_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// if true
	// end
	given := []token.Token{
		gen(token.If, "if"),
		gen(token.True, "true"),
		gen(token.Newline, "\n"),
		gen(token.End, "end"),
		gen(token.Newline, "\n"),
	}

	exp := ifStmt(
		given[0],
		lit(given[1]),
		nil,
		given[3],
	)

	doParseTest(t, given, exp)
}

func Test_if_2(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// if true
	//   a = 0
	// end
	given := []token.Token{
		gen(token.If, "if"),
		gen(token.True, "true"),
		gen(token.Newline, "\n"),
		gen(token.Identifier, "a"),
		gen(token.Assign, "="), // 4
		gen(token.Number, "0"),
		gen(token.Newline, "\n"),
		gen(token.End, "end"),
		gen(token.Newline, "\n"), // 8
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

func Test_if_3(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// if true
	//   a = 0
	//
	// Missing 'end' of statement block
	given := []token.Token{
		gen(token.If, "if"),
		gen(token.True, "true"),
		gen(token.Newline, "\n"),
		gen(token.Identifier, "a"),
		gen(token.Assign, "="),
		gen(token.Number, "0"),
		gen(token.Newline, "\n"),
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrMissingEndOfBlock,
		ErrBadIfStmt,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_if_4(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// if true
	// end
	//
	// Missing terminator
	given := []token.Token{
		gen(token.If, "if"),
		gen(token.True, "true"),
		gen(token.End, "end"),
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrMissingTerminator,
		ErrBadIfStmt,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_if_5(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// if
	// end
	//
	// Missing condition expression
	given := []token.Token{
		gen(token.If, "if"),
		gen(token.Newline, "\n"),
		gen(token.End, "end"),
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrMissingExpr,
		ErrBadIfStmt,
		ErrBadStmt,
		ErrParsing,
	)
}
