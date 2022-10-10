package tokentest

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/utilities/pos"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

func Tok(tt token.TokenType, v string) token.Token {
	return token.MakeToken(tt, v, pos.RawRangeForString(0, 0, 0, v))
}

type TokenGenerator func(token.TokenType, string) token.Token

func NewTokenGenerator() TokenGenerator {
	prev := pos.Range{}

	return func(tt token.TokenType, v string) token.Token {
		prev.ShiftString(v)
		return token.MakeToken(tt, v, prev)
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
