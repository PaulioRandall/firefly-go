package terminator

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/pos"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

func tok(tt token.TokenType, v string) token.Token {
	return token.MakeToken(tt, v, pos.Range{})
}

func assert(t *testing.T, given, exp []token.Token) {
	r := inout.NewListReader(given)
	w := inout.NewListWriter[token.Token]()

	e := Terminate(r, w)

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, w.List())
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

func assertRemovesNewlineAfter(t *testing.T, given token.Token) {
	in := []token.Token{
		given,
		tok(token.Newline, "\n"),
	}

	exp := []token.Token{
		given,
	}

	assert(t, in, exp)
}

func Test_4(t *testing.T) {
	assertRemovesNewlineAfter(t, tok(token.ParenOpen, "("))
}

func Test_5(t *testing.T) {
	assertRemovesNewlineAfter(t, tok(token.BraceOpen, "{"))
}

func Test_6(t *testing.T) {
	assertRemovesNewlineAfter(t, tok(token.BracketOpen, "["))
}

func assertRemovesNewlineBefore(t *testing.T, given token.Token) {
	in := []token.Token{
		tok(token.Number, "0"),
		tok(token.Newline, "\n"),
		given,
	}

	exp := []token.Token{
		tok(token.Number, "0"),
		given,
	}

	assert(t, in, exp)
}

func Test_7(t *testing.T) {
	assertRemovesNewlineBefore(t, tok(token.ParenClose, ")"))
}

func Test_8(t *testing.T) {
	assertRemovesNewlineBefore(t, tok(token.BraceClose, "}"))
}

func Test_9(t *testing.T) {
	assertRemovesNewlineBefore(t, tok(token.BracketClose, "]"))
}
