package aligner

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

	e := Align(r, w)

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, w.List())
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

func Test_5_Align(t *testing.T) {
	given := []token.Token{
		tok(token.BracketOpen, "["),
		tok(token.Newline, "\n"),
		tok(token.BracketClose, "]"),
	}

	exp := []token.Token{
		tok(token.BracketOpen, "["),
		tok(token.BracketClose, "]"),
	}

	assert(t, given, exp)
}

func Test_6_Align(t *testing.T) {
	given := []token.Token{
		tok(token.BraceOpen, "{"),
		tok(token.Newline, "\n"),
		tok(token.BraceClose, "}"),
	}

	exp := []token.Token{
		tok(token.BraceOpen, "{"),
		tok(token.BraceClose, "}"),
	}

	assert(t, given, exp)
}

func Test_7_Align(t *testing.T) {
	given := []token.Token{
		tok(token.ParenOpen, "("),
		tok(token.Newline, "\n"),
		tok(token.ParenClose, ")"),
	}

	exp := []token.Token{
		tok(token.ParenOpen, "("),
		tok(token.ParenClose, ")"),
	}

	assert(t, given, exp)
}

func Test_8_Align(t *testing.T) {
	given := []token.Token{
		tok(token.BracketOpen, "["),
		tok(token.Newline, "\n1"),
		tok(token.Var, "a"),
		tok(token.Newline, "\n2"),
		tok(token.Var, "b"),
		tok(token.Newline, "\n3"),
		tok(token.Var, "c"),
		tok(token.Newline, "\n4"),
		tok(token.BracketClose, "]"),
	}

	exp := []token.Token{
		tok(token.BracketOpen, "["),
		tok(token.Var, "a"),
		tok(token.Comma, "\n2"),
		tok(token.Var, "b"),
		tok(token.Comma, "\n3"),
		tok(token.Var, "c"),
		tok(token.BracketClose, "]"),
	}

	assert(t, given, exp)
}

func Test_9_Align(t *testing.T) {
	given := []token.Token{
		tok(token.BracketOpen, "["),
		tok(token.Newline, "\n1"),
		tok(token.Var, "a"),
		tok(token.Newline, "\n2"),
		tok(token.Var, "b"),
		tok(token.Comma, ","),
		tok(token.Var, "c"),
		tok(token.Newline, "\n3"),
		tok(token.Var, "d"),
		tok(token.Newline, "\n4"),
		tok(token.BracketClose, "]"),
	}

	exp := []token.Token{
		tok(token.BracketOpen, "["),
		tok(token.Var, "a"),
		tok(token.Comma, "\n2"),
		tok(token.Var, "b"),
		tok(token.Comma, ","),
		tok(token.Var, "c"),
		tok(token.Comma, "\n3"),
		tok(token.Var, "d"),
		tok(token.BracketClose, "]"),
	}

	assert(t, given, exp)
}

func Test_10_Align(t *testing.T) {
	given := []token.Token{
		tok(token.ParenOpen, "("),
		tok(token.Newline, "\n1"),
		tok(token.BracketOpen, "["),
		tok(token.Newline, "\n2"),
		tok(token.Var, "a"),
		tok(token.Newline, "\n3"),
		tok(token.Var, "b"),
		tok(token.Newline, "\n4"),
		tok(token.BracketClose, "]"),
		tok(token.Newline, "\n5"),
		tok(token.ParenClose, ")"),
	}

	exp := []token.Token{
		tok(token.ParenOpen, "("),
		tok(token.BracketOpen, "["),
		tok(token.Var, "a"),
		tok(token.Comma, "\n3"),
		tok(token.Var, "b"),
		tok(token.BracketClose, "]"),
		tok(token.ParenClose, ")"),
	}

	assert(t, given, exp)
}

func Test_11_Align(t *testing.T) {
	given := []token.Token{
		tok(token.ParenOpen, "("),
		tok(token.Newline, "\n1"),
		tok(token.BracketOpen, "["),
		tok(token.BracketClose, "]"),
	}

	exp := []token.Token{
		tok(token.ParenOpen, "("),
		tok(token.BracketOpen, "["),
		tok(token.BracketClose, "]"),
	}

	assert(t, given, exp)
}

func Test_12_Align(t *testing.T) {
	given := []token.Token{
		tok(token.ParenOpen, "("),
		tok(token.BracketOpen, "["),
		tok(token.Newline, "\n1"),
		tok(token.ParenClose, ")"),
	}

	exp := []token.Token{
		tok(token.ParenOpen, "("),
		tok(token.BracketOpen, "["),
		tok(token.ParenClose, ")"),
	}

	assert(t, given, exp)
}
