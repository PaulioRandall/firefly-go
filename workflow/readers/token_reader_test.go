package readers

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/workflow/token"
	"github.com/stretchr/testify/require"
)

func tok(tt token.TokenType, v string) token.Token {
	return token.MakeToken(tt, v, token.MakeInlineRange(0, 0, 0, len(v)))
}

func Test_1_tokenReader_More(t *testing.T) {
	given := NewTokenReader([]token.Token{})
	require.False(t, given.More())
}

func Test_2_tokenReader_Peek(t *testing.T) {
	given := NewTokenReader([]token.Token{
		tok(token.Var, ""),
	})

	exp := tok(token.Var, "")
	act := given.Peek()

	require.Equal(t, exp, act)
	require.True(t, given.More())
}

func Test_3_tokenReader_Read(t *testing.T) {
	given := NewTokenReader([]token.Token{
		tok(token.Var, ""),
	})

	exp := tok(token.Var, "")
	act := given.Read()

	require.Equal(t, exp, act)
	require.False(t, given.More())
}
