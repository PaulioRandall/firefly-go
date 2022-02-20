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

	var exp []token.Lexeme
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}
