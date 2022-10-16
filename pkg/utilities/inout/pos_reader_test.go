package inout

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func newPR(given ...token.Token) PosReader[token.Token] {
	lr := NewListReader[token.Token](given)
	br := NewBufReader[token.Token](lr)
	return NewPosReader[token.Token](br)
}

func Test_1_posReader(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
	}

	ppr := newPR(given...)
	tk := ppr.Peek()

	require.Equal(t, given[0], tk)
	require.True(t, ppr.More())
}

func Test_2_posReader(t *testing.T) {
	ppr := newPR()

	require.Panics(t, func() {
		_ = ppr.Peek()
	})
}

func Test_3_posReader(t *testing.T) {
	ppr := newPR()

	require.Panics(t, func() {
		_ = ppr.Read()
	})
}

func Test_4_posReader(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
	}

	ppr := newPR(given...)
	tk := ppr.Read()

	require.Equal(t, given[0], tk)
	require.False(t, ppr.More())
}
