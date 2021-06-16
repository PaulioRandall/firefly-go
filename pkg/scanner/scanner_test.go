package scanner

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func doTestScanAll(t *testing.T, scroll string, exp []token.Lexeme) {

	sr := NewStringScrollReader(
		[]rune(scroll),
	)

	act, e := ScanAll(sr)

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
	// GIVEN a valid firefly scroll containing valid numbers and operators
	// WHEN scanning all tokens in the scroll
	// THEN the scroll should be correctly parsed without error

	scroll := "1 + 2 - 3 * 4 / 5"

	exp := []token.Lexeme{
		lex(token.TokenNumber, "1"),
		lex(token.TokenSpace, " "),
		lex(token.TokenOperator, "+"),
		lex(token.TokenSpace, " "),
		lex(token.TokenNumber, "2"),
		lex(token.TokenSpace, " "),
		lex(token.TokenOperator, "-"),
		lex(token.TokenSpace, " "),
		lex(token.TokenNumber, "3"),
		lex(token.TokenSpace, " "),
		lex(token.TokenOperator, "*"),
		lex(token.TokenSpace, " "),
		lex(token.TokenNumber, "4"),
		lex(token.TokenSpace, " "),
		lex(token.TokenOperator, "/"),
		lex(token.TokenSpace, " "),
		lex(token.TokenNumber, "5"),
	}

	doTestScanAll(t, scroll, exp)
}

func TestScanAll_2(t *testing.T) {
	// GIVEN a valid firefly scroll containing a newline
	// WHEN scanning all tokens in the scroll
	// THEN the scroll should be correctly parsed without error

	scroll := "1\n2"

	exp := []token.Lexeme{
		lex(token.TokenNumber, "1"),
		lex(token.TokenNewline, "\n"),
		lex(token.TokenNumber, "2"),
	}

	doTestScanAll(t, scroll, exp)
}

func TestScanAll_3(t *testing.T) {
	// GIVEN a firefly scroll containing an invalid token
	// WHEN scanning all tokens in the scroll
	// THEN the an error should be returned

	sr := NewStringScrollReader(
		[]rune("#"),
	)

	_, e := ScanAll(sr)

	require.NotNil(t, e, "Expected error when given invalid token")
}
