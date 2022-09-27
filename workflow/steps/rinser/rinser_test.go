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

func Test_3_RinseAll(t *testing.T) {
	given := []token.Token{
		tok(token.Comment),
	}

	var exp []token.Token

	assertRinseAll(t, given, exp)
}

func Test_4_RinseAll(t *testing.T) {
	given := []token.Token{
		tok(token.Var),
	}

	exp := []token.Token{
		tok(token.Var),
	}

	assertRinseAll(t, given, exp)
}

func Test_5_RinseAll(t *testing.T) {
	given := []token.Token{
		tok(token.Var),
		tok(token.Space),
		tok(token.Assign),
		tok(token.Space),
		tok(token.Number),
		tok(token.Space),
		tok(token.Comment),
		tok(token.Newline),
	}

	exp := []token.Token{
		tok(token.Var),
		tok(token.Assign),
		tok(token.Number),
		tok(token.Newline),
	}

	assertRinseAll(t, given, exp)
}

func Test_6_RinseAll(t *testing.T) {
	given := []token.Token{
		tok(token.String),
		tok(token.Newline),
		tok(token.Newline),
		tok(token.Newline),
		tok(token.Number),
	}

	exp := []token.Token{
		tok(token.String),
		tok(token.Newline),
		tok(token.Number),
	}

	assertRinseAll(t, given, exp)
}
