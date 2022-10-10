package tokentest

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/utilities/pos"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

func Test_1_TokenGenerator(t *testing.T) {
	gen := NewTokenGenerator()

	act := gen(token.EQU, "==")
	exp := token.Token{
		TokenType: token.EQU,
		Value:     "==",
		Range:     pos.RawRangeForString(0, 0, 0, "=="),
	}

	require.Equal(t, exp, act)
}

func Test_2_TokenGenerator(t *testing.T) {
	gen := NewTokenGenerator()

	_ = gen(token.EQU, "==")
	act := gen(token.Newline, "\n")

	exp := token.Token{
		TokenType: token.Newline,
		Value:     "\n",
		Range:     pos.RawRangeForString(2, 0, 2, "\n"),
	}

	require.Equal(t, exp, act)
}

func Test_3_TokenGenerator(t *testing.T) {
	gen := NewTokenGenerator()

	_ = gen(token.EQU, "==")
	_ = gen(token.NEQ, "!=")
	act := gen(token.LTE, "<=")

	exp := token.Token{
		TokenType: token.LTE,
		Value:     "<=",
		Range:     pos.RawRangeForString(4, 0, 4, "<="),
	}

	require.Equal(t, exp, act)
}
