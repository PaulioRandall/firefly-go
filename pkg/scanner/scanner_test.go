package scanner

import (
	"testing"

	"github.com/stretchr/testify/require"

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

func Test_ScanAll_1(t *testing.T) {
	r := readers.NewRuneStringReader("")

	act, e := ScanAll(r)
	var exp []token.Token

	require.Nil(t, e)
	require.Equal(t, exp, act)
}

func Test_ScanAll_2(t *testing.T) {
	r := readers.NewRuneStringReader("~")
	_, e := ScanAll(r)
	require.NotNil(t, e)
}

func Test_ScanAll_3(t *testing.T) {
	given := "if"
	exp := singletonTokenList(token.If, given, len(given))
	doScanAllTest(t, given, exp)
}

func Test_ScanAll_4(t *testing.T) {
	given := "for"
	exp := singletonTokenList(token.For, given, len(given))
	doScanAllTest(t, given, exp)
}

func Test_ScanAll_5(t *testing.T) {
	given := "watch"
	exp := singletonTokenList(token.Watch, given, len(given))
	doScanAllTest(t, given, exp)
}

func Test_ScanAll_6(t *testing.T) {
	given := "when"
	exp := singletonTokenList(token.When, given, len(given))
	doScanAllTest(t, given, exp)
}

func Test_ScanAll_7(t *testing.T) {
	given := "E"
	exp := singletonTokenList(token.E, given, len(given))
	doScanAllTest(t, given, exp)
}

func Test_ScanAll_8(t *testing.T) {
	given := "F"
	exp := singletonTokenList(token.F, given, len(given))
	doScanAllTest(t, given, exp)
}

func Test_ScanAll_9(t *testing.T) {
	given := "end"
	exp := singletonTokenList(token.End, given, len(given))
	doScanAllTest(t, given, exp)
}

func Test_ScanAll_10(t *testing.T) {
	given := "abc"
	exp := singletonTokenList(token.Var, given, len(given))
	doScanAllTest(t, given, exp)
}

func Test_ScanAll_11(t *testing.T) {
	given := "abc_xyz"
	exp := singletonTokenList(token.Var, given, len(given))
	doScanAllTest(t, given, exp)
}

func Test_ScanAll_12(t *testing.T) {
	given := "_"
	exp := singletonTokenList(token.Var, given, len(given))
	doScanAllTest(t, given, exp)
}

func Test_ScanAll_13(t *testing.T) {
	given := "forest"
	exp := singletonTokenList(token.Var, given, len(given))
	doScanAllTest(t, given, exp)
}

func Test_ScanAll_14(t *testing.T) {
	set := map[token.TokenType]string{
		token.Add: "+",
		token.Sub: "-",
		token.Mul: "*",
		token.Div: "/",
		token.Mod: "%",
		token.LT:  "<",
		token.GT:  ">",
		token.LTE: "<=",
		token.GTE: ">=",
		token.EQU: "==",
		token.NEQ: "!=",
		token.ASS: "=",
		token.DEF: ":=",
	}

	for tt, given := range set {
		exp := singletonTokenList(tt, given, len(given))
		doScanAllTest(t, given, exp)
	}
}
