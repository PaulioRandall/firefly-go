package slicer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/token"
)

type mockTokenReader struct {
	idx int
	tks []token.Lexeme
}

func (tr *mockTokenReader) More() bool {
	return len(tr.tks) > tr.idx
}

func (tr *mockTokenReader) Read() (token.Lexeme, error) {
	if !tr.More() {
		return token.Lexeme{}, errors.New("EOF")
	}
	tk := tr.tks[tr.idx]
	tr.idx++
	return tk, nil
}

func lex(tk token.Token, v string) token.Lexeme {
	return token.Lexeme{
		Token: tk,
		Value: v,
	}
}

func TestSliceAll_1(t *testing.T) {
	// GIVEN two statements as a slice of tokens
	// WHEN slicing all tokens into statements
	// THEN the two statments are separated and returned without error

	tr := &mockTokenReader{
		// 1 + 2
		// 3 * 4
		tks: []token.Lexeme{
			lex(token.TokenNumber, "1"),
			lex(token.TokenNewline, "\n"),
			lex(token.TokenNumber, "2"),
		},
	}

	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.TokenNumber, "1"),
		},
		[]token.Lexeme{
			lex(token.TokenNumber, "2"),
		},
	}

	act, e := SliceAll(tr)

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}
