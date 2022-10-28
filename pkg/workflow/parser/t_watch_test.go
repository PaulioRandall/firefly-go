package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func Test_watch_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// watch e
	// end
	given := []token.Token{
		gen(token.Watch, "watch"),   // 0
		gen(token.Identifier, "e"),  // 1
		gen(token.Terminator, "\n"), // 2
		gen(token.End, "end"),       // 3
		gen(token.Terminator, "\n"),
	}

	exp := watchStmt(
		given[0],
		varExpr(given[1]),
		nil,
		given[3],
	)

	doParseTest(t, given, exp)
}

func Test_watch_2(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// watch e
	//   true
	// end
	given := []token.Token{
		gen(token.Watch, "watch"),   // 0
		gen(token.Identifier, "e"),  // 1
		gen(token.Terminator, "\n"), // 2
		gen(token.True, "true"),     // 3
		gen(token.Terminator, "\n"), // 4
		gen(token.End, "end"),       // 5
		gen(token.Terminator, "\n"),
	}

	body := stmts(
		lit(given[3]),
	)

	exp := watchStmt(
		given[0],
		varExpr(given[1]),
		body,
		given[5],
	)

	doParseTest(t, given, exp)
}

func Test_watch_3(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// watch e
	//   true
	//   e = "error"
	// end
	given := []token.Token{
		gen(token.Watch, "watch"),    // 0
		gen(token.Identifier, "e"),   // 1
		gen(token.Terminator, "\n"),  // 2
		gen(token.True, "true"),      // 3
		gen(token.Terminator, "\n"),  // 4
		gen(token.Identifier, "e"),   // 5
		gen(token.Assign, "="),       // 6
		gen(token.String, `"error"`), // 7
		gen(token.Terminator, "\n"),  // 8
		gen(token.End, "end"),        // 9
		gen(token.Terminator, "\n"),
	}

	body := stmts(
		lit(given[3]),
		assStmt(
			vars(given[5]),
			given[6],
			lits(given[7]),
		),
	)

	exp := watchStmt(
		given[0],
		varExpr(given[1]),
		body,
		given[9],
	)

	doParseTest(t, given, exp)
}

func Test_watch_4(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// watch e
	//
	// Missing 'end' of block
	given := []token.Token{
		gen(token.Watch, "watch"),   // 0
		gen(token.Identifier, "e"),  // 1
		gen(token.Terminator, "\n"), // 2
		gen(token.Terminator, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrMissingEndOfBlock,
		ErrBadWatchStmt,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_watch_5(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// watch e
	//
	// Missing terminator after variable
	given := []token.Token{
		gen(token.Watch, "watch"),  // 0
		gen(token.Identifier, "e"), // 1
		gen(token.End, "end"),      // 2
		gen(token.Terminator, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrMissingTerminator,
		ErrBadWatchStmt,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_watch_6(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// watch 1
	// end
	//
	// Subject must be a variable
	given := []token.Token{
		gen(token.Watch, "watch"),   // 0
		gen(token.Number, "1"),      // 1
		gen(token.Terminator, "\n"), // 2
		gen(token.End, "end"),       // 3
		gen(token.Terminator, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrMissingVariable,
		ErrBadWatchStmt,
		ErrBadStmt,
		ErrParsing,
	)
}
