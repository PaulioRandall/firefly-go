package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func Test_parseWhen_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// when
	// end
	given := []token.Token{
		gen(token.When, "when"),
		gen(token.Terminator, "\n"),
		gen(token.End, "end"),
		gen(token.Terminator, "\n"),
	}

	exp := whenStmt(
		given[0],
		nil,
		nil,
		given[2],
	)

	doParseTest(t, given, exp)
}

func Test_parseWhen_2(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// when 1
	// end
	given := []token.Token{
		gen(token.When, "when"),
		gen(token.Number, "1"),
		gen(token.Terminator, "\n"),
		gen(token.End, "end"),
		gen(token.Terminator, "\n"),
	}

	exp := whenStmt(
		given[0],
		lit(given[1]),
		nil,
		given[3],
	)

	doParseTest(t, given, exp)
}

func Test_parseWhen_3(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// when
	//   true:
	// end
	given := []token.Token{
		gen(token.When, "when"),
		gen(token.Terminator, "\n"),
		gen(token.True, "true"),
		gen(token.Colon, ":"),
		gen(token.Terminator, "\n"),
		gen(token.End, "end"),
		gen(token.Terminator, "\n"),
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

func Test_parseWhen_4(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// when
	//   1 == 2:
	// end
	given := []token.Token{
		gen(token.When, "when"), // 0
		gen(token.Terminator, "\n"),
		gen(token.Number, "1"), // 2
		gen(token.EQU, "=="),   // 3
		gen(token.Number, "2"), // 4
		gen(token.Colon, ":"),  // 5
		gen(token.Terminator, "\n"),
		gen(token.End, "end"), // 7
		gen(token.Terminator, "\n"),
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

func Test_parseWhen_5(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// when
	//   a == 1:
	//   a == 2:
	//   a == 3:
	// end
	given := []token.Token{
		gen(token.When, "when"), // 0
		gen(token.Terminator, "\n"),
		gen(token.Identifier, "a"), // 2
		gen(token.EQU, "=="),       // 3
		gen(token.Number, "1"),     // 4
		gen(token.Colon, ":"),      // 5
		gen(token.Terminator, "\n"),
		gen(token.Identifier, "a"), // 7
		gen(token.EQU, "=="),       // 8
		gen(token.Number, "2"),     // 9
		gen(token.Colon, ":"),      // 10
		gen(token.Terminator, "\n"),
		gen(token.Identifier, "a"), // 12
		gen(token.EQU, "=="),       // 13
		gen(token.Number, "3"),     // 14
		gen(token.Colon, ":"),      // 15
		gen(token.Terminator, "\n"),
		gen(token.End, "end"), // 17
		gen(token.Terminator, "\n"),
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

func Test_parseWhen_6(t *testing.T) {
	// when a
	//   is 1:
	// end
}

func Test_parseWhen_7(t *testing.T) {
	// when a
	//   is 1:
	//   is 2:
	//   is 3:
	// end
}

func Test_parseWhen_9(t *testing.T) {
	// when a
	//   is 1:
	//   a == 2:
	//   is 3:
	//   a == 4:
	// end
}

func Test_parseWhen_10(t *testing.T) {
	// when a
	//   is 1:
	//   a == 2:
	//   true:
	// end
}

func Test_parseWhen_11(t *testing.T) {
	// when
	//   is 1
	// end

	// Error!
}

func Test_parseWhen_12(t *testing.T) {
	// when
	//   a == 1
	// end

	// Error!
}

func Test_parseWhen_13(t *testing.T) {
	// when
	//   a == 1

	// Error!
}

func Test_parseWhen_14(t *testing.T) {
	// when a
	//   is:
	// end

	// Error!
}
