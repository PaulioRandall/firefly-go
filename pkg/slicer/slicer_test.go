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

	// GIVEN three statements as a slice of lexemes
	lr := NewSliceLexemeReader(
		// 1
		// 2
		// 3
		Statement{
			lex(token.TokenNumber, "1"),
			lex(token.TokenNewline, "\n"),
			lex(token.TokenNumber, "2"),
			lex(token.TokenNewline, "\n"),
			lex(token.TokenNumber, "3"),
		},
	)

	exp := []Statement{
		Statement{
			lex(token.TokenNumber, "1"),
		},
		Statement{
			lex(token.TokenNumber, "2"),
		},
		Statement{
			lex(token.TokenNumber, "3"),
		},
	}

	// WHEN slicing all lexemes into statements
	act, e := SliceAll(lr)

	// THEN the two statments are separated and returned without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func TestSliceAll_2(t *testing.T) {

	// GIVEN a single statement as a slice of lexemes
	lr := NewSliceLexemeReader(
		// 1 + 2
		Statement{
			lex(token.TokenNumber, "1"),
			lex(token.TokenOperator, "+"),
			lex(token.TokenNumber, "2"),
		},
	)

	exp := []Statement{
		Statement{
			lex(token.TokenNumber, "1"),
			lex(token.TokenOperator, "+"),
			lex(token.TokenNumber, "2"),
		},
	}

	// WHEN slicing all lexemes into statements
	act, e := SliceAll(lr)

	// THEN the one statments is returned without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}
