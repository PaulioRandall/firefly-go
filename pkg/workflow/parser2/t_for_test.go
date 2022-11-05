package parser2

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

	// for true end
	//
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

	// for true
	//
	// Missing end of block
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

	// for i = 0; i < 5 i = i+1
	// end
	//
	// Missing terminator after condition
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

	// for ;
	// end
	//
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

func Test_for_each_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// for i in []
	// end
	given := []token.Token{
		gen(token.For, "for"),        // 0
		gen(token.Identifier, "i"),   // 1
		gen(token.In, "in"),          // 2
		gen(token.BracketOpen, "["),  // 3
		gen(token.BracketClose, "]"), // 4
		gen(token.Newline, "\n"),     // 5

		gen(token.End, "end"), // 6
		gen(token.Newline, "\n"),
	}

	vars := seriesOfVar(given[1])
	iterable := listExpr(given[3], nil, given[4])

	exp := forEachStmt(
		given[0],
		vars,
		iterable,
		nil,
		given[6],
	)

	doParseTest(t, given, exp)
}

func Test_for_each_2(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// for i, v in []
	// end
	given := []token.Token{
		gen(token.For, "for"),        // 0
		gen(token.Identifier, "i"),   // 1
		gen(token.Comma, ","),        // 2
		gen(token.Identifier, "v"),   // 3
		gen(token.In, "in"),          // 4
		gen(token.BracketOpen, "["),  // 5
		gen(token.BracketClose, "]"), // 6
		gen(token.Newline, "\n"),     // 7

		gen(token.End, "end"), // 8
		gen(token.Newline, "\n"),
	}

	vars := seriesOfVar(given[1], given[3])
	iterable := listExpr(given[5], nil, given[6])

	exp := forEachStmt(
		given[0],
		vars,
		iterable,
		nil,
		given[8],
	)

	doParseTest(t, given, exp)
}

func Test_for_each_3(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// for i, v, more in []
	// end
	given := []token.Token{
		gen(token.For, "for"),         // 0
		gen(token.Identifier, "i"),    // 1
		gen(token.Comma, ","),         // 2
		gen(token.Identifier, "v"),    // 3
		gen(token.Comma, ","),         // 4
		gen(token.Identifier, "more"), // 5
		gen(token.In, "in"),           // 6
		gen(token.BracketOpen, "["),   // 7
		gen(token.BracketClose, "]"),  // 8
		gen(token.Newline, "\n"),      // 9

		gen(token.End, "end"), // 10
		gen(token.Newline, "\n"),
	}

	vars := seriesOfVar(given[1], given[3], given[5])
	iterable := listExpr(given[7], nil, given[8])

	exp := forEachStmt(
		given[0],
		vars,
		iterable,
		nil,
		given[10],
	)

	doParseTest(t, given, exp)
}

func Test_for_each_4(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// for i in x
	// end
	given := []token.Token{
		gen(token.For, "for"),      // 0
		gen(token.Identifier, "i"), // 1
		gen(token.In, "in"),        // 2
		gen(token.Identifier, "x"), // 3
		gen(token.Newline, "\n"),   // 4

		gen(token.End, "end"), // 5
		gen(token.Newline, "\n"),
	}

	vars := seriesOfVar(given[1])
	iterable := varExpr(given[3])

	exp := forEachStmt(
		given[0],
		vars,
		iterable,
		nil,
		given[5],
	)

	doParseTest(t, given, exp)
}

func Test_for_each_5(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// for k, v in {"one": 1}
	// end
	given := []token.Token{
		gen(token.For, "for"),      // 0
		gen(token.Identifier, "k"), // 1
		gen(token.Comma, ","),      // 2
		gen(token.Identifier, "v"), // 3
		gen(token.In, "in"),        // 4
		gen(token.BraceOpen, "{"),  // 5
		gen(token.String, `"one"`), // 6
		gen(token.Colon, ":"),      // 7
		gen(token.Number, "1"),     // 8
		gen(token.BraceClose, "}"), // 9
		gen(token.Newline, "\n"),   // 410

		gen(token.End, "end"), // 11
		gen(token.Newline, "\n"),
	}

	vars := seriesOfVar(given[1], given[3])

	iterable := mapExpr(
		given[5],
		mapEntries(
			mapEntry(given[6], given[8]),
		),
		given[9],
	)

	exp := forEachStmt(
		given[0],
		vars,
		iterable,
		nil,
		given[11],
	)

	doParseTest(t, given, exp)
}

func Test_for_each_6(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// for i in
	// end
	given := []token.Token{
		gen(token.For, "for"),      // 0
		gen(token.Identifier, "i"), // 1
		gen(token.In, "in"),        // 2
		gen(token.Newline, "\n"),   // 4

		gen(token.End, "end"), // 5
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrMissingExpr,
		ErrBadForEachLoop,
		ErrBadStmt,
		ErrParsing,
	)
}
