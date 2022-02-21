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

func thenNot(t *testing.T, not, act []token.Lexeme) {
	require.NotEqual(t, not, act)
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
		lex(token.TK_NUMBER, "123"),
	)

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_number_2(t *testing.T) {
	r := given("123.456")

	act, e := ScanAll(r)
	exp := expect(
		lex(token.TK_NUMBER, "123.456"),
	)

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_number_3(t *testing.T) {
	r := given("1_234_567")

	act, e := ScanAll(r)
	exp := expect(
		lex(token.TK_NUMBER, "1_234_567"),
	)

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_number_4(t *testing.T) {
	r := given("123.")

	_, e := ScanAll(r)

	thenError(t, e, "Expected missing digit after decimal point error")
}

func TestScanAll_number_5(t *testing.T) {
	r := given("123.a")

	_, e := ScanAll(r)

	thenError(t, e, "Expected missing digit after decimal point error")
}

func TestScanAll_string_1(t *testing.T) {
	r := given(`""`)

	act, e := ScanAll(r)
	exp := expect(
		lex(token.TK_STRING, `""`),
	)

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_string_2(t *testing.T) {
	r := given(`"abc"`)

	act, e := ScanAll(r)
	exp := expect(
		lex(token.TK_STRING, `"abc"`),
	)

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_string_3(t *testing.T) {
	r := given(`"ab\"cd"`)

	act, e := ScanAll(r)
	exp := expect(
		lex(token.TK_STRING, `"ab\"cd"`),
	)

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_string_4(t *testing.T) {
	r := given(`"abc\\\\\\xyz"`)

	act, e := ScanAll(r)
	exp := expect(
		lex(token.TK_STRING, `"abc\\\\\\xyz"`),
	)

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_string_5(t *testing.T) {
	r := given(`"abc xyz"`)

	act, e := ScanAll(r)
	exp := expect(
		lex(token.TK_STRING, `"abc xyz"`),
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

func TestScanAll_ident_1(t *testing.T) {
	r := given("abc")

	act, e := ScanAll(r)
	exp := expect(
		lex(token.TK_IDENT, "abc"),
	)

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_ident_2(t *testing.T) {
	r := given("abc_xyz")

	act, e := ScanAll(r)
	exp := expect(
		lex(token.TK_IDENT, "abc_xyz"),
	)

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_ident_3(t *testing.T) {
	r := given("_")

	act, e := ScanAll(r)
	exp := expect(
		lex(token.TK_IDENT, "_"),
	)

	thenNoError(t, e)
	then(t, exp, act)
}

func TestScanAll_operators_1(t *testing.T) {

	testScan := func(in string) {
		r := given(in)

		act, e := ScanAll(r)
		exp := expect(
			lex(token.TK_OPERATOR, in),
		)

		thenNoError(t, e)
		then(t, exp, act)
	}

	testScan("+")
	testScan("-")
	testScan("*")
	testScan("/")
	testScan("%")

	testScan("<")
	testScan(">")
	testScan("<=")
	testScan(">=")
	testScan("==")
	testScan("!=")

	testScan("<<")
	testScan(">>")
}

func TestScanAll_operator_2(t *testing.T) {
	r := given(">!")

	act, _ := ScanAll(r)
	not := expect(
		lex(token.TK_OPERATOR, ">!"),
	)

	thenNot(t, not, act)
}

func TestScanAll_operator_3(t *testing.T) {
	r := given("!>")

	act, _ := ScanAll(r)
	not := expect(
		lex(token.TK_OPERATOR, "!>"),
	)

	thenNot(t, not, act)
}
