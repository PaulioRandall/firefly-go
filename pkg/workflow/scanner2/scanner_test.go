package scanner2

import (
	"testing"

	//"github.com/PaulioRandall/go-trackerr"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
	"github.com/stretchr/testify/require"
)

type mockReader struct {
	runes []rune
}

func newMockReader(s string) *mockReader {
	return &mockReader{
		runes: []rune(s),
	}
}

func (r *mockReader) More() bool {
	return len(r.runes) > 0
}

func (r *mockReader) Read() rune {
	if len(r.runes) == 0 {
		return rune(0)
	}

	ru := r.runes[0]
	r.runes = r.runes[1:]
	return ru
}

func (r *mockReader) Putback(ru rune) {
	r.runes = append([]rune{ru}, r.runes...)
}

type mockWriter struct {
	tks []token.Token
}

func (w *mockWriter) Write(tk token.Token) error {
	w.tks = append(w.tks, tk)
	return nil
}

/*
func assertError(t *testing.T, given string, exp ...error) {
	r := inout.NewListReader([]rune(given))
	w := inout.NewListWriter[token.Token]()

	e := Scan(r, w)
	require.True(t, trackerr.AllOrdered(e, exp...))
}
*/

func doTestScan(s string) ([]token.Token, error) {
	r := newMockReader(s)
	w := &mockWriter{}
	e := Scan(r, w)
	return w.tks, e
}

func doSingleTokenTest(t *testing.T, given string, tt token.TokenType, v string) {
	act, e := doTestScan(given)
	require.Nil(t, e)

	exp := expSingleToken(tt, v)
	require.Equal(t, exp, act)
}

func expSingleToken(tt token.TokenType, v string) []token.Token {
	return []token.Token{
		token.Token{
			TokenType: tt,
			Value:     v,
		},
	}
}

func Test_1(t *testing.T) {
	t.Log("Empty scanner input should return empty token list, not an error")

	act, e := doTestScan("")

	require.Nil(t, e)
	require.Empty(t, act)
}

func Test_2(t *testing.T) {
	doSingleTokenTest(t, "\n", token.Newline, "\n")
}
