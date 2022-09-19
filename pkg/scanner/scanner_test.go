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
	words := map[string]token.TokenType{
		"if":      token.If,
		"for":     token.For,
		"watch":   token.Watch,
		"when":    token.When,
		"E":       token.E,
		"F":       token.F,
		"end":     token.End,
		"true":    token.True,
		"false":   token.False,
		"abc":     token.Var,
		"abc_xyz": token.Var,
		"forest":  token.Var,
		"For":     token.Var,
		"FOR":     token.Var,
		"e":       token.Var,
	}

	for given, tt := range words {
		exp := singletonTokenList(tt, given, len(given))
		doScanAllTest(t, given, exp)
	}
}

func Test_ScanAll_4(t *testing.T) {
	operators := map[string]token.TokenType{
		"=":  token.Ass,
		":=": token.Def,
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
		";":  token.Terminator,
	}

	for given, tt := range operators {
		exp := singletonTokenList(tt, given, len(given))
		doScanAllTest(t, given, exp)
	}
}
