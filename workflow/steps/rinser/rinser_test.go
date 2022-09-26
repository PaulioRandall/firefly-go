package rinser

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/readers/tokenreader"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

func tok(tt token.TokenType) token.Token {
	return token.MakeToken(tt, "", token.Range{})
}

func assertRinseAll(t *testing.T, given, exp []token.Token) {
	tr := tokenreader.FromList(given...)

	act, e := RinseAll(tr)

	require.True(t, errors.Is(e, err.EOF))
	require.Equal(t, exp, act)
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

func Test_2_RinseAll(t *testing.T) {
	given := []token.Token{
		tok(token.Space),
	}

	var exp []token.Token

	assertRinseAll(t, given, exp)
}
