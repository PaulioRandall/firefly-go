package inout

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func tok(tt token.TokenType, v string) token.Token {
	return tokentest.Tok(tt, v)
}

func newPPR(given ...token.Token) PanicPosReader[token.Token] {
	return NewPanicPosReader[token.Token](NewListReader[token.Token](given))
}

// TODO: Test Putback & Prev funcs

func Test_1_panicPosReader(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
	}

	ppr := newPPR(given...)
	tk := ppr.Peek()

	require.Equal(t, given[0], tk)
	require.True(t, ppr.More())
}

func Test_2v(t *testing.T) {
	ppr := newPPR()

	require.Panics(t, func() {
		_ = ppr.Peek()
	})
}

func Test_3_panicPosReader(t *testing.T) {
	ppr := newPPR()

	require.Panics(t, func() {
		_ = ppr.Read()
	})
}

func Test_4_panicPosReader(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
	}

	ppr := newPPR(given...)
	tk := ppr.Read()

	require.Equal(t, given[0], tk)
	require.False(t, ppr.More())
}
