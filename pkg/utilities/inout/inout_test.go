package inout

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func tok(tt token.TokenType, v string) token.Token {
	return tokentest.Tok(tt, v)
}

func requireError(t *testing.T, exp, act error, msg string) {
	require.True(t, err.Is(act, exp), msg)
}
