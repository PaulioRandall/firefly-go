package scanner

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/readers"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func Test_ScanAll_1(t *testing.T) {
	r := readers.NewRuneStringReader("")

	act, e := ScanAll(r)
	var exp []token.Token

	require.Nil(t, e)
	require.Equal(t, exp, act)
}

func Test_ScanAll_2(t *testing.T) {
	r := readers.NewRuneStringReader("~")
	_, e := ScanAll(r)
	require.NotNil(t, e)
}

func Test_ScanAll_3(t *testing.T) {
	r := readers.NewRuneStringReader("if")

	act, e := ScanAll(r)
	exp := []token.Token{
		token.Token{
			Type:    token.If,
			Value:   "if",
			FilePos: token.MakeInlineRange(0, 0, 0, 2),
		},
	}

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}
