package aligner

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func tok(tt token.TokenType, v string) token.Token {
	return tokentest.Tok(tt, v)
}

func assert(t *testing.T, given, exp []token.Token) {
	r := inout.NewListReader(given)
	w := inout.NewListWriter[token.Token]()

	e := Align(r, w)

	require.Nil(t, e, "%+v", e)
	tokentest.RequireEqual(t, exp, w.List())
}

func Test_1_Align(t *testing.T) {
	var given []token.Token
	var exp []token.Token
	assert(t, given, exp)
}

func Test_2_Align(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
		tok(token.Number, "0"),
	}

	exp := []token.Token{
		tok(token.String, `""`),
		tok(token.Number, "0"),
	}

	assert(t, given, exp)
}

func Test_3_Align(t *testing.T) {
	given := []token.Token{
		tok(token.BracketOpen, "["),
	}

	exp := []token.Token{
		tok(token.BracketOpen, "["),
	}

	assert(t, given, exp)
}

func Test_4_Align(t *testing.T) {
	given := []token.Token{
		tok(token.BracketOpen, "["),
		tok(token.BracketClose, "]"),
	}

	exp := []token.Token{
		tok(token.BracketOpen, "["),
		tok(token.BracketClose, "]"),
	}

	assert(t, given, exp)
}

func Test_8_Align(t *testing.T) {
	// [a,
	// b,
	// c,]
	given := []token.Token{
		tok(token.BracketOpen, "["),
		tok(token.Var, "a"),
		tok(token.Comma, ",1"),
		tok(token.Terminator, "\n1"),
		tok(token.Var, "b"),
		tok(token.Comma, ",2"),
		tok(token.Terminator, "\n2"),
		tok(token.Var, "c"),
		tok(token.Comma, ",3"),
		tok(token.BracketClose, "]"),
	}

	// [a,b,c]
	exp := []token.Token{
		tok(token.BracketOpen, "["),
		tok(token.Var, "a"),
		tok(token.Comma, ",1"),
		tok(token.Var, "b"),
		tok(token.Comma, ",2"),
		tok(token.Var, "c"),
		tok(token.BracketClose, "]"),
	}

	assert(t, given, exp)
}

func Test_9_Align(t *testing.T) {
	// [a,
	// b,c,
	// d,]
	given := []token.Token{
		tok(token.BracketOpen, "["),
		tok(token.Var, "a"),
		tok(token.Comma, ",1"),
		tok(token.Terminator, "\n2"),
		tok(token.Var, "b"),
		tok(token.Comma, ",2"),
		tok(token.Var, "c"),
		tok(token.Comma, ",3"),
		tok(token.Terminator, "\n3"),
		tok(token.Var, "d"),
		tok(token.Comma, ",4"),
		tok(token.BracketClose, "]"),
	}

	// [a,b,c,d]
	exp := []token.Token{
		tok(token.BracketOpen, "["),
		tok(token.Var, "a"),
		tok(token.Comma, ",1"),
		tok(token.Var, "b"),
		tok(token.Comma, ",2"),
		tok(token.Var, "c"),
		tok(token.Comma, ",3"),
		tok(token.Var, "d"),
		tok(token.BracketClose, "]"),
	}

	assert(t, given, exp)
}

func Test_10_Align(t *testing.T) {
	given := []token.Token{
		tok(token.Comma, ","),
		tok(token.BracketClose, "]"),
	}

	exp := []token.Token{
		tok(token.BracketClose, "]"),
	}

	assert(t, given, exp)
}

func Test_11_Align(t *testing.T) {
	given := []token.Token{
		tok(token.Comma, ","),
		tok(token.BraceClose, "}"),
	}

	exp := []token.Token{
		tok(token.BraceClose, "}"),
	}

	assert(t, given, exp)
}

func Test_12_Align(t *testing.T) {
	given := []token.Token{
		tok(token.Comma, ","),
		tok(token.ParenClose, ")"),
	}

	exp := []token.Token{
		tok(token.ParenClose, ")"),
	}

	assert(t, given, exp)
}
