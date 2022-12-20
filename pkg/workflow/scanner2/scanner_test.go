package scanner2

import (
	"strings"
	"testing"

	"github.com/PaulioRandall/go-trackerr"
	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"
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

func Test_40(t *testing.T) {
	_, e := doTestScan(`"`)
	require.True(t, trackerr.AllOrdered(e, ErrScanning, ErrUnterminatedString))
}

func Test_41(t *testing.T) {
	_, e := doTestScan(`"""`)
	require.True(t, trackerr.AllOrdered(e, ErrScanning, ErrUnterminatedString))
}

func Test_42(t *testing.T) {
	_, e := doTestScan(`"\`)
	require.True(t, trackerr.AllOrdered(e, ErrScanning, ErrUnterminatedString))
}

func Test_43(t *testing.T) {
	_, e := doTestScan(`"\"`)
	require.True(t, trackerr.AllOrdered(e, ErrScanning, ErrUnterminatedString))
}

func Test_44(t *testing.T) {
	_, e := doTestScan(`"\\\"`)
	require.True(t, trackerr.AllOrdered(e, ErrScanning, ErrUnterminatedString))
}

func Test_45(t *testing.T) {
	doSingleTokenTest(t, "0", token.Number)
}

func Test_46(t *testing.T) {
	doSingleTokenTest(t, "0.00000", token.Number)
}

func Test_47(t *testing.T) {
	doSingleTokenTest(t, "0.1", token.Number)
}

func Test_48(t *testing.T) {
	doSingleTokenTest(t, "1", token.Number)
}

func Test_49(t *testing.T) {
	doSingleTokenTest(t, "1.1234567890", token.Number)
}

func Test_50(t *testing.T) {
	doSingleTokenTest(t, "123456789.987654321", token.Number)
}

func Test_51(t *testing.T) {
	doSingleTokenTest(t, "9", token.Number)
}

func Test_52(t *testing.T) {
	_, e := doTestScan("0.")
	require.True(t, trackerr.AllOrdered(e, ErrScanning, ErrMissingFractional))
}

func Test_53(t *testing.T) {
	_, e := doTestScan("0.a")
	require.True(t, trackerr.AllOrdered(e, ErrScanning, ErrMissingFractional))
}

func Test_54(t *testing.T) {
	doSingleTokenTest(t, "&&", token.And)
}

func Test_55(t *testing.T) {
	doSingleTokenTest(t, "||", token.Or)
}

func Test_56(t *testing.T) {
	_, e := doTestScan("~")
	require.True(t, trackerr.AllOrdered(e, ErrScanning, ErrUnknownSymbol))
}

func Test_60(t *testing.T) {
	doSingleTokenTest(t, "abc", token.Ident)
}

func Test_61(t *testing.T) {
	doSingleTokenTest(t, "abc_xyz", token.Ident)
}

func Test_62(t *testing.T) {
	doSingleTokenTest(t, "forest", token.Ident)
}

func Test_63(t *testing.T) {
	doSingleTokenTest(t, "For", token.Ident)
}

func Test_64(t *testing.T) {
	doSingleTokenTest(t, "FOR", token.Ident)
}

func Test_65(t *testing.T) {
	doSingleTokenTest(t, "e", token.Ident)
}

func Test_70(t *testing.T) {
	doSingleTokenTest(t, "if", token.If)
}

func Test_71(t *testing.T) {
	doSingleTokenTest(t, "for", token.For)
}

func Test_72(t *testing.T) {
	doSingleTokenTest(t, "in", token.In)
}

func Test_74(t *testing.T) {
	doSingleTokenTest(t, "watch", token.Watch)
}

func Test_75(t *testing.T) {
	doSingleTokenTest(t, "when", token.When)
}

func Test_76(t *testing.T) {
	doSingleTokenTest(t, "is", token.Is)
}

func Test_110(t *testing.T) {
	doSingleTokenTest(t, "def", token.Def)
}

func Test_77(t *testing.T) {
	doSingleTokenTest(t, "F", token.Func)
}

func Test_78(t *testing.T) {
	doSingleTokenTest(t, "P", token.Proc)
}

func Test_79(t *testing.T) {
	doSingleTokenTest(t, "true", token.Bool)
}

func Test_80(t *testing.T) {
	doSingleTokenTest(t, "false", token.Bool)
}

func Test_200(t *testing.T) {
	given := "x = 1"

	exp := []token.Token{
		makeToken(token.Ident, "x"),
		makeToken(token.Space, " "),
		makeToken(token.Assign, "="),
		makeToken(token.Space, " "),
		makeToken(token.Number, "1"),
	}

	act, e := doTestScan(given)

	require.Nil(t, e)
	require.Equal(t, exp, act)
}

func Test_201(t *testing.T) {

	lines := []string{
		`x = true`,
		`y, z = 123.456, "string"`,
		``,
		`// A procedure`,
		`def f P(a, b) c, d`,
		`	when a`,
		`		is 1: @println("one")`,
		`		a == b: @println("a == b")`,
		`		true: @println("meh")`,
		`	end`,
		`end`,
		``,
	}

	exp := []token.Token{}
	addToken := func(tt token.TokenType, v string) {
		exp = append(exp, makeToken(tt, v))
	}

	// `x = true`
	addToken(token.Ident, "x")
	addToken(token.Space, " ")
	addToken(token.Assign, "=")
	addToken(token.Space, " ")
	addToken(token.Bool, "true")
	addToken(token.Newline, "\n")

	// `y, z = 123.456, "string"`
	addToken(token.Ident, "y")
	addToken(token.Comma, ",")
	addToken(token.Space, " ")
	addToken(token.Ident, "z")
	addToken(token.Space, " ")
	addToken(token.Assign, "=")
	addToken(token.Space, " ")
	addToken(token.Number, "123.456")
	addToken(token.Comma, ",")
	addToken(token.Space, " ")
	addToken(token.String, `"string"`)
	addToken(token.Newline, "\n")

	addToken(token.Newline, "\n")

	// `// A procedure`
	addToken(token.Comment, "// A procedure")
	addToken(token.Newline, "\n")

	// `def f P(a, b) c, d`
	addToken(token.Def, "def")
	addToken(token.Space, " ")
	addToken(token.Ident, "f")
	addToken(token.Space, " ")
	addToken(token.Proc, "P")
	addToken(token.ParenOpen, "(")
	addToken(token.Ident, "a")
	addToken(token.Comma, ",")
	addToken(token.Space, " ")
	addToken(token.Ident, "b")
	addToken(token.ParenClose, ")")
	addToken(token.Space, " ")
	addToken(token.Ident, "c")
	addToken(token.Comma, ",")
	addToken(token.Space, " ")
	addToken(token.Ident, "d")
	addToken(token.Newline, "\n")

	// `	when a`
	addToken(token.Space, "\t")
	addToken(token.When, "when")
	addToken(token.Space, " ")
	addToken(token.Ident, "a")
	addToken(token.Newline, "\n")

	// `		is 1: @println("one")`
	addToken(token.Space, "\t\t")
	addToken(token.Is, "is")
	addToken(token.Space, " ")
	addToken(token.Number, "1")
	addToken(token.Colon, ":")
	addToken(token.Space, " ")
	addToken(token.Spell, "@")
	addToken(token.Ident, "println")
	addToken(token.ParenOpen, "(")
	addToken(token.String, `"one"`)
	addToken(token.ParenClose, ")")
	addToken(token.Newline, "\n")

	// `		a == b: @println("a == b")`
	addToken(token.Space, "\t\t")
	addToken(token.Ident, "a")
	addToken(token.Space, " ")
	addToken(token.Equ, "==")
	addToken(token.Space, " ")
	addToken(token.Ident, "b")
	addToken(token.Colon, ":")
	addToken(token.Space, " ")
	addToken(token.Spell, "@")
	addToken(token.Ident, "println")
	addToken(token.ParenOpen, "(")
	addToken(token.String, `"a == b"`)
	addToken(token.ParenClose, ")")
	addToken(token.Newline, "\n")

	// `		true: @println("meh")`
	addToken(token.Space, "\t\t")
	addToken(token.Bool, "true")
	addToken(token.Colon, ":")
	addToken(token.Space, " ")
	addToken(token.Spell, "@")
	addToken(token.Ident, "println")
	addToken(token.ParenOpen, "(")
	addToken(token.String, `"meh"`)
	addToken(token.ParenClose, ")")
	addToken(token.Newline, "\n")

	// `	end`
	addToken(token.Space, "\t")
	addToken(token.End, "end")
	addToken(token.Newline, "\n")

	// `end`
	addToken(token.End, "end")
	addToken(token.Newline, "\n")

	given := strings.Join(lines, "\n")
	act, e := doTestScan(given)

	require.Nil(t, e)
	require.Equal(t, exp, act)
}
