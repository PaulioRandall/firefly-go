package scanner

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/token"
)

type mockScrollReader struct {
	idx    int
	scroll []rune
}

func (sr *mockScrollReader) More() bool {
	return len(sr.scroll) > sr.idx
}

func (sr *mockScrollReader) Read() (rune, error) {
	if !sr.More() {
		return rune(0), errors.New("EOF")
	}
	ru := sr.scroll[sr.idx]
	sr.idx++
	return ru, nil
}

func (sr *mockScrollReader) PutBack(ru rune) error {
	last := sr.scroll[sr.idx:]
	first := append(sr.scroll[:sr.idx], ru)
	sr.scroll = append(first, last...)
	return nil
}

func doTestScanAll(t *testing.T, scroll string, exp []token.Lexeme) {

	sr := &mockScrollReader{
		scroll: []rune(scroll),
	}

	act, e := ScanAll(sr)

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func TestScanAll_1(t *testing.T) {
	// GIVEN a valid firefly scroll containing valid numbers and operators
	// THEN the scroll should be correctly parsed without error

	scroll := "1 + 2 - 3 * 4 / 5"

	exp := []token.Lexeme{
		token.Lexeme{token.TokenNumber, "1"},
		token.Lexeme{token.TokenSpace, " "},
		token.Lexeme{token.TokenOperator, "+"},
		token.Lexeme{token.TokenSpace, " "},
		token.Lexeme{token.TokenNumber, "2"},
		token.Lexeme{token.TokenSpace, " "},
		token.Lexeme{token.TokenOperator, "-"},
		token.Lexeme{token.TokenSpace, " "},
		token.Lexeme{token.TokenNumber, "3"},
		token.Lexeme{token.TokenSpace, " "},
		token.Lexeme{token.TokenOperator, "*"},
		token.Lexeme{token.TokenSpace, " "},
		token.Lexeme{token.TokenNumber, "4"},
		token.Lexeme{token.TokenSpace, " "},
		token.Lexeme{token.TokenOperator, "/"},
		token.Lexeme{token.TokenSpace, " "},
		token.Lexeme{token.TokenNumber, "5"},
	}

	doTestScanAll(t, scroll, exp)
}

func TestScanAll_2(t *testing.T) {
	// GIVEN a valid firefly scroll containing a newline
	// THEN the scroll should be correctly parsed without error

	scroll := "1\n2"

	exp := []token.Lexeme{
		token.Lexeme{token.TokenNumber, "1"},
		token.Lexeme{token.TokenNewline, "\n"},
		token.Lexeme{token.TokenNumber, "2"},
	}

	doTestScanAll(t, scroll, exp)
}

func TestScanAll_3(t *testing.T) {
	// GIVEN a firefly scroll containing an invalid token
	// THEN the an error should be returned

	sr := &mockScrollReader{
		scroll: []rune("#"),
	}

	_, e := ScanAll(sr)

	require.NotNil(t, e, "Expected error when given invalid token")
}
