package tokentest

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func Tok(tt token.TokenType, v string) token.Token {
	return token.MakeTokenAt(tt, v, pos.Pos{})
}

type TokenGenerator func(token.TokenType, string) token.Token

func NewTokenGenerator() TokenGenerator {
	var from, to pos.Pos

	return func(tt token.TokenType, v string) token.Token {
		from, to = pos.RangeFor(to, v)
		return token.MakeToken(tt, v, from, to)
	}
}

func RequireEqual(t *testing.T, exp, act []token.Token) {
	failMsg := "Failed to match expected tokens"

	for i, expTk := range exp {
		if len(act) == i {
			require.Failf(t, failMsg, "At index %d\n\texpected %v\n\tbut no more tokens", i, expTk.Debug())
		}

		actTk := act[i]
		require.Equal(t, expTk, actTk, "At index %d\n\texpected %v\n\tbut got %v", i, expTk.Debug(), actTk.Debug())
	}

	if len(act) > len(exp) {
		i := len(exp)
		require.Failf(t, failMsg, "Required tokens checked but at index %d\n\treceived unexpected %v", i, act[i].Debug())
	}
}
