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

func Test_5_ScanAll(t *testing.T) {
	operators := map[string]token.TokenType{
		"=":  token.Ass,
		":=": token.Def,
		";":  token.Terminator,
		",":  token.Comma,
		":":  token.Colon,
		"@":  token.Spell,
		"+":  token.Add,
		"-":  token.Sub,
		"*":  token.Mul,
		"/":  token.Div,
		"%":  token.Mod,
		"<":  token.LT,
		">":  token.GT,
		"<=": token.LTE,
		">=": token.GTE,
		"==": token.EQU,
		"!=": token.NEQ,
		"(":  token.ParenOpen,
		")":  token.ParenClose,
		"{":  token.BraceOpen,
		"}":  token.BraceClose,
		"[":  token.BracketOpen,
		"]":  token.BracketClose,
	}

	for given, tt := range operators {
		exp := singletonTokenList(tt, given, len(given))
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
