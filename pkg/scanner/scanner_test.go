package scanner

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func lex(tk token.Token, v string) token.Lexeme {
	return token.Lexeme{
		Token: tk,
		Value: v,
	}
}

func TestScanAll_1(t *testing.T) {

	// GIVEN valid firefly code containing valid numbers and operators
	sr := token.NewStringScrollReader(
		[]rune("1 + 2 - 3 * 4 / 5"),
	)

	exp := token.Statement{
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
	act, e := ScanAll(sr)

	// THEN the code should be correctly parsed without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func TestScanAll_2(t *testing.T) {

	// GIVEN valid firefly code containing a newline
	sr := token.NewStringScrollReader(
		[]rune("1\n2"),
	)

	exp := token.Statement{
		lex(token.TokenNumber, "1"),
		lex(token.TokenNewline, "\n"),
		lex(token.TokenNumber, "2"),
	}

	// WHEN scanning all tokens
	act, e := ScanAll(sr)

	// THEN the code should be correctly parsed without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func TestScanAll_3(t *testing.T) {

	// GIVEN firefly code containing an invalid token
	sr := token.NewStringScrollReader(
		[]rune("#"),
	)

	// WHEN scanning all tokens
	_, e := ScanAll(sr)

	// THEN an error should be returned
	require.NotNil(t, e, "Expected error when given invalid token")
}
