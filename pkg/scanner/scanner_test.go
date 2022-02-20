package scanner

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func given(in string) RuneReader {
	return token.NewRuneReader([]rune(in))
}

func lex(tk token.Token, v string) token.Lexeme {
	return token.Lexeme{
		Token: tk,
		Value: v,
	}
}

func TestScanAll_0(t *testing.T) {
	r := given("")

	act, e := ScanAll(r)

	require.Nil(t, e, "%+v", e)

	var exp []token.Lexeme
	require.Equal(t, exp, act)
}

func TestScanAll_bool_1(t *testing.T) {
	r := given("true")

	act, e := ScanAll(r)

	require.Nil(t, e, "%+v", e)

	exp := []token.Lexeme{
		lex(token.TK_BOOL, "true"),
	}
	require.Equal(t, exp, act)
}

func TestScanAll_bool_2(t *testing.T) {
	r := given("false")

	act, e := ScanAll(r)

	require.Nil(t, e, "%+v", e)

	exp := []token.Lexeme{
		lex(token.TK_BOOL, "false"),
	}
	require.Equal(t, exp, act)
}
