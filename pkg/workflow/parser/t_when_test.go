package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func Test_when_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// when
	// end
	given := []token.Token{
		gen(token.When, "when"),
		gen(token.Newline, "\n"),
		gen(token.End, "end"),
		gen(token.Newline, "\n"),
	}

	exp := whenStmt(
		given[0],
		nil,
		nil,
		given[2],
	)

	doParseTest(t, given, exp)
}

func Test_when_2(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// when 1
	// end
	given := []token.Token{
		gen(token.When, "when"),  // 0
		gen(token.Number, "1"),   // 1
		gen(token.Newline, "\n"), //
		gen(token.End, "end"),    // 3
		gen(token.Newline, "\n"),
	}

	exp := whenStmt(
		given[0],
		lit(given[1]),
		nil,
		given[3],
	)

	doParseTest(t, given, exp)
}

func Test_when_3(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// when
	//   true:
	// end
	given := []token.Token{
		gen(token.When, "when"),  // 0
		gen(token.Newline, "\n"), //
		gen(token.Bool, "true"),  // 2
		gen(token.Colon, ":"),    // 3
		gen(token.Newline, "\n"), //
		gen(token.End, "end"),    // 5
		gen(token.Newline, "\n"),
	}

	cases := whenCases(
		whenCase(lit(given[2]), nil),
	)

	exp := whenStmt(
		given[0],
		nil,
		cases,
		given[5],
	)

	doParseTest(t, given, exp)
}

func Test_when_4(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// when
	//   1 == 2:
	// end
	given := []token.Token{
		gen(token.When, "when"),  // 0
		gen(token.Newline, "\n"), //
		gen(token.Number, "1"),   // 2
		gen(token.EQU, "=="),     // 3
		gen(token.Number, "2"),   // 4
		gen(token.Colon, ":"),    // 5
		gen(token.Newline, "\n"), //
		gen(token.End, "end"),    // 7
		gen(token.Newline, "\n"),
	}

	firstCase := binOp(
		lit(given[2]),
		given[3],
		lit(given[4]),
	)

	cases := whenCases(
		whenCase(firstCase, nil),
	)

	exp := whenStmt(
		given[0],
		nil,
		cases,
		given[7],
	)

	doParseTest(t, given, exp)
}

func Test_when_5(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// when
	//   a == 1:
	//   a == 2:
	//   a == 3:
	// end
	given := []token.Token{
		gen(token.When, "when"),    // 0
		gen(token.Newline, "\n"),   //
		gen(token.Identifier, "a"), // 2
		gen(token.EQU, "=="),       // 3
		gen(token.Number, "1"),     // 4
		gen(token.Colon, ":"),      // 5
		gen(token.Newline, "\n"),   //
		gen(token.Identifier, "a"), // 7
		gen(token.EQU, "=="),       // 8
		gen(token.Number, "2"),     // 9
		gen(token.Colon, ":"),      // 10
		gen(token.Newline, "\n"),   //
		gen(token.Identifier, "a"), // 12
		gen(token.EQU, "=="),       // 13
		gen(token.Number, "3"),     // 14
		gen(token.Colon, ":"),      // 15
		gen(token.Newline, "\n"),   //
		gen(token.End, "end"),      // 17
		gen(token.Newline, "\n"),
	}

	firstCase := binOp(
		varExpr(given[2]),
		given[3],
		lit(given[4]),
	)

	secondCase := binOp(
		varExpr(given[7]),
		given[8],
		lit(given[9]),
	)

	thirdCase := binOp(
		varExpr(given[12]),
		given[13],
		lit(given[14]),
	)

	cases := whenCases(
		whenCase(firstCase, nil),
		whenCase(secondCase, nil),
		whenCase(thirdCase, nil),
	)

	exp := whenStmt(
		given[0],
		nil,
		cases,
		given[17],
	)

	doParseTest(t, given, exp)
}

func Test_when_6(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// when a
	//   is 1:
	// end
	given := []token.Token{
		gen(token.When, "when"),    // 0
		gen(token.Identifier, "a"), // 1
		gen(token.Newline, "\n"),   //
		gen(token.Is, "is"),        // 3
		gen(token.Number, "1"),     // 4
		gen(token.Colon, ":"),      // 5
		gen(token.Newline, "\n"),   //
		gen(token.End, "end"),      // 7
		gen(token.Newline, "\n"),
	}

	firstCase := is(
		given[3],
		lit(given[4]),
	)

	cases := whenCases(
		whenCase(firstCase, nil),
	)

	exp := whenStmt(
		given[0],
		varExpr(given[1]),
		cases,
		given[7],
	)

	doParseTest(t, given, exp)
}

func Test_when_7(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// when a
	//   is 1:
	//   is 2:
	//   is 3:
	// end
	given := []token.Token{
		gen(token.When, "when"),    // 0
		gen(token.Identifier, "a"), // 1
		gen(token.Newline, "\n"),   //
		gen(token.Is, "is"),        // 3
		gen(token.Number, "1"),     // 4
		gen(token.Colon, ":"),      // 5
		gen(token.Newline, "\n"),   //
		gen(token.Is, "is"),        // 7
		gen(token.Number, "2"),     // 8
		gen(token.Colon, ":"),      // 9
		gen(token.Newline, "\n"),   //
		gen(token.Is, "is"),        // 11
		gen(token.Number, "3"),     // 12
		gen(token.Colon, ":"),      // 13
		gen(token.Newline, "\n"),   //
		gen(token.End, "end"),      // 15
		gen(token.Newline, "\n"),
	}

	firstCase := is(
		given[3],
		lit(given[4]),
	)

	secondCase := is(
		given[7],
		lit(given[8]),
	)

	thirdCase := is(
		given[11],
		lit(given[12]),
	)

	cases := whenCases(
		whenCase(firstCase, nil),
		whenCase(secondCase, nil),
		whenCase(thirdCase, nil),
	)

	exp := whenStmt(
		given[0],
		varExpr(given[1]),
		cases,
		given[15],
	)

	doParseTest(t, given, exp)
}

func Test_when_9(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// when a
	//   is 1:
	//   a == 2:
	//   is 3:
	//   a == 4:
	// end
	given := []token.Token{
		gen(token.When, "when"),    // 0
		gen(token.Identifier, "a"), // 1
		gen(token.Newline, "\n"),   //
		gen(token.Is, "is"),        // 3
		gen(token.Number, "1"),     // 4
		gen(token.Colon, ":"),      // 5
		gen(token.Newline, "\n"),   //
		gen(token.Identifier, "a"), // 7
		gen(token.EQU, "=="),       // 8
		gen(token.Number, "2"),     // 9
		gen(token.Colon, ":"),      // 10
		gen(token.Newline, "\n"),   //
		gen(token.Is, "is"),        // 12
		gen(token.Number, "3"),     // 13
		gen(token.Colon, ":"),      // 14
		gen(token.Newline, "\n"),   //
		gen(token.Identifier, "a"), // 16
		gen(token.EQU, "=="),       // 17
		gen(token.Number, "4"),     // 18
		gen(token.Colon, ":"),      // 19
		gen(token.Newline, "\n"),   //
		gen(token.End, "end"),      // 21
		gen(token.Newline, "\n"),
	}

	firstCase := is(
		given[3],
		lit(given[4]),
	)

	secondCase := binOp(
		varExpr(given[7]),
		given[8],
		lit(given[9]),
	)

	thirdCase := is(
		given[12],
		lit(given[13]),
	)

	fourthCase := binOp(
		varExpr(given[16]),
		given[17],
		lit(given[18]),
	)

	cases := whenCases(
		whenCase(firstCase, nil),
		whenCase(secondCase, nil),
		whenCase(thirdCase, nil),
		whenCase(fourthCase, nil),
	)

	exp := whenStmt(
		given[0],
		varExpr(given[1]),
		cases,
		given[21],
	)

	doParseTest(t, given, exp)
}

func Test_when_10(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// when a
	//   is 1:
	//   a == 2:
	//   true:
	// end
	given := []token.Token{
		gen(token.When, "when"),    // 0
		gen(token.Identifier, "a"), // 1
		gen(token.Newline, "\n"),   //
		gen(token.Is, "is"),        // 3
		gen(token.Number, "1"),     // 4
		gen(token.Colon, ":"),      // 5
		gen(token.Newline, "\n"),   //
		gen(token.Identifier, "a"), // 7
		gen(token.EQU, "=="),       // 8
		gen(token.Number, "2"),     // 9
		gen(token.Colon, ":"),      // 10
		gen(token.Newline, "\n"),   //
		gen(token.Bool, "true"),    // 12
		gen(token.Colon, ":"),      // 13
		gen(token.Newline, "\n"),   //
		gen(token.End, "end"),      // 15
		gen(token.Newline, "\n"),
	}

	firstCase := is(
		given[3],
		lit(given[4]),
	)

	secondCase := binOp(
		varExpr(given[7]),
		given[8],
		lit(given[9]),
	)

	thirdCase := lit(given[12])

	cases := whenCases(
		whenCase(firstCase, nil),
		whenCase(secondCase, nil),
		whenCase(thirdCase, nil),
	)

	exp := whenStmt(
		given[0],
		varExpr(given[1]),
		cases,
		given[15],
	)

	doParseTest(t, given, exp)
}

func Test_when_11(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// when a
	//   is 1
	// end
	//
	// Missing colon
	given := []token.Token{
		gen(token.When, "when"),    // 0
		gen(token.Identifier, "a"), // 1
		gen(token.Newline, "\n"),   //
		gen(token.Is, "is"),        // 3
		gen(token.Number, "1"),     // 4
		gen(token.Newline, "\n"),   //
		gen(token.End, "end"),      // 6
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrBadWhenCase,
		ErrBadWhenStmt,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_when_12(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// when
	//   a == 1
	// end
	//
	// Missing colon
	given := []token.Token{
		gen(token.When, "when"),    // 0
		gen(token.Identifier, "a"), // 1
		gen(token.Newline, "\n"),   //
		gen(token.Identifier, "a"), // 3
		gen(token.EQU, "=="),       // 4
		gen(token.Number, "1"),     // 5
		gen(token.Newline, "\n"),   //
		gen(token.End, "end"),      // 6
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrBadWhenCase,
		ErrBadWhenStmt,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_when_13(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// when
	//   a == 1:
	given := []token.Token{
		gen(token.When, "when"),    // 0
		gen(token.Identifier, "a"), // 1
		gen(token.Newline, "\n"),   //
		gen(token.Identifier, "a"), // 3
		gen(token.EQU, "=="),       // 4
		gen(token.Number, "1"),     // 5
		gen(token.Colon, ":"),      // 6
		gen(token.Newline, "\n"),
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrBadWhenCase,
		ErrBadWhenStmt,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_when_14(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// when a
	//   is:
	// end
	given := []token.Token{
		gen(token.When, "when"),    // 0
		gen(token.Identifier, "a"), // 1
		gen(token.Newline, "\n"),   //
		gen(token.Is, "is"),        // 3
		gen(token.Colon, ":"),      // 4
		gen(token.Newline, "\n"),   //
		gen(token.End, "end"),      // 6
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrBadWhenCase,
		ErrBadWhenStmt,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_when_15(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// when
	//   true: 1
	// end
	given := []token.Token{
		gen(token.When, "when"),  // 0
		gen(token.Newline, "\n"), //
		gen(token.Bool, "true"),  // 2
		gen(token.Colon, ":"),    // 3
		gen(token.Number, "1"),   // 4
		gen(token.Newline, "\n"), //
		gen(token.End, "end"),    // 6
		gen(token.Newline, "\n"),
	}

	firstCase := lit(given[2])
	firstBody := lit(given[4])

	cases := whenCases(
		whenCase(firstCase, firstBody),
	)

	exp := whenStmt(
		given[0],
		nil,
		cases,
		given[6],
	)

	doParseTest(t, given, exp)
}

func Test_when_16(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// when a
	//   is 0: a = 1
	//   true: 2 == 3
	// end
	given := []token.Token{
		gen(token.When, "when"),    // 0
		gen(token.Identifier, "a"), // 1
		gen(token.Newline, "\n"),   //

		gen(token.Is, "is"),        // 3
		gen(token.Number, "0"),     // 4
		gen(token.Colon, ":"),      // 5
		gen(token.Identifier, "a"), // 6
		gen(token.Assign, "="),     // 7
		gen(token.Number, "1"),     // 8
		gen(token.Newline, "\n"),   //

		gen(token.Bool, "true"),  // 10
		gen(token.Colon, ":"),    // 11
		gen(token.Number, "2"),   // 12
		gen(token.EQU, "=="),     // 13
		gen(token.Number, "3"),   // 14
		gen(token.Newline, "\n"), //

		gen(token.End, "end"), // 16
		gen(token.Newline, "\n"),
	}

	subject := varExpr(given[1])

	firstCase := is(
		given[3],
		lit(given[4]),
	)

	firstBody := assStmt(
		vars(given[6]),
		given[7],
		lits(given[8]),
	)

	secondCase := lit(given[10])

	secondBody := binOp(
		lit(given[12]),
		given[13],
		lit(given[14]),
	)

	cases := whenCases(
		whenCase(firstCase, firstBody),
		whenCase(secondCase, secondBody),
	)

	exp := whenStmt(
		given[0],
		subject,
		cases,
		given[16],
	)

	doParseTest(t, given, exp)
}
