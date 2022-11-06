package tokentest

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func Test_1_TokenGenerator(t *testing.T) {
	gen := NewTokenGenerator()

	act := gen(token.Equ, "==")

	exp := token.MakeTokenAt(token.Equ, "==", pos.Pos{})

	require.Equal(t, exp, act)
}

func Test_2_TokenGenerator(t *testing.T) {
	gen := NewTokenGenerator()

	_ = gen(token.Equ, "==")
	act := gen(token.Newline, "\n")

	exp := token.MakeTokenAt(token.Newline, "\n", pos.At(2, 0, 2))

	require.Equal(t, exp, act)
}

func Test_3_TokenGenerator(t *testing.T) {
	gen := NewTokenGenerator()

	_ = gen(token.Equ, "==")
	_ = gen(token.Neq, "!=")
	act := gen(token.Lte, "<=")

	exp := token.MakeTokenAt(token.Lte, "<=", pos.At(4, 0, 4))

	require.Equal(t, exp, act)
}
