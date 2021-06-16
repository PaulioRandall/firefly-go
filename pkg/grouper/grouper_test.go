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

	// GIVEN three statements in a lexeme reader
	lr := token.NewSliceLexemeReader(
		// 1
		// 2
		// 3
		token.Statement{
			lex(token.TokenNumber, "1"),
			lex(token.TokenNewline, "\n"),
			lex(token.TokenNumber, "2"),
			lex(token.TokenNewline, "\n"),
			lex(token.TokenNumber, "3"),
		},
	)

	exp := []token.Statement{
		token.Statement{
			lex(token.TokenNumber, "1"),
		},
		token.Statement{
			lex(token.TokenNumber, "2"),
		},
		token.Statement{
			lex(token.TokenNumber, "3"),
		},
	}

	// WHEN grouping all lexemes into statements
	act, e := GroupAll(lr)

	// THEN the two statments are separated and returned without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func TestSliceAll_2(t *testing.T) {

	// GIVEN a single statement in a lexeme reader
	lr := token.NewSliceLexemeReader(
		// 1 + 2
		token.Statement{
			lex(token.TokenNumber, "1"),
			lex(token.TokenOperator, "+"),
			lex(token.TokenNumber, "2"),
		},
	)

	exp := []token.Statement{
		token.Statement{
			lex(token.TokenNumber, "1"),
			lex(token.TokenOperator, "+"),
			lex(token.TokenNumber, "2"),
		},
	}

	// WHEN grouping all lexemes into statements
	act, e := GroupAll(lr)

	// THEN the one statments is returned without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}
