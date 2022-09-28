package token

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_1_TokenGenerator(t *testing.T) {
	gen := NewTokenGenerator()

	act := gen(EQU, "==")
	exp := Token{
		TokenType: EQU,
		Value:     "==",
		Range:     MakeInlineRange(0, 0, 0, len("==")),
	}

	require.Equal(t, exp, act)
}

func Test_2_TokenGenerator(t *testing.T) {
	gen := NewTokenGenerator()

	_ = gen(EQU, "==")
	act := gen(Newline, "\n")

	exp := Token{
		TokenType: Newline,
		Value:     "\n",
		Range: MakeRange(
			MakePos(2, 0, 2),
			MakePos(3, 1, 0),
		),
	}

	require.Equal(t, exp, act)
}

func Test_3_TokenGenerator(t *testing.T) {
	gen := NewTokenGenerator()

	_ = gen(EQU, "==")
	_ = gen(NEQ, "!=")
	act := gen(LTE, "<=")

	exp := Token{
		TokenType: LTE,
		Value:     "<=",
		Range:     MakeInlineRange(4, 0, 4, len("<=")),
	}

	require.Equal(t, exp, act)
}
