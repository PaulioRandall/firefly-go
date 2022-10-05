package terminator

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/pos"
	"github.com/PaulioRandall/firefly-go/workflow/token"

	"github.com/PaulioRandall/firefly-go/workflow/token/tokentest"
)

func tok(tt token.TokenType, v string) token.Token {
	return token.MakeToken(tt, v, pos.Range{})
}

func assert(t *testing.T, given, exp []token.Token) {
	r := inout.NewListReader(given)
	w := inout.NewListWriter[token.Token]()

	e := Terminate(r, w)

	require.Nil(t, e, "%+v", e)
	tokentest.RequireEqual(t, exp, w.List())
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
		tok(token.Newline, "\n"),
	}

	exp := []token.Token{
		tok(token.Terminator, "\n"),
	}

	assert(t, given, exp)
}

func Test_3(t *testing.T) {
	given := []token.Token{
		tok(token.Number, "0"),
		tok(token.Newline, "\n"),
		tok(token.Var, "a"),
	}

	exp := []token.Token{
		tok(token.Number, "0"),
		tok(token.Terminator, "\n"),
		tok(token.Var, "a"),
	}

	assert(t, given, exp)
}
