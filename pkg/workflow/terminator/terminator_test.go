package terminator

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func tok(tt token.TokenType, v string) token.Token {
	return token.MakeToken(tt, v, pos.Pos{}, pos.Pos{})
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
		tok(token.Identifier, "a"),
	}

	exp := []token.Token{
		tok(token.Number, "0"),
		tok(token.Terminator, "\n"),
		tok(token.Identifier, "a"),
	}

	assert(t, given, exp)
}
