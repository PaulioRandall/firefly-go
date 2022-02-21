package scanner

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func given(in string) RuneReader {
	return token.NewRuneReader([]rune(in))
}

func expect(out ...token.Lexeme) []token.Lexeme {
	if len(out) == 0 {
		return nil
	}
	return out
}

func then(t *testing.T, exp, act []token.Lexeme) {
	require.Equal(t, exp, act)
}

func thenNoError(t *testing.T, e error) {
	require.Nil(t, e, "%+v", e)
}

func thenError(t *testing.T, e error, failMsgAndArgs ...interface{}) {
	require.NotNil(t, e, failMsgAndArgs...)
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
	exp := expect()

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_bool_1(t *testing.T) {
	r := given("true")

	act, e := ScanAll(r)
	exp := expect(
		lex(token.TK_BOOL, "true"),
	)

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_bool_2(t *testing.T) {
	r := given("false")

	act, e := ScanAll(r)
	exp := expect(
		lex(token.TK_BOOL, "false"),
	)

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_number_1(t *testing.T) {
	r := given("123")

	act, e := ScanAll(r)
	exp := expect(
		lex(token.TK_NUM, "123"),
	)

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_number_2(t *testing.T) {
	r := given("123.456")

	act, e := ScanAll(r)
	exp := expect(
		lex(token.TK_NUM, "123.456"),
	)

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_number_3(t *testing.T) {
	r := given("1_234_567")

	act, e := ScanAll(r)
	exp := expect(
		lex(token.TK_NUM, "1_234_567"),
	)

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_number_4(t *testing.T) {
	r := given("123.")

	_, e := ScanAll(r)

	thenError(t, e, "Expected missing digit error after decimal point")
}

func TestScanAll_number_5(t *testing.T) {
	r := given("123.a")

	_, e := ScanAll(r)

	thenError(t, e, "Expected missing digit error after decimal point")
}

func TestScanAll_string_1(t *testing.T) {
	r := given(`""`)

	act, e := ScanAll(r)
	exp := expect(
		lex(token.TK_STR, `""`),
	)

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_string_2(t *testing.T) {
	r := given(`"abc"`)

	act, e := ScanAll(r)
	exp := expect(
		lex(token.TK_STR, `"abc"`),
	)

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_string_3(t *testing.T) {
	r := given(`"ab\"cd"`)

	act, e := ScanAll(r)
	exp := expect(
		lex(token.TK_STR, `"ab\"cd"`),
	)

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_string_4(t *testing.T) {
	r := given(`"abc\\\\\\xyz"`)

	act, e := ScanAll(r)
	exp := expect(
		lex(token.TK_STR, `"abc\\\\\\xyz"`),
	)

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_string_5(t *testing.T) {
	r := given(`"abc xyz"`)

	act, e := ScanAll(r)
	exp := expect(
		lex(token.TK_STR, `"abc xyz"`),
	)

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_string_6(t *testing.T) {
	r := given(`"abc`)

	_, e := ScanAll(r)

	thenError(t, e, "Expected unterminated string error")
}

func TestScanAll_string_7(t *testing.T) {
	r := given(`"\"`)

	_, e := ScanAll(r)

	thenError(t, e, "Expected unterminated string error")
}

func TestScanAll_string_8(t *testing.T) {
	r := given(`"\\""`)

	_, e := ScanAll(r)

	thenError(t, e, "Expected unterminated string error")
}

func TestScanAll_string_9(t *testing.T) {
	r := given(`"abc` + "\n")

	_, e := ScanAll(r)

	thenError(t, e, "Expected unterminated string error")
}
