package rinser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/token"

	"github.com/PaulioRandall/firefly-go/workflow/token/tokentest"
)

func tok(tt token.TokenType, v string) token.Token {
	return tokentest.Tok(tt, v)
}

func assert(t *testing.T, given, exp []token.Token) {
	r := inout.NewListReader(given)
	w := inout.NewListWriter[token.Token]()

	e := Clean(r, w)

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, w.List())
}

func Test_1_Clean(t *testing.T) {
	given := []token.Token{}

	var exp []token.Token

	assert(t, given, exp)
}

func Test_2_Clean(t *testing.T) {
	given := []token.Token{
		tok(token.Space, " "),
	}

	var exp []token.Token

	assert(t, given, exp)
}

func Test_3_Clean(t *testing.T) {
	given := []token.Token{
		tok(token.Comment, "//"),
	}

	var exp []token.Token

	assert(t, given, exp)
}

func Test_4_Clean(t *testing.T) {
	given := []token.Token{
		tok(token.Var, "abc"),
	}

	exp := []token.Token{
		tok(token.Var, "abc"),
	}

	assert(t, given, exp)
}

func Test_5_Clean(t *testing.T) {
	given := []token.Token{
		tok(token.Var, "abc"),
		tok(token.Space, " "),
		tok(token.Assign, "="),
		tok(token.Space, " "),
		tok(token.Number, "0"),
		tok(token.Space, " "),
		tok(token.Comment, "//"),
		tok(token.Newline, "\n"),
	}

	exp := []token.Token{
		tok(token.Var, "abc"),
		tok(token.Assign, "="),
		tok(token.Number, "0"),
		tok(token.Newline, "\n"),
	}

	assert(t, given, exp)
}

func Test_6_Clean(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
		tok(token.Newline, "\n"),
		tok(token.Newline, "\n"),
		tok(token.Newline, "\n"),
		tok(token.Number, "0"),
	}

	exp := []token.Token{
		tok(token.String, `""`),
		tok(token.Newline, "\n"),
		tok(token.Number, "0"),
	}

	assert(t, given, exp)
}

func Test_7_Clean(t *testing.T) {
	given := []token.Token{
		tok(token.Space, "\n"),
	}

	var exp []token.Token

	assert(t, given, exp)
}

func Test_8_Clean(t *testing.T) {
	given := []token.Token{
		tok(token.Space, "\n"),
		tok(token.Space, "\n"),
	}

	var exp []token.Token

	assert(t, given, exp)
}
