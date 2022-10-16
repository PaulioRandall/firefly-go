package auditor

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func tok(tt token.TokenType, v string) token.Token {
	return tokentest.Tok(tt, v)
}

func aud(given ...token.Token) *Auditor {
	return NewAuditor(inout.NewListReader(given))
}

func Test_1(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
	}

	a := aud(given...)
	tk := a.Peek()

	require.Equal(t, given[0], tk)
	require.True(t, a.More())
}

func Test_2(t *testing.T) {
	a := aud()

	require.Panics(t, func() {
		_ = a.Peek()
	})
}

func Test_3(t *testing.T) {
	a := aud()

	require.Panics(t, func() {
		_ = a.Read()
	})
}

func Test_4(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
	}

	a := aud(given...)
	tk := a.Read()

	require.Equal(t, given[0], tk)
	require.False(t, a.More())
}
