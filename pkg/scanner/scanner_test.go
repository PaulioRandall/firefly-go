package scanner

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/token"
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

func TestScanAll_1(t *testing.T) {

	// GIVEN a single digit number
	in := []rune("9")

	exp := []token.Lexeme{
		lex(token.TokenNumber, "9"),
	}

	// WHEN scanning all tokens
	// THEN the code should be parsed without error
	// AND the output should match the 'exp'
	happyTest(t, in, exp)
}

func TestScanAll_2(t *testing.T) {

	// GIVEN a multi-digit number
	in := []rune("999")

	exp := []token.Lexeme{
		lex(token.TokenNumber, "999"),
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
	// AND the output should match the 'exp'
	doTest := func(op string, tk token.Token) {
		in := []rune(op)
		exp := []token.Lexeme{lex(tk, op)}
		happyTest(t, in, exp)
	}

	doTest("+", token.TokenAdd)
	doTest("-", token.TokenSub)
	doTest("*", token.TokenMul)
	doTest("/", token.TokenDiv)
}

func TestScanAll_4(t *testing.T) {

	// GIVEN parentheses
	// WHEN scanning all tokens
	// THEN the parentheses should be parsed without error
	// AND the output should match the 'exp'
	doTest := func(op string, tk token.Token) {
		in := []rune(op)
		exp := []token.Lexeme{lex(tk, op)}
		happyTest(t, in, exp)
	}

	doTest("(", token.TokenParenOpen)
	doTest(")", token.TokenParenClose)
}

func TestScanAll_100(t *testing.T) {

	// GIVEN a long expression
	r := token.NewRuneReader(
		[]rune("1 + 2 - 3 * 4 / 5"),
	)

	exp := []token.Lexeme{
		lex(token.TokenNumber, "1"),
		lex(token.TokenSpace, " "),
		lex(token.TokenAdd, "+"),
		lex(token.TokenSpace, " "),
		lex(token.TokenNumber, "2"),
		lex(token.TokenSpace, " "),
		lex(token.TokenSub, "-"),
		lex(token.TokenSpace, " "),
		lex(token.TokenNumber, "3"),
		lex(token.TokenSpace, " "),
		lex(token.TokenMul, "*"),
		lex(token.TokenSpace, " "),
		lex(token.TokenNumber, "4"),
		lex(token.TokenSpace, " "),
		lex(token.TokenDiv, "/"),
		lex(token.TokenSpace, " "),
		lex(token.TokenNumber, "5"),
	}

	// WHEN scanning all tokens
	act, e := ScanAll(r)

	// THEN the code should be correctly parsed without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func TestScanAll_101(t *testing.T) {

	// GIVEN multiple statements
	r := token.NewRuneReader(
		[]rune("1\n2\n3\n"),
	)

	exp := []token.Lexeme{
		lex(token.TokenNumber, "1"),
		lex(token.TokenNewline, "\n"),
		lex(token.TokenNumber, "2"),
		lex(token.TokenNewline, "\n"),
		lex(token.TokenNumber, "3"),
		lex(token.TokenNewline, "\n"),
	}

	// WHEN scanning all tokens
	act, e := ScanAll(r)

	// THEN the code should be correctly parsed without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func TestScanAll_102(t *testing.T) {

	// GIVEN firefly code containing an invalid token
	r := token.NewRuneReader(
		[]rune("#"),
	)

	// WHEN scanning the token
	_, e := ScanAll(r)

	// THEN an error should be returned
	require.NotNil(t, e, "Expected error when given invalid token")
}
