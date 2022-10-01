package formaliser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/token"
	"github.com/PaulioRandall/firefly-go/workflow/tokenreader"
)

func tok(tt token.TokenType, v string) token.Token {
	return token.MakeToken(tt, v, token.Range{})
}

func assert(t *testing.T, given, exp []token.Token) {
	tr := tokenreader.FromList(given...)
	act := Formalise(tr)
	require.Equal(t, exp, act)
}

func Test_1(t *testing.T) {
	given := []token.Token{
		tok(token.Number, "0"),
	}

	exp := []token.Token{
		tok(token.Number, "0"),
	}

	assert(t, given, exp)
}

func Test_2(t *testing.T) {
	given := []token.Token{
		tok(token.Number, "0"),
		tok(token.Newline, "\n"),
	}

	exp := []token.Token{
		tok(token.Number, "0"),
		tok(token.Terminator, "\n"),
	}

	assert(t, given, exp)
}

func Test_3(t *testing.T) {
	given := []token.Token{
		tok(token.Number, "1"),
		tok(token.Add, "+"),
		tok(token.Newline, "\n"),
		tok(token.Number, "2"),
	}

	exp := []token.Token{
		tok(token.Number, "1"),
		tok(token.Add, "+"),
		tok(token.Number, "2"),
	}

	assert(t, given, exp)
}
