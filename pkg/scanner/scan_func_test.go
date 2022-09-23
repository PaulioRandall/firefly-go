package scanner

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/err"
	"github.com/PaulioRandall/firefly-go/pkg/readers"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func tok(tt token.TokenType, v string) token.Token {
	return token.MakeToken(tt, v, token.MakeInlineRange(0, 0, 0, len(v)))
}

func assertAllTokensScan(t *testing.T, given string, exp []token.Token) {
	r := readers.NewRuneStringReader(given)

	act, e := ScanAll(r)
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func assertTokenScans(t *testing.T, given string, exp token.TokenType) {
	r := readers.NewRuneStringReader(given)

	actTk, e := ScanAll(r)
	expTk := []token.Token{
		tok(exp, given),
	}

	require.Nil(t, e, "Expected %q but got %+v", exp.String(), err.DebugString(e))
	require.NotEmpty(t, actTk)
	require.Equal(t, expTk, actTk,
		"Expected %q but got %q", exp.String(), actTk[0].Type.String(),
	)
}

func assertScanError(t *testing.T, given string, exp error) {
	r := readers.NewRuneStringReader(given)
	_, e := ScanAll(r)
	require.True(t, errors.Is(e, exp), "Expected %+v", exp.Error())
}

func Test_1_ScanAll(t *testing.T) {
	r := readers.NewRuneStringReader("")

	act, e := ScanAll(r)

	require.Nil(t, e)
	require.Empty(t, act)
}

func Test_7_ScanAll(t *testing.T) {
	given := "\n"
	exp := []token.Token{
		token.MakeToken(
			token.Newline,
			given,
			token.MakeRange(
				token.MakePos(0, 0, 0),
				token.MakePos(1, 1, 0),
			),
		),
	}
	assertAllTokensScan(t, given, exp)
}

func Test_10_ScanAll(t *testing.T) {
	assertTokenScans(t, "=", token.Ass)
}

func Test_11_ScanAll(t *testing.T) {
	assertTokenScans(t, ":=", token.Def)
}

func Test_12_ScanAll(t *testing.T) {
	assertTokenScans(t, ";", token.Terminator)
}

func Test_13_ScanAll(t *testing.T) {
	assertTokenScans(t, ",", token.Comma)
}

func Test_14_ScanAll(t *testing.T) {
	assertTokenScans(t, ":", token.Colon)
}

func Test_15_ScanAll(t *testing.T) {
	assertTokenScans(t, "@", token.Spell)
}

func Test_16_ScanAll(t *testing.T) {
	assertTokenScans(t, "+", token.Add)
}

func Test_17_ScanAll(t *testing.T) {
	assertTokenScans(t, "-", token.Sub)
}

func Test_18_ScanAll(t *testing.T) {
	assertTokenScans(t, "*", token.Mul)
}

func Test_19_ScanAll(t *testing.T) {
	assertTokenScans(t, "/", token.Div)
}

func Test_20_ScanAll(t *testing.T) {
	assertTokenScans(t, "%", token.Mod)
}

func Test_21_ScanAll(t *testing.T) {
	assertTokenScans(t, "<", token.LT)
}

func Test_22_ScanAll(t *testing.T) {
	assertTokenScans(t, ">", token.GT)
}

func Test_23_ScanAll(t *testing.T) {
	assertTokenScans(t, "<=", token.LTE)
}

func Test_24_ScanAll(t *testing.T) {
	assertTokenScans(t, ">=", token.GTE)
}

func Test_25_ScanAll(t *testing.T) {
	assertTokenScans(t, "==", token.EQU)
}

func Test_26_ScanAll(t *testing.T) {
	assertTokenScans(t, "!=", token.NEQ)
}

func Test_27_ScanAll(t *testing.T) {
	assertTokenScans(t, "(", token.ParenOpen)
}

func Test_28_ScanAll(t *testing.T) {
	assertTokenScans(t, ")", token.ParenClose)
}

func Test_29_ScanAll(t *testing.T) {
	assertTokenScans(t, "{", token.BraceOpen)
}

func Test_30_ScanAll(t *testing.T) {
	assertTokenScans(t, "}", token.BraceClose)
}

func Test_31_ScanAll(t *testing.T) {
	assertTokenScans(t, "[", token.BracketOpen)
}

func Test_32_ScanAll(t *testing.T) {
	assertTokenScans(t, "]", token.BracketClose)
}

func Test_33_ScanAll(t *testing.T) {
	assertTokenScans(t, `""`, token.String)
}

func Test_34_ScanAll(t *testing.T) {
	assertTokenScans(t, `"a"`, token.String)
}

func Test_35_ScanAll(t *testing.T) {
	assertTokenScans(t, `"abc"`, token.String)
}

func Test_36_ScanAll(t *testing.T) {
	assertTokenScans(t, `"   "`, token.String)
}

func Test_37_ScanAll(t *testing.T) {
	assertTokenScans(t, `"\\"`, token.String)
}

func Test_38_ScanAll(t *testing.T) {
	assertTokenScans(t, `"\\\\\\"`, token.String)
}

func Test_39_ScanAll(t *testing.T) {
	assertTokenScans(t, `"\"\"\""`, token.String)
}

func Test_40_ScanAll(t *testing.T) {
	assertTokenScans(t, " ", token.Space)
}

func Test_41_ScanAll(t *testing.T) {
	assertTokenScans(t, "\t", token.Space)
}

func Test_42_ScanAll(t *testing.T) {
	assertTokenScans(t, "\v", token.Space)
}

func Test_43_ScanAll(t *testing.T) {
	assertTokenScans(t, "\r", token.Space)
}

func Test_44_ScanAll(t *testing.T) {
	assertTokenScans(t, "\f", token.Space)
}

func Test_45_ScanAll(t *testing.T) {
	assertTokenScans(t, "  \t\v \f\r   \v\v\t", token.Space)
}

func Test_50_ScanAll(t *testing.T) {
	assertTokenScans(t, "0", token.Number)
}

func Test_51_ScanAll(t *testing.T) {
	assertTokenScans(t, "0.00000", token.Number)
}

func Test_52_ScanAll(t *testing.T) {
	assertTokenScans(t, "0.1", token.Number)
}

func Test_53_ScanAll(t *testing.T) {
	assertTokenScans(t, "1", token.Number)
}

func Test_54_ScanAll(t *testing.T) {
	assertTokenScans(t, "1.1234567890", token.Number)
}

func Test_55_ScanAll(t *testing.T) {
	assertTokenScans(t, "123456789.987654321", token.Number)
}

func Test_56_ScanAll(t *testing.T) {
	assertTokenScans(t, "9", token.Number)
}

func Test_60_ScanAll(t *testing.T) {
	assertTokenScans(t, "abc", token.Var)
}

func Test_61_ScanAll(t *testing.T) {
	assertTokenScans(t, "abc_xyz", token.Var)
}

func Test_62_ScanAll(t *testing.T) {
	assertTokenScans(t, "forest", token.Var)
}

func Test_63_ScanAll(t *testing.T) {
	assertTokenScans(t, "For", token.Var)
}

func Test_64_ScanAll(t *testing.T) {
	assertTokenScans(t, "FOR", token.Var)
}

func Test_65_ScanAll(t *testing.T) {
	assertTokenScans(t, "e", token.Var)
}

func Test_70_ScanAll(t *testing.T) {
	assertTokenScans(t, "if", token.If)
}

func Test_71_ScanAll(t *testing.T) {
	assertTokenScans(t, "for", token.For)
}

func Test_72_ScanAll(t *testing.T) {
	assertTokenScans(t, "in", token.In)
}

func Test_73_ScanAll(t *testing.T) {
	assertTokenScans(t, "if", token.If)
}

func Test_74_ScanAll(t *testing.T) {
	assertTokenScans(t, "watch", token.Watch)
}

func Test_75_ScanAll(t *testing.T) {
	assertTokenScans(t, "when", token.When)
}

func Test_76_ScanAll(t *testing.T) {
	assertTokenScans(t, "is", token.Is)
}

func Test_77_ScanAll(t *testing.T) {
	assertTokenScans(t, "E", token.E)
}

func Test_78_ScanAll(t *testing.T) {
	assertTokenScans(t, "F", token.F)
}

func Test_79_ScanAll(t *testing.T) {
	assertTokenScans(t, "true", token.True)
}

func Test_80_ScanAll(t *testing.T) {
	assertTokenScans(t, "false", token.False)
}

func Test_81_ScanAll(t *testing.T) {
	assertTokenScans(t, `"\\"`, token.String)
}

func Test_100_ScanAll(t *testing.T) {
	assertScanError(t, "~", ErrUnknownSymbol)
}

func Test_101_ScanAll(t *testing.T) {
	assertScanError(t, `"`, ErrUnterminatedString)
}

func Test_102_ScanAll(t *testing.T) {
	assertScanError(t, `"""`, ErrUnterminatedString)
}

func Test_103_ScanAll(t *testing.T) {
	assertScanError(t, `"\`, ErrUnterminatedString)
}

func Test_104_ScanAll(t *testing.T) {
	assertScanError(t, `"\"`, ErrUnterminatedString)
}

func Test_105_ScanAll(t *testing.T) {
	assertScanError(t, `"\\\"`, ErrUnterminatedString)
}

func Test_106_ScanAll(t *testing.T) {
	assertScanError(t, "=!", ErrUnknownSymbol)
}

func Test_107_ScanAll(t *testing.T) {
	assertScanError(t, ".", ErrUnknownSymbol)
}

func Test_108_ScanAll(t *testing.T) {
	assertScanError(t, "0.", ErrMissingFractional)
}

func Test_109_ScanAll(t *testing.T) {
	assertScanError(t, "0.a", ErrMissingFractional)
}

func Test_200_ScanAll(t *testing.T) {
	given := "x = 1"

	gen := token.NewTokenGenerator()
	exp := []token.Token{
		gen(token.Var, "x"),
		gen(token.Space, " "),
		gen(token.Ass, "="),
		gen(token.Space, " "),
		gen(token.Number, "1"),
	}

	assertAllTokensScan(t, given, exp)
}

func Test_201_ScanAll(t *testing.T) {
	given := strings.Join([]string{
		`x = true`,
		`y, z = 123.456, "string"`,
		``,
		`f := F(a, b) c, d`,
		`	when a`,
		`		is 1: @println("one")`,
		`		a == b: @println("a == b")`,
		`		true: @println("meh")`,
		`	end`,
		`end`,
		``,
	}, "\n")

	gen := token.NewTokenGenerator()
	exp := []token.Token{
		// `x = true`
		gen(token.Var, "x"),
		gen(token.Space, " "),
		gen(token.Ass, "="),
		gen(token.Space, " "),
		gen(token.True, "true"),
		gen(token.Newline, "\n"),
		// `y, z = 1, "string"`
		gen(token.Var, "y"),
		gen(token.Comma, ","),
		gen(token.Space, " "),
		gen(token.Var, "z"),
		gen(token.Space, " "),
		gen(token.Ass, "="),
		gen(token.Space, " "),
		gen(token.Number, "123.456"),
		gen(token.Comma, ","),
		gen(token.Space, " "),
		gen(token.String, `"string"`),
		gen(token.Newline, "\n"),
		// ``
		gen(token.Newline, "\n"),
		// `f := F(a, b) c, d {`
		gen(token.Var, "f"),
		gen(token.Space, " "),
		gen(token.Def, ":="),
		gen(token.Space, " "),
		gen(token.F, "F"),
		gen(token.ParenOpen, "("),
		gen(token.Var, "a"),
		gen(token.Comma, ","),
		gen(token.Space, " "),
		gen(token.Var, "b"),
		gen(token.ParenClose, ")"),
		gen(token.Space, " "),
		gen(token.Var, "c"),
		gen(token.Comma, ","),
		gen(token.Space, " "),
		gen(token.Var, "d"),
		gen(token.Newline, "\n"),
		// `	when a`
		gen(token.Space, "\t"),
		gen(token.When, "when"),
		gen(token.Space, " "),
		gen(token.Var, "a"),
		gen(token.Newline, "\n"),
		// ` 	is 1: @println("one")`
		gen(token.Space, "\t\t"),
		gen(token.Is, "is"),
		gen(token.Space, " "),
		gen(token.Number, "1"),
		gen(token.Colon, ":"),
		gen(token.Space, " "),
		gen(token.Spell, "@"),
		gen(token.Var, "println"),
		gen(token.ParenOpen, "("),
		gen(token.String, `"one"`),
		gen(token.ParenClose, ")"),
		gen(token.Newline, "\n"),
		// `		a == b: @println("b")`,
		gen(token.Space, "\t\t"),
		gen(token.Var, "a"),
		gen(token.Space, " "),
		gen(token.EQU, "=="),
		gen(token.Space, " "),
		gen(token.Var, "b"),
		gen(token.Colon, ":"),
		gen(token.Space, " "),
		gen(token.Spell, "@"),
		gen(token.Var, "println"),
		gen(token.ParenOpen, "("),
		gen(token.String, `"a == b"`),
		gen(token.ParenClose, ")"),
		gen(token.Newline, "\n"),
		// `		true: @println("meh")`
		gen(token.Space, "\t\t"),
		gen(token.True, "true"),
		gen(token.Colon, ":"),
		gen(token.Space, " "),
		gen(token.Spell, "@"),
		gen(token.Var, "println"),
		gen(token.ParenOpen, "("),
		gen(token.String, `"meh"`),
		gen(token.ParenClose, ")"),
		gen(token.Newline, "\n"),
		// `	}`
		gen(token.Space, "\t"),
		gen(token.End, "end"),
		gen(token.Newline, "\n"),
		// `}`
		gen(token.End, "end"),
		gen(token.Newline, "\n"),
	}

	assertAllTokensScan(t, given, exp)
}
