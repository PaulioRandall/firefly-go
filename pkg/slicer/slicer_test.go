package slicer

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

func TestSliceAll_1(t *testing.T) {
	// GIVEN two statements as a slice of lexemes
	// WHEN slicing all lexemes into statements
	// THEN the two statments are separated and returned without error

	lr := NewSliceLexemeReader(
		// 1 + 2
		// 3 * 4
		[]token.Lexeme{
			lex(token.TokenNumber, "1"),
			lex(token.TokenNewline, "\n"),
			lex(token.TokenNumber, "2"),
			lex(token.TokenNewline, "\n"),
			lex(token.TokenNumber, "3"),
		},
	)

	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.TokenNumber, "1"),
		},
		[]token.Lexeme{
			lex(token.TokenNumber, "2"),
		},
		[]token.Lexeme{
			lex(token.TokenNumber, "3"),
		},
	}

	act, e := SliceAll(lr)

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func TestSliceAll_2(t *testing.T) {
	// GIVEN a single statement as a slice of lexemes
	// WHEN slicing all lexemes into statements
	// THEN the one statments is returned without error

	lr := NewSliceLexemeReader(
		// 1 + 2
		// 3 * 4
		[]token.Lexeme{
			lex(token.TokenNumber, "1"),
			lex(token.TokenOperator, "+"),
			lex(token.TokenNumber, "2"),
		},
	)

	exp := [][]token.Lexeme{
		[]token.Lexeme{
			lex(token.TokenNumber, "1"),
			lex(token.TokenOperator, "+"),
			lex(token.TokenNumber, "2"),
		},
	}

	act, e := SliceAll(lr)

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}
