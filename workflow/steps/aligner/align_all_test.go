package aligner

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/readers/tokenreader"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

func tok(tt token.TokenType, v string) token.Token {
	return token.MakeToken(tt, v, token.Range{})
}

func assertAlignAll(t *testing.T, given, exp []token.Token) {
	tr := tokenreader.FromList(given...)
	act := AlignAll(tr)
	require.Equal(t, exp, act)
}

func Test_1_AlignAll(t *testing.T) {
	var given []token.Token
	var exp []token.Token
	assertAlignAll(t, given, exp)
}

func Test_2_AlignAll(t *testing.T) {
	given := []token.Token{
		tok(token.String, `""`),
		tok(token.Number, "0"),
	}

	exp := []token.Token{
		tok(token.String, `""`),
		tok(token.Number, "0"),
	}

	assertAlignAll(t, given, exp)
}

func Test_3_AlignAll(t *testing.T) {
	given := []token.Token{
		tok(token.BracketOpen, "["),
	}

	exp := []token.Token{
		tok(token.BracketOpen, "["),
	}

	assertAlignAll(t, given, exp)
}

func Test_4_AlignAll(t *testing.T) {
	given := []token.Token{
		tok(token.BracketOpen, "["),
		tok(token.BracketClose, "]"),
	}

	exp := []token.Token{
		tok(token.BracketOpen, "["),
		tok(token.BracketClose, "]"),
	}

	assertAlignAll(t, given, exp)
}

func Test_5_AlignAll(t *testing.T) {
	given := []token.Token{
		tok(token.BracketOpen, "["),
		tok(token.Newline, "\n"),
		tok(token.BracketClose, "]"),
	}

	exp := []token.Token{
		tok(token.BracketOpen, "["),
		tok(token.BracketClose, "]"),
	}

	assertAlignAll(t, given, exp)
}

func Test_6_AlignAll(t *testing.T) {
	given := []token.Token{
		tok(token.BraceOpen, "{"),
		tok(token.Newline, "\n"),
		tok(token.BraceClose, "}"),
	}

	exp := []token.Token{
		tok(token.BraceOpen, "{"),
		tok(token.BraceClose, "}"),
	}

	assertAlignAll(t, given, exp)
}

func Test_7_AlignAll(t *testing.T) {
	given := []token.Token{
		tok(token.ParenOpen, "("),
		tok(token.Newline, "\n"),
		tok(token.ParenClose, ")"),
	}

	exp := []token.Token{
		tok(token.ParenOpen, "("),
		tok(token.ParenClose, ")"),
	}

	assertAlignAll(t, given, exp)
}

func Test_8_AlignAll(t *testing.T) {
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

	assertAlignAll(t, given, exp)
}

func Test_9_AlignAll(t *testing.T) {
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

	assertAlignAll(t, given, exp)
}

func Test_10_AlignAll(t *testing.T) {
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

	assertAlignAll(t, given, exp)
}

func Test_11_AlignAll(t *testing.T) {
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

	assertAlignAll(t, given, exp)
}

func Test_12_AlignAll(t *testing.T) {
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

	assertAlignAll(t, given, exp)
}
