package scanner

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/err"
	"github.com/PaulioRandall/firefly-go/pkg/readers"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func singletonTokenList(tt token.TokenType, v string, vLen int) []token.Token {
	return []token.Token{
		token.MakeToken(
			tt,
			v,
			token.MakeInlineRange(0, 0, 0, vLen),
		),
	}
}

func checkAllTokensScan(t *testing.T, given string, exp []token.Token) {
	r := readers.NewRuneStringReader(given)

	act, e := ScanAll(r)
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func checkTokenScans(t *testing.T, given string, exp token.TokenType) {
	r := readers.NewRuneStringReader(given)

	actTk, e := ScanAll(r)
	expTk := singletonTokenList(exp, given, len(given))

	require.Nil(t, e, "Expected %q but got %+v", exp.String(), err.DebugString(e))
	require.NotEmpty(t, actTk)
	require.Equal(t, expTk, actTk,
		"Expected %q but got %q", exp.String(), actTk[0].Type.String(),
	)
}

func checkForScanError(t *testing.T, given string, exp error) {
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
	checkAllTokensScan(t, given, exp)
}

func Test_10_ScanAll(t *testing.T) {
	checkTokenScans(t, "=", token.Ass)
}

func Test_11_ScanAll(t *testing.T) {
	checkTokenScans(t, ":=", token.Def)
}

func Test_12_ScanAll(t *testing.T) {
	checkTokenScans(t, ";", token.Terminator)
}

func Test_13_ScanAll(t *testing.T) {
	checkTokenScans(t, ",", token.Comma)
}

func Test_14_ScanAll(t *testing.T) {
	checkTokenScans(t, ":", token.Colon)
}

func Test_15_ScanAll(t *testing.T) {
	checkTokenScans(t, "@", token.Spell)
}

func Test_16_ScanAll(t *testing.T) {
	checkTokenScans(t, "+", token.Add)
}

func Test_17_ScanAll(t *testing.T) {
	checkTokenScans(t, "-", token.Sub)
}

func Test_18_ScanAll(t *testing.T) {
	checkTokenScans(t, "*", token.Mul)
}

func Test_19_ScanAll(t *testing.T) {
	checkTokenScans(t, "/", token.Div)
}

func Test_20_ScanAll(t *testing.T) {
	checkTokenScans(t, "%", token.Mod)
}

func Test_21_ScanAll(t *testing.T) {
	checkTokenScans(t, "<", token.LT)
}

func Test_22_ScanAll(t *testing.T) {
	checkTokenScans(t, ">", token.GT)
}

func Test_23_ScanAll(t *testing.T) {
	checkTokenScans(t, "<=", token.LTE)
}

func Test_24_ScanAll(t *testing.T) {
	checkTokenScans(t, ">=", token.GTE)
}

func Test_25_ScanAll(t *testing.T) {
	checkTokenScans(t, "==", token.EQU)
}

func Test_26_ScanAll(t *testing.T) {
	checkTokenScans(t, "!=", token.NEQ)
}

func Test_27_ScanAll(t *testing.T) {
	checkTokenScans(t, "(", token.ParenOpen)
}

func Test_28_ScanAll(t *testing.T) {
	checkTokenScans(t, ")", token.ParenClose)
}

func Test_29_ScanAll(t *testing.T) {
	checkTokenScans(t, "{", token.BraceOpen)
}

func Test_30_ScanAll(t *testing.T) {
	checkTokenScans(t, "}", token.BraceClose)
}

func Test_31_ScanAll(t *testing.T) {
	checkTokenScans(t, "[", token.BracketOpen)
}

func Test_32_ScanAll(t *testing.T) {
	checkTokenScans(t, "]", token.BracketClose)
}

func Test_33_ScanAll(t *testing.T) {
	checkTokenScans(t, `""`, token.String)
}

func Test_34_ScanAll(t *testing.T) {
	checkTokenScans(t, `"a"`, token.String)
}

func Test_35_ScanAll(t *testing.T) {
	checkTokenScans(t, `"abc"`, token.String)
}

func Test_36_ScanAll(t *testing.T) {
	checkTokenScans(t, `"   "`, token.String)
}

func Test_37_ScanAll(t *testing.T) {
	checkTokenScans(t, `"\\"`, token.String)
}

func Test_38_ScanAll(t *testing.T) {
	checkTokenScans(t, `"\\\\\\"`, token.String)
}

func Test_39_ScanAll(t *testing.T) {
	checkTokenScans(t, `"\"\"\""`, token.String)
}

func Test_40_ScanAll(t *testing.T) {
	checkTokenScans(t, " ", token.Space)
}

func Test_41_ScanAll(t *testing.T) {
	checkTokenScans(t, "\t", token.Space)
}

func Test_42_ScanAll(t *testing.T) {
	checkTokenScans(t, "\v", token.Space)
}

func Test_43_ScanAll(t *testing.T) {
	checkTokenScans(t, "\r", token.Space)
}

func Test_44_ScanAll(t *testing.T) {
	checkTokenScans(t, "\f", token.Space)
}

func Test_45_ScanAll(t *testing.T) {
	checkTokenScans(t, "  \t\v \f\r   \v\v\t", token.Space)
}

func Test_50_ScanAll(t *testing.T) {
	checkTokenScans(t, "0", token.Number)
}

func Test_51_ScanAll(t *testing.T) {
	checkTokenScans(t, "0.00000", token.Number)
}

func Test_52_ScanAll(t *testing.T) {
	checkTokenScans(t, "0.1", token.Number)
}

func Test_53_ScanAll(t *testing.T) {
	checkTokenScans(t, "1", token.Number)
}

func Test_54_ScanAll(t *testing.T) {
	checkTokenScans(t, "1.1234567890", token.Number)
}

func Test_55_ScanAll(t *testing.T) {
	checkTokenScans(t, "123456789.987654321", token.Number)
}

func Test_56_ScanAll(t *testing.T) {
	checkTokenScans(t, "9", token.Number)
}

func Test_60_ScanAll(t *testing.T) {
	checkTokenScans(t, "abc", token.Var)
}

func Test_61_ScanAll(t *testing.T) {
	checkTokenScans(t, "abc_xyz", token.Var)
}

func Test_62_ScanAll(t *testing.T) {
	checkTokenScans(t, "forest", token.Var)
}

func Test_63_ScanAll(t *testing.T) {
	checkTokenScans(t, "For", token.Var)
}

func Test_64_ScanAll(t *testing.T) {
	checkTokenScans(t, "FOR", token.Var)
}

func Test_65_ScanAll(t *testing.T) {
	checkTokenScans(t, "e", token.Var)
}

func Test_70_ScanAll(t *testing.T) {
	checkTokenScans(t, "if", token.If)
}

func Test_71_ScanAll(t *testing.T) {
	checkTokenScans(t, "for", token.For)
}

func Test_72_ScanAll(t *testing.T) {
	checkTokenScans(t, "in", token.In)
}

func Test_73_ScanAll(t *testing.T) {
	checkTokenScans(t, "if", token.If)
}

func Test_74_ScanAll(t *testing.T) {
	checkTokenScans(t, "watch", token.Watch)
}

func Test_75_ScanAll(t *testing.T) {
	checkTokenScans(t, "when", token.When)
}

func Test_76_ScanAll(t *testing.T) {
	checkTokenScans(t, "is", token.Is)
}

func Test_77_ScanAll(t *testing.T) {
	checkTokenScans(t, "E", token.E)
}

func Test_78_ScanAll(t *testing.T) {
	checkTokenScans(t, "F", token.F)
}

func Test_79_ScanAll(t *testing.T) {
	checkTokenScans(t, "true", token.True)
}

func Test_80_ScanAll(t *testing.T) {
	checkTokenScans(t, "false", token.False)
}

func Test_100_ScanAll(t *testing.T) {
	checkForScanError(t, `"`, err.EOF)
}

func Test_101_ScanAll(t *testing.T) {
	checkForScanError(t, "~", ErrUnknownSymbol)
}
