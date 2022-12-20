package cleaner2

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

type mockReader struct {
	tks []token.Token
}

func newMockReader(tks ...token.Token) *mockReader {
	return &mockReader{
		tks: tks,
	}
}

func (r *mockReader) More() bool {
	return len(r.tks) > 0
}

func (r *mockReader) Read() token.Token {
	if len(r.tks) == 0 {
		panic("Empty reader")
	}

	tk := r.tks[0]
	r.tks = r.tks[1:]
	return tk
}

type mockWriter struct {
	tks []token.Token
}

func (w *mockWriter) Write(tk token.Token) error {
	w.tks = append(w.tks, tk)
	return nil
}

func doTestClean(given ...token.Token) ([]token.Token, error) {
	r := newMockReader(given...)
	w := &mockWriter{}
	e := Clean(r, w)
	return w.tks, e
}

func doTestEqual(t *testing.T, given, exp []token.Token) {
	act, e := doTestClean(given...)
	require.Nil(t, e)
	require.Equal(t, exp, act)
}

func tok(tt token.TokenType, v string) token.Token {
	return token.Token{
		TokenType: tt,
		Value:     v,
	}
}

func Test_1(t *testing.T) {
	act, e := doTestClean()
	require.Nil(t, e)
	require.Empty(t, act)
}

func Test_2(t *testing.T) {
	given := []token.Token{
		tok(token.Space, " "),
	}

	var exp []token.Token

	doTestEqual(t, given, exp)
}

func Test_3(t *testing.T) {
	given := []token.Token{
		tok(token.Comment, "//"),
	}

	var exp []token.Token

	doTestEqual(t, given, exp)
}

func Test_4(t *testing.T) {
	given := []token.Token{
		tok(token.Ident, "abc"),
	}

	exp := []token.Token{
		tok(token.Ident, "abc"),
	}

	doTestEqual(t, given, exp)
}

func Test_5(t *testing.T) {
	given := []token.Token{
		tok(token.Ident, "abc"),
		tok(token.Space, " "),
		tok(token.Assign, "="),
		tok(token.Space, " "),
		tok(token.Number, "0"),
		tok(token.Space, " "),
		tok(token.Comment, "//"),
		tok(token.Newline, "\n"),
	}

	exp := []token.Token{
		tok(token.Ident, "abc"),
		tok(token.Assign, "="),
		tok(token.Number, "0"),
		tok(token.Newline, "\n"),
	}

	doTestEqual(t, given, exp)
}

func Test_6(t *testing.T) {
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

	doTestEqual(t, given, exp)
}

func Test_7(t *testing.T) {
	given := []token.Token{
		tok(token.Space, "\n"),
	}

	var exp []token.Token

	doTestEqual(t, given, exp)
}

func Test_8(t *testing.T) {
	given := []token.Token{
		tok(token.Space, "\n"),
		tok(token.Space, "\n"),
	}

	var exp []token.Token

	doTestEqual(t, given, exp)
}

func Test_9(t *testing.T) {
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

	doTestEqual(t, given, exp)
}

func assertRemovesNewlineAfter(t *testing.T, given token.Token) {
	in := []token.Token{
		given,
		tok(token.Newline, "\n"),
	}

	exp := []token.Token{
		given,
	}

	doTestEqual(t, in, exp)
}

func Test_10(t *testing.T) {
	assertRemovesNewlineAfter(t, tok(token.ParenOpen, "("))
}

func Test_11(t *testing.T) {
	assertRemovesNewlineAfter(t, tok(token.BraceOpen, "{"))
}

func Test_12(t *testing.T) {
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

	doTestEqual(t, in, exp)
}

func Test_13(t *testing.T) {
	assertRemovesNewlineBefore(t, tok(token.ParenClose, ")"))
}

func Test_14(t *testing.T) {
	assertRemovesNewlineBefore(t, tok(token.BraceClose, "}"))
}

func Test_15(t *testing.T) {
	assertRemovesNewlineBefore(t, tok(token.BracketClose, "]"))
}
