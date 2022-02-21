package scanner

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func given(in string) RuneReader {
	return token.NewRuneReader([]rune(in))
}

func lex(tk token.Token, v string) token.Lexeme {
	return token.Lexeme{
		Token: tk,
		Value: v,
	}
}

func TestScanAll_0(t *testing.T) {
	r := given("")

	act, e := ScanAll(r)

	require.Nil(t, e, "%+v", e)

	var exp []token.Lexeme
	require.Equal(t, exp, act)
}

func TestScanAll_bool_1(t *testing.T) {
	r := given("true")

	act, e := ScanAll(r)

	require.Nil(t, e, "%+v", e)

	exp := []token.Lexeme{
		lex(token.TK_BOOL, "true"),
	}
	require.Equal(t, exp, act)
}

func TestScanAll_bool_2(t *testing.T) {
	r := given("false")

	act, e := ScanAll(r)

	require.Nil(t, e, "%+v", e)

	exp := []token.Lexeme{
		lex(token.TK_BOOL, "false"),
	}
	require.Equal(t, exp, act)
}

func TestScanAll_number_1(t *testing.T) {
	r := given("123")

	act, e := ScanAll(r)

	require.Nil(t, e, "%+v", e)

	exp := []token.Lexeme{
		lex(token.TK_NUM, "123"),
	}
	require.Equal(t, exp, act)
}

func TestScanAll_number_2(t *testing.T) {
	r := given("123.456")

	act, e := ScanAll(r)

	require.Nil(t, e, "%+v", e)

	exp := []token.Lexeme{
		lex(token.TK_NUM, "123.456"),
	}
	require.Equal(t, exp, act)
}

func TestScanAll_number_3(t *testing.T) {
	r := given("1_234_567")

	act, e := ScanAll(r)

	require.Nil(t, e, "%+v", e)

	exp := []token.Lexeme{
		lex(token.TK_NUM, "1_234_567"),
	}
	require.Equal(t, exp, act)
}

func TestScanAll_number_4(t *testing.T) {
	r := given("123.")

	_, e := ScanAll(r)

	require.NotNil(t, e, "Expected error")
}

func TestScanAll_number_5(t *testing.T) {
	r := given("123.a")

	_, e := ScanAll(r)

	require.NotNil(t, e, "Expected error")
}

func TestScanAll_string_1(t *testing.T) {
	r := given(`""`)

	act, e := ScanAll(r)

	require.Nil(t, e, "%+v", e)

	exp := []token.Lexeme{
		lex(token.TK_STR, `""`),
	}
	require.Equal(t, exp, act)
}

func TestScanAll_string_2(t *testing.T) {
	r := given(`"abc"`)

	act, e := ScanAll(r)

	require.Nil(t, e, "%+v", e)

	exp := []token.Lexeme{
		lex(token.TK_STR, `"abc"`),
	}
	require.Equal(t, exp, act)
}

func TestScanAll_string_3(t *testing.T) {
	r := given(`"ab\"cd"`)

	act, e := ScanAll(r)

	require.Nil(t, e, "%+v", e)

	exp := []token.Lexeme{
		lex(token.TK_STR, `"ab\"cd"`),
	}
	require.Equal(t, exp, act)
}

func TestScanAll_string_4(t *testing.T) {
	r := given(`"abc\\\\\\xyz"`)

	act, e := ScanAll(r)

	require.Nil(t, e, "%+v", e)

	exp := []token.Lexeme{
		lex(token.TK_STR, `"abc\\\\\\xyz"`),
	}
	require.Equal(t, exp, act)
}

func TestScanAll_string_5(t *testing.T) {
	r := given(`"abc xyz"`)

	act, e := ScanAll(r)

	require.Nil(t, e, "%+v", e)

	exp := []token.Lexeme{
		lex(token.TK_STR, `"abc xyz"`),
	}
	require.Equal(t, exp, act)
}

func TestScanAll_string_6(t *testing.T) {
	r := given(`"abc`)

	_, e := ScanAll(r)

	require.NotNil(t, e, "Expected error")
}

func TestScanAll_string_7(t *testing.T) {
	r := given(`"\"`)

	_, e := ScanAll(r)

	require.NotNil(t, e, "Expected error")
}

func TestScanAll_string_8(t *testing.T) {
	r := given(`"\\""`)

	_, e := ScanAll(r)

	require.NotNil(t, e, "Expected error")
}
