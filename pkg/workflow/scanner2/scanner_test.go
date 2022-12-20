package scanner2

import (
	"testing"

	"github.com/PaulioRandall/go-trackerr"

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

func doSingleTokenTest(t *testing.T, given string, tt token.TokenType) {
	act, e := doTestScan(given)
	require.Nil(t, e, trackerr.ErrorStack(e))

	exp := expSingleToken(tt, given)
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
	doSingleTokenTest(t, "\n", token.Newline)
}

func Test_3(t *testing.T) {
	doSingleTokenTest(t, "+", token.Add)
}

func Test_4(t *testing.T) {
	doSingleTokenTest(t, "-", token.Sub)
}

func Test_5(t *testing.T) {
	doSingleTokenTest(t, "*", token.Mul)
}

func Test_6(t *testing.T) {
	doSingleTokenTest(t, "/", token.Div)
}

func Test_7(t *testing.T) {
	doSingleTokenTest(t, "%", token.Mod)
}

func Test_8(t *testing.T) {
	doSingleTokenTest(t, "<", token.Lt)
}

func Test_9(t *testing.T) {
	doSingleTokenTest(t, ">", token.Gt)
}

func Test_10(t *testing.T) {
	doSingleTokenTest(t, "<=", token.Lte)
}

func Test_11(t *testing.T) {
	doSingleTokenTest(t, ">=", token.Gte)
}

func Test_12(t *testing.T) {
	doSingleTokenTest(t, "==", token.Equ)
}

func Test_13(t *testing.T) {
	doSingleTokenTest(t, "!=", token.Neq)
}

func Test_14(t *testing.T) {
	doSingleTokenTest(t, "=", token.Assign)
}

func Test_15(t *testing.T) {
	doSingleTokenTest(t, ":", token.Colon)
}

func Test_16(t *testing.T) {
	doSingleTokenTest(t, ";", token.Terminator)
}

func Test_17(t *testing.T) {
	doSingleTokenTest(t, ",", token.Comma)
}

func Test_18(t *testing.T) {
	doSingleTokenTest(t, "@", token.Spell)
}

func Test_19(t *testing.T) {
	doSingleTokenTest(t, "(", token.ParenOpen)
}

func Test_20(t *testing.T) {
	doSingleTokenTest(t, ")", token.ParenClose)
}

func Test_21(t *testing.T) {
	doSingleTokenTest(t, "{", token.BraceOpen)
}

func Test_22(t *testing.T) {
	doSingleTokenTest(t, "}", token.BraceClose)
}

func Test_23(t *testing.T) {
	doSingleTokenTest(t, "[", token.BracketOpen)
}

func Test_24(t *testing.T) {
	doSingleTokenTest(t, "]", token.BracketClose)
}

func Test_25(t *testing.T) {
	doSingleTokenTest(t, " ", token.Space)
}

func Test_26(t *testing.T) {
	doSingleTokenTest(t, "\t", token.Space)
}

func Test_27(t *testing.T) {
	doSingleTokenTest(t, "\v", token.Space)
}

func Test_28(t *testing.T) {
	doSingleTokenTest(t, "\r", token.Space)
}

func Test_29(t *testing.T) {
	doSingleTokenTest(t, "\f", token.Space)
}

func Test_30(t *testing.T) {
	doSingleTokenTest(t, "  \t\v \f\r   \v\v\t", token.Space)
}

func Test_31(t *testing.T) {
	doSingleTokenTest(t, "//", token.Comment)
}

func Test_32(t *testing.T) {
	doSingleTokenTest(t, "// abc", token.Comment)
}

func Test_33(t *testing.T) {
	doSingleTokenTest(t, `""`, token.String)
}

func Test_34(t *testing.T) {
	doSingleTokenTest(t, `"a"`, token.String)
}

func Test_35(t *testing.T) {
	doSingleTokenTest(t, `"abc"`, token.String)
}

func Test_36(t *testing.T) {
	doSingleTokenTest(t, `"   "`, token.String)
}

func Test_37(t *testing.T) {
	doSingleTokenTest(t, `"\\"`, token.String)
}

func Test_38(t *testing.T) {
	doSingleTokenTest(t, `"\\\\\\"`, token.String)
}

func Test_39(t *testing.T) {
	doSingleTokenTest(t, `"\"\"\""`, token.String)
}
