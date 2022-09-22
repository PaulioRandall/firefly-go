package scanner

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/err"
	"github.com/PaulioRandall/firefly-go/pkg/readers"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func singletonTokenList(tt token.TokenType, val string, valLen int) []token.Token {
	return []token.Token{
		token.MakeToken(
			tt,
			val,
			token.MakeInlineRange(0, 0, 0, valLen),
		),
	}
}

func doScanAllTest(t *testing.T, given string, exp []token.Token) {
	r := readers.NewRuneStringReader(given)

	act, e := ScanAll(r)
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func doScanTokenTest(t *testing.T, given string, exp token.TokenType) {
	r := readers.NewRuneStringReader(given)

	actTk, e := ScanAll(r)
	expTk := singletonTokenList(exp, given, len(given))

	require.Nil(t, e, "Expected %q but got %+v", exp.String(), err.DebugString(e))
	require.NotEmpty(t, actTk)
	require.Equal(t, expTk, actTk,
		"Expected %q but got %q", exp.String(), actTk[0].Type.String(),
	)
}

func Test_1_ScanAll(t *testing.T) {
	r := readers.NewRuneStringReader("")

	act, e := ScanAll(r)
	var exp []token.Token

	require.Nil(t, e)
	require.Equal(t, exp, act)
}

func Test_2_ScanAll(t *testing.T) {
	r := readers.NewRuneStringReader("~")
	_, e := ScanAll(r)
	require.NotNil(t, e)
}

func Test_3_ScanAll(t *testing.T) {
	words := map[string]token.TokenType{
		"if":    token.If,
		"for":   token.For,
		"in":    token.In,
		"watch": token.Watch,
		"when":  token.When,
		"is":    token.Is,
		"E":     token.E,
		"F":     token.F,
		"end":   token.End,
		"true":  token.True,
		"false": token.False,
	}

	for given, tt := range words {
		exp := singletonTokenList(tt, given, len(given))
		doScanAllTest(t, given, exp)
	}
}

func Test_4_ScanAll(t *testing.T) {
	vars := []string{
		"abc",
		"abc_xyz",
		"forest",
		"For",
		"FOR",
		"e",
	}

	for _, given := range vars {
		exp := singletonTokenList(token.Var, given, len(given))
		doScanAllTest(t, given, exp)
	}
}

func Test_6_ScanAll(t *testing.T) {
	vars := []string{
		"0",
		"0.00000",
		"0.1",
		"1",
		"1.1234567890",
		"123456789.987654321",
		"9",
	}

	for _, given := range vars {
		exp := singletonTokenList(token.Number, given, len(given))
		doScanAllTest(t, given, exp)
	}
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
	doScanAllTest(t, given, exp)
}

func Test_8_ScanAll(t *testing.T) {
	vars := []string{
		" ",
		"\t",
		"\f",
		"\v",
		"\r",
		"  \t\v \f\r   \v\v\t",
	}

	for _, given := range vars {
		exp := singletonTokenList(token.Space, given, len(given))
		doScanAllTest(t, given, exp)
	}
}

func Test_9_ScanAll(t *testing.T) {
	vars := []string{
		`""`,
		`"a"`,
		`"abc"`,
		`"   "`,
		`"\\"`,
		`"\\\\\\"`,
		`"\"\"\""`,
	}

	for _, given := range vars {
		exp := singletonTokenList(token.String, given, len(given))
		doScanAllTest(t, given, exp)
	}
}

func Test_10_ScanAll(t *testing.T) {
	doScanTokenTest(t, "=", token.Ass)
}

func Test_11_ScanAll(t *testing.T) {
	doScanTokenTest(t, ":=", token.Def)
}

func Test_12_ScanAll(t *testing.T) {
	doScanTokenTest(t, ";", token.Terminator)
}

func Test_13_ScanAll(t *testing.T) {
	doScanTokenTest(t, ",", token.Comma)
}

func Test_14_ScanAll(t *testing.T) {
	doScanTokenTest(t, ":", token.Colon)
}

func Test_15_ScanAll(t *testing.T) {
	doScanTokenTest(t, "@", token.Spell)
}

func Test_16_ScanAll(t *testing.T) {
	doScanTokenTest(t, "+", token.Add)
}

func Test_17_ScanAll(t *testing.T) {
	doScanTokenTest(t, "-", token.Sub)
}

func Test_18_ScanAll(t *testing.T) {
	doScanTokenTest(t, "*", token.Mul)
}

func Test_19_ScanAll(t *testing.T) {
	doScanTokenTest(t, "/", token.Div)
}

func Test_20_ScanAll(t *testing.T) {
	doScanTokenTest(t, "%", token.Mod)
}

func Test_21_ScanAll(t *testing.T) {
	doScanTokenTest(t, "<", token.LT)
}

func Test_22_ScanAll(t *testing.T) {
	doScanTokenTest(t, ">", token.GT)
}

func Test_23_ScanAll(t *testing.T) {
	doScanTokenTest(t, "<=", token.LTE)
}

func Test_24_ScanAll(t *testing.T) {
	doScanTokenTest(t, ">=", token.GTE)
}

func Test_25_ScanAll(t *testing.T) {
	doScanTokenTest(t, "==", token.EQU)
}

func Test_26_ScanAll(t *testing.T) {
	doScanTokenTest(t, "!=", token.NEQ)
}

func Test_27_ScanAll(t *testing.T) {
	doScanTokenTest(t, "(", token.ParenOpen)
}

func Test_28_ScanAll(t *testing.T) {
	doScanTokenTest(t, ")", token.ParenClose)
}

func Test_29_ScanAll(t *testing.T) {
	doScanTokenTest(t, "{", token.BraceOpen)
}

func Test_30_ScanAll(t *testing.T) {
	doScanTokenTest(t, "}", token.BraceClose)
}

func Test_31_ScanAll(t *testing.T) {
	doScanTokenTest(t, "[", token.BracketOpen)
}

func Test_32_ScanAll(t *testing.T) {
	doScanTokenTest(t, "]", token.BracketClose)
}
