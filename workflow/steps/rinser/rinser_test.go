package rinser

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/readers/tokenreader"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

func tok(tt token.TokenType, v string) token.Token {
	return token.MakeToken(tt, v, token.MakeInlineRange(0, 0, 0, len(v)))
}

func assertRinseAll(t *testing.T, given, exp []token.Token) {
	tr := tokenreader.FromList(given...)

	act, e := RinseAll(tr)
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func assertRinse(t *testing.T, given, expTk token.Token) {
	tr := tokenreader.FromList(given)

	act, e := RinseAll(tr)
	exp := []token.Token{expTk}

	require.Nil(t, e, "Expected %q but got error: %+v", exp, err.DebugString(e))
	require.Equal(t, exp, act,
		"Expected %q but got %q", exp, act,
	)
}

func assertRinseError(t *testing.T, given []token.Token, exp error) {
	tr := tokenreader.FromList(given...)
	_, e := RinseAll(tr)
	require.True(t, errors.Is(e, exp), "Expected %+v", exp.Error())
}

func Test_1_RinseAll(t *testing.T) {
	tr := tokenreader.FromList()

	act, e := RinseAll(tr)

	require.True(t, errors.Is(e, err.EOF))
	require.Empty(t, act)
}
