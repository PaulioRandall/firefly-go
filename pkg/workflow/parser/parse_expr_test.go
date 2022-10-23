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

func Test_parseExpr_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// 1
	given := []token.Token{
		gen(token.Number, "1"),
		gen(token.Terminator, "\n"),
	}

	doParseTest(t, given, lit(given[0]))
}

func Test_parseExpr_2(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// "abc"
	given := []token.Token{
		gen(token.String, `"abc"`),
		gen(token.Terminator, "\n"),
	}

	doParseTest(t, given, lit(given[0]))
}

func Test_parseExpr_3(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// true
	given := []token.Token{
		gen(token.True, "true"),
		gen(token.Terminator, "\n"),
	}

	doParseTest(t, given, lit(given[0]))
}

func Test_parseExpr_4(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// false
	given := []token.Token{
		gen(token.False, "false"),
		gen(token.Terminator, "\n"),
	}

	doParseTest(t, given, lit(given[0]))
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
