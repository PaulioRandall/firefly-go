package tokenreader

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/pos"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

func tok(tt token.TokenType) token.Token {
	return token.MakeToken(tt, "", pos.Range{})
}

func Test_1_tokenReader_More(t *testing.T) {
	given := FromList()
	require.False(t, given.More())
}

func Test_2_tokenReader_Peek(t *testing.T) {
	given := FromList(
		tok(token.Var),
	)

	exp := tok(token.Var)
	act := given.Peek()

	require.Equal(t, exp, act)
	require.True(t, given.More())
}

func Test_3_tokenReader_Read(t *testing.T) {
	given := FromList(
		tok(token.Var),
	)

	exp := tok(token.Var)
	act := given.Read()

	require.Equal(t, exp, act)
	require.False(t, given.More())
}
