package scanner

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/firefly/token"
)

func happyTest(t *testing.T, in []rune, exp []token.Lexeme) {
	r := token.NewRuneReader(in)
	act, e := ScanAll(r)
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func lex(tk token.Token, v string) token.Lexeme {
	return token.Lexeme{
		Token: tk,
		Value: v,
	}
}

func TestScanAll_0(t *testing.T) {

	// GIVEN nothing
	in := []rune("")

	var exp []token.Lexeme

	// WHEN scanning all tokens
	// THEN the code should be parsed without error
	// AND the output should be a nil slice
	happyTest(t, in, exp)
}

func TestScanAll_1(t *testing.T) {

	// GIVEN a single digit number
	in := []rune("9")

	exp := []token.Lexeme{
		lex(token.TK_NUMBER, "9"),
	}

	// WHEN scanning all tokens
	// THEN the code should be parsed without error
	// AND the output should match 'exp'
	happyTest(t, in, exp)
}

func TestScanAll_2(t *testing.T) {

	// GIVEN a multi-digit number
	in := []rune("999")

	exp := []token.Lexeme{
		lex(token.TK_NUMBER, "999"),
	}

	// WHEN scanning all tokens
	// THEN the code should be parsed without error
	// AND the output should match the 'exp'
	happyTest(t, in, exp)
}

func TestScanAll_3(t *testing.T) {

	// GIVEN an operator
	// WHEN scanning all tokens
	// THEN the code should be parsed without error
	// AND the output should match 'exp'
	doTest := func(op string, tk token.Token) {
		in := []rune(op)
		exp := []token.Lexeme{lex(tk, op)}
		happyTest(t, in, exp)
	}

	doTest("+", token.TK_ADD)
	doTest("-", token.TK_SUB)
	doTest("*", token.TK_MUL)
	doTest("/", token.TK_DIV)
}

func TestScanAll_4(t *testing.T) {

	// GIVEN parentheses
	// WHEN scanning all tokens
	// THEN the parentheses should be parsed without error
	// AND the output should match 'exp'
	doTest := func(op string, tk token.Token) {
		in := []rune(op)
		exp := []token.Lexeme{lex(tk, op)}
		happyTest(t, in, exp)
	}

	doTest("(", token.TK_PAREN_OPEN)
	doTest(")", token.TK_PAREN_CLOSE)
}

func TestScanAll_5(t *testing.T) {

	// GIVEN a long expression
	r := token.NewRuneReader(
		[]rune("1 + 2 - 3 * 4 / 5"),
	)

	exp := []token.Lexeme{
		lex(token.TK_NUMBER, "1"),
		lex(token.TK_SPACE, " "),
		lex(token.TK_ADD, "+"),
		lex(token.TK_SPACE, " "),
		lex(token.TK_NUMBER, "2"),
		lex(token.TK_SPACE, " "),
		lex(token.TK_SUB, "-"),
		lex(token.TK_SPACE, " "),
		lex(token.TK_NUMBER, "3"),
		lex(token.TK_SPACE, " "),
		lex(token.TK_MUL, "*"),
		lex(token.TK_SPACE, " "),
		lex(token.TK_NUMBER, "4"),
		lex(token.TK_SPACE, " "),
		lex(token.TK_DIV, "/"),
		lex(token.TK_SPACE, " "),
		lex(token.TK_NUMBER, "5"),
	}

	// WHEN scanning all tokens
	act, e := ScanAll(r)

	// THEN the code should be correctly parsed without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func TestScanAll_6(t *testing.T) {

	// GIVEN multiple statements
	r := token.NewRuneReader(
		[]rune("1\n2\n3\n"),
	)

	exp := []token.Lexeme{
		lex(token.TK_NUMBER, "1"),
		lex(token.TK_NEWLINE, "\n"),
		lex(token.TK_NUMBER, "2"),
		lex(token.TK_NEWLINE, "\n"),
		lex(token.TK_NUMBER, "3"),
		lex(token.TK_NEWLINE, "\n"),
	}

	// WHEN scanning all tokens
	act, e := ScanAll(r)

	// THEN the code should be correctly parsed without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func TestScanAll_7(t *testing.T) {

	// GIVEN an invalid token
	r := token.NewRuneReader(
		[]rune("#"),
	)

	// WHEN scanning the token
	_, e := ScanAll(r)

	// THEN an error should be returned
	require.NotNil(t, e, "Expected error when given invalid token")
}
