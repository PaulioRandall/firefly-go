package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func doBinExprTest(t *testing.T, given ...token.Token) {
	exp := binOp(
		lit(given[0]),
		given[1],
		lit(given[2]),
	)

	doParseTest(t, given, exp)
}

func Test_parseExpr_6(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// 1 + 2 + 3
	// (1 + 2) + 3
	given := []token.Token{
		gen(token.Number, "1"),
		gen(token.Add, "+"),
		gen(token.Number, "2"),
		gen(token.Add, "+"),
		gen(token.Number, "3"),
		gen(token.Terminator, "\n"),
	}

	// 1 + 2
	a := binOp(
		lit(given[0]),
		given[1],
		lit(given[2]),
	)

	// a + 3
	exp := binOp(
		a,
		given[3],
		lit(given[4]),
	)

	doParseTest(t, given, exp)
}

func Test_parseExpr_7(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// 1 + 2 * 3
	// 1 + (2 * 3)
	given := []token.Token{
		gen(token.Number, "1"),
		gen(token.Add, "+"),
		gen(token.Number, "2"),
		gen(token.Mul, "*"),
		gen(token.Number, "3"),
		gen(token.Terminator, "\n"),
	}

	// 2 * 3
	a := binOp(
		lit(given[2]),
		given[3],
		lit(given[4]),
	)

	// 1 + a
	exp := binOp(
		lit(given[0]),
		given[1],
		a,
	)

	doParseTest(t, given, exp)
}

func Test_parseExpr_8(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// 1 * 2 + 3
	// (1 * 2) + 3
	given := []token.Token{
		gen(token.Number, "1"),
		gen(token.Mul, "*"),
		gen(token.Number, "2"),
		gen(token.Add, "+"),
		gen(token.Number, "3"),
		gen(token.Terminator, "\n"),
	}

	// 1 * 2
	a := binOp(
		lit(given[0]),
		given[1],
		lit(given[2]),
	)

	// a + 3
	exp := binOp(
		a,
		given[3],
		lit(given[4]),
	)

	doParseTest(t, given, exp)
}

func Test_parseExpr_9(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// 1 * 2 + 3 * 4
	// (1 * 2) + (3 * 4)
	given := []token.Token{
		gen(token.Number, "1"),
		gen(token.Mul, "*"),
		gen(token.Number, "2"),
		gen(token.Add, "+"),
		gen(token.Number, "3"),
		gen(token.Mul, "*"),
		gen(token.Number, "4"),
		gen(token.Terminator, "\n"),
	}

	// 1 * 2
	a := binOp(
		lit(given[0]),
		given[1],
		lit(given[2]),
	)

	// 3 * 4
	b := binOp(
		lit(given[4]),
		given[5],
		lit(given[6]),
	)

	// a + b
	exp := binOp(
		a,
		given[3],
		b,
	)

	doParseTest(t, given, exp)
}

func Test_parseExpr_10(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	doBinExprTest(t,
		// 1 + 2
		gen(token.Number, "1"),
		gen(token.Add, "+"),
		gen(token.Number, "2"),
		gen(token.Terminator, "\n"),
	)
}

func Test_parseExpr_11(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	doBinExprTest(t,
		// 1 - 2
		gen(token.Number, "1"),
		gen(token.Sub, "-"),
		gen(token.Number, "2"),
		gen(token.Terminator, "\n"),
	)
}

func Test_parseExpr_12(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	doBinExprTest(t,
		// 1 * 2
		gen(token.Number, "1"),
		gen(token.Mul, "*"),
		gen(token.Number, "2"),
		gen(token.Terminator, "\n"),
	)
}

func Test_parseExpr_13(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	doBinExprTest(t,
		// 1 / 2
		gen(token.Number, "1"),
		gen(token.Div, "/"),
		gen(token.Number, "2"),
		gen(token.Terminator, "\n"),
	)
}

func Test_parseExpr_14(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	doBinExprTest(t,
		// 1 % 2
		gen(token.Number, "1"),
		gen(token.Mod, "%"),
		gen(token.Number, "2"),
		gen(token.Terminator, "\n"),
	)
}

func Test_parseExpr_15(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	doBinExprTest(t,
		// 1 < 2
		gen(token.Number, "1"),
		gen(token.LT, "<"),
		gen(token.Number, "2"),
		gen(token.Terminator, "\n"),
	)
}

func Test_parseExpr_16(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	doBinExprTest(t,
		// 1 > 2
		gen(token.Number, "1"),
		gen(token.GT, ">"),
		gen(token.Number, "2"),
		gen(token.Terminator, "\n"),
	)
}

func Test_parseExpr_17(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	doBinExprTest(t,
		// 1 <= 2
		gen(token.Number, "1"),
		gen(token.LTE, "<="),
		gen(token.Number, "2"),
		gen(token.Terminator, "\n"),
	)
}

func Test_parseExpr_18(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	doBinExprTest(t,
		// 1 >= 2
		gen(token.Number, "1"),
		gen(token.GTE, ">="),
		gen(token.Number, "2"),
		gen(token.Terminator, "\n"),
	)
}

func Test_parseExpr_19(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	doBinExprTest(t,
		// 1 == 2
		gen(token.Number, "1"),
		gen(token.EQU, "=="),
		gen(token.Number, "2"),
		gen(token.Terminator, "\n"),
	)
}

func Test_parseExpr_20(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	doBinExprTest(t,
		// 1 != 2
		gen(token.Number, "1"),
		gen(token.NEQ, "!="),
		gen(token.Number, "2"),
		gen(token.Terminator, "\n"),
	)
}

func Test_parseExpr_21(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// (1 + 2)
	given := []token.Token{
		gen(token.ParenOpen, "("),
		gen(token.Number, "1"),
		gen(token.Add, "+"),
		gen(token.Number, "2"),
		gen(token.ParenClose, ")"),
		gen(token.Terminator, "\n"),
	}

	// 1 + 2
	exp := binOp(
		lit(given[1]),
		given[2],
		lit(given[3]),
	)

	doParseTest(t, given, exp)
}

func Test_parseExpr_22(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// (1 + 2) * 3
	given := []token.Token{
		gen(token.ParenOpen, "("),  // 0
		gen(token.Number, "1"),     // 1
		gen(token.Add, "+"),        // 2
		gen(token.Number, "2"),     // 3
		gen(token.ParenClose, ")"), // 4
		gen(token.Mul, "*"),        // 5
		gen(token.Number, "3"),     // 6
		gen(token.Terminator, "\n"),
	}

	// (1 + 2)
	a := binOp(
		lit(given[1]),
		given[2],
		lit(given[3]),
	)

	// a + 3
	exp := binOp(
		a,
		given[5],
		lit(given[6]),
	)

	doParseTest(t, given, exp)
}

func Test_parseExpr_23(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// 1 * (2 + 3)
	given := []token.Token{
		gen(token.Number, "1"),     // 0
		gen(token.Mul, "*"),        // 1
		gen(token.ParenOpen, "("),  // 2
		gen(token.Number, "2"),     // 3
		gen(token.Add, "+"),        // 4
		gen(token.Number, "3"),     // 5
		gen(token.ParenClose, ")"), // 6
		gen(token.Terminator, "\n"),
	}

	// (2 + 3)
	a := binOp(
		lit(given[3]),
		given[4],
		lit(given[5]),
	)

	// 1 * a
	exp := binOp(
		lit(given[0]),
		given[1],
		a,
	)

	doParseTest(t, given, exp)
}

func Test_parseExpr_24(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// (1 + 2) * (3 + 4)
	given := []token.Token{
		gen(token.ParenOpen, "("),  // 0
		gen(token.Number, "1"),     // 1
		gen(token.Add, "+"),        // 2
		gen(token.Number, "2"),     // 3
		gen(token.ParenClose, ")"), // 4
		gen(token.Mul, "*"),        // 5
		gen(token.ParenOpen, "("),  // 6
		gen(token.Number, "3"),     // 7
		gen(token.Add, "+"),        // 8
		gen(token.Number, "4"),     // 9
		gen(token.ParenClose, ")"), // 10
		gen(token.Terminator, "\n"),
	}

	// (1 + 2)
	a := binOp(
		lit(given[1]),
		given[2],
		lit(given[3]),
	)

	// (3 + 4)
	b := binOp(
		lit(given[7]),
		given[8],
		lit(given[9]),
	)

	// a * b
	exp := binOp(
		a,
		given[5],
		b,
	)

	doParseTest(t, given, exp)
}

func Test_parseExpr_25(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// (1 * (2 + 3) / 4)
	given := []token.Token{
		gen(token.ParenOpen, "("),  // 0
		gen(token.Number, "1"),     // 1
		gen(token.Mul, "*"),        // 2
		gen(token.ParenOpen, "("),  // 3
		gen(token.Number, "2"),     // 4
		gen(token.Add, "+"),        // 5
		gen(token.Number, "3"),     // 6
		gen(token.ParenClose, ")"), // 7
		gen(token.Div, "/"),        // 8
		gen(token.Number, "4"),     // 9
		gen(token.ParenClose, ")"), // 10
		gen(token.Terminator, "\n"),
	}

	// (2 + 3)
	a := binOp(
		lit(given[4]),
		given[5],
		lit(given[6]),
	)

	// 1 * a
	b := binOp(
		lit(given[1]),
		given[2],
		a,
	)

	// b / 4
	exp := binOp(
		b,
		given[8],
		lit(given[9]),
	)

	doParseTest(t, given, exp)
}

func Test_parseExpr_26(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// 1 * (2 + (3 - 4))
	given := []token.Token{
		gen(token.Number, "1"),     // 0
		gen(token.Mul, "*"),        // 1
		gen(token.ParenOpen, "("),  // 2
		gen(token.Number, "2"),     // 3
		gen(token.Add, "+"),        // 4
		gen(token.ParenOpen, "("),  // 5
		gen(token.Number, "3"),     // 6
		gen(token.Sub, "-"),        // 7
		gen(token.Number, "4"),     // 8
		gen(token.ParenClose, ")"), // 9
		gen(token.ParenClose, ")"), // 10
		gen(token.Terminator, "\n"),
	}

	// (3 - 4)
	a := binOp(
		lit(given[6]),
		given[7],
		lit(given[8]),
	)

	// (2 + a)
	b := binOp(
		lit(given[3]),
		given[4],
		a,
	)

	// 1 * b
	exp := binOp(
		lit(given[0]),
		given[1],
		b,
	)

	doParseTest(t, given, exp)
}

func Test_parseExpr_27(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// (((1 + 2)))
	given := []token.Token{
		gen(token.ParenOpen, "("),  // 0
		gen(token.ParenOpen, "("),  // 1
		gen(token.ParenOpen, "("),  // 2
		gen(token.Number, "1"),     // 3
		gen(token.Add, "+"),        // 4
		gen(token.Number, "2"),     // 5
		gen(token.ParenClose, ")"), // 6
		gen(token.ParenClose, ")"), // 7
		gen(token.ParenClose, ")"), // 8
		gen(token.Terminator, "\n"),
	}

	// (((1 + 2))))
	exp := binOp(
		lit(given[3]),
		given[4],
		lit(given[5]),
	)

	doParseTest(t, given, exp)
}

func Test_parseExpr_28(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// (1)
	given := []token.Token{
		gen(token.ParenOpen, "("),  // 0
		gen(token.Number, "1"),     // 1
		gen(token.ParenClose, ")"), // 2
		gen(token.Terminator, "\n"),
	}

	// (1)
	exp := lit(given[1])

	doParseTest(t, given, exp)
}
