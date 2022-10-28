package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func Test_for_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// for true
	// end
	given := []token.Token{
		gen(token.For, "for"),    // 0
		gen(token.True, "true"),  // 1
		gen(token.Newline, "\n"), // 2

		gen(token.End, "end"), // 3
		gen(token.Newline, "\n"),
	}

	exp := forStmt(
		given[0],
		nil,
		lit(given[1]),
		nil,
		nil,
		given[3],
	)

	doParseTest(t, given, exp)
}

func Test_for_2(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// for ;;
	// end
	given := []token.Token{
		gen(token.For, "for"),      // 0
		gen(token.Terminator, ";"), // 1
		gen(token.Terminator, ";"), // 2
		gen(token.Newline, "\n"),   // 3

		gen(token.End, "end"), // 4
		gen(token.Newline, "\n"),
	}

	exp := forStmt(
		given[0],
		nil,
		nil,
		nil,
		nil,
		given[4],
	)

	doParseTest(t, given, exp)
}

func Test_for_3(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// for i = 0; i < 5; i = i+1
	// end
	given := []token.Token{
		gen(token.For, "for"), // 0

		gen(token.Identifier, "i"), // 1
		gen(token.Assign, "="),     // 2
		gen(token.Number, "0"),     // 3
		gen(token.Terminator, ";"), // 4

		gen(token.Identifier, "i"), // 5
		gen(token.LT, "<"),         // 6
		gen(token.Number, "5"),     // 7
		gen(token.Terminator, ";"), // 8

		gen(token.Identifier, "i"), // 9
		gen(token.Assign, "="),     // 10
		gen(token.Identifier, "i"), // 11
		gen(token.Add, "+"),        // 12
		gen(token.Number, "1"),     // 13
		gen(token.Newline, "\n"),   // 14

		gen(token.End, "end"), // 15
		gen(token.Newline, "\n"),
	}

	initaliser := assStmt(
		vars(given[1]),
		given[2],
		lits(given[3]),
	)

	condition := binOp(
		varExpr(given[5]),
		given[6],
		lit(given[7]),
	)

	advancement := assStmt(
		vars(given[9]),
		given[10],
		[]ast.Expr{
			binOp(
				varExpr(given[11]),
				given[12],
				lit(given[13]),
			),
		},
	)

	exp := forStmt(
		given[0],
		initaliser,
		condition,
		advancement,
		nil,
		given[15],
	)

	doParseTest(t, given, exp)
}

func Test_for_4(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// Missing terminator after condition
	given := []token.Token{
		gen(token.For, "for"),   // 0
		gen(token.True, "true"), // 1

		gen(token.End, "end"), // 2
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrMissingTerminator,
		ErrBadForLoop,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_for_5(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// Missing terminator after at end of block
	given := []token.Token{
		gen(token.For, "for"),   // 0
		gen(token.True, "true"), // 1
		gen(token.Newline, "\n"),

		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrMissingEndOfBlock,
		ErrBadForLoop,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_for_6(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// Missing terminator after at end of block
	given := []token.Token{
		gen(token.For, "for"), // 0

		gen(token.Identifier, "i"), // 1
		gen(token.Assign, "="),     // 2
		gen(token.Number, "0"),     // 3
		gen(token.Terminator, ";"), // 4

		gen(token.Identifier, "i"), // 5
		gen(token.LT, "<"),         // 6
		gen(token.Number, "5"),     // 7

		gen(token.Identifier, "i"), // 8
		gen(token.Assign, "="),     // 9
		gen(token.Identifier, "i"), // 10
		gen(token.Add, "+"),        // 11
		gen(token.Number, "1"),     // 12
		gen(token.Newline, "\n"),   // 13

		gen(token.End, "end"), // 14
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrMissingTerminator,
		ErrBadForLoopControl,
		ErrBadForLoop,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_for_7(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// Missing terminator in for loop controls
	given := []token.Token{
		gen(token.For, "for"),      // 0
		gen(token.Terminator, ";"), // 1
		gen(token.Newline, "\n"),   // 2

		gen(token.End, "end"), // 3
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrMissingTerminator,
		ErrBadForLoopControl,
		ErrBadForLoop,
		ErrBadStmt,
		ErrParsing,
	)
}
