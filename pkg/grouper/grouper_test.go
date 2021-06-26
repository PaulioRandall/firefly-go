package grouper

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

func TestGroupAll_1(t *testing.T) {

	// GIVEN three statements in a lexeme reader
	r := token.NewLexemeReader(
		// 1
		// 2
		// 3
		token.Statement{
			lex(token.TK_NUMBER, "1"),
			lex(token.TK_NEWLINE, "\n"),
			lex(token.TK_NUMBER, "2"),
			lex(token.TK_NEWLINE, "\n"),
			lex(token.TK_NUMBER, "3"),
		},
	)

	exp := token.Program{
		token.Statement{
			lex(token.TK_NUMBER, "1"),
		},
		token.Statement{
			lex(token.TK_NUMBER, "2"),
		},
		token.Statement{
			lex(token.TK_NUMBER, "3"),
		},
	}

	// WHEN grouping all lexemes into statements
	act, e := GroupAll(r)

	// THEN the two statments are separated and returned without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func TestGroupAll_2(t *testing.T) {

	// GIVEN a single statement in a lexeme reader
	r := token.NewLexemeReader(
		// 1 + 2
		token.Statement{
			lex(token.TK_NUMBER, "1"),
			lex(token.TK_ADD, "+"),
			lex(token.TK_NUMBER, "2"),
		},
	)

	exp := token.Program{
		token.Statement{
			lex(token.TK_NUMBER, "1"),
			lex(token.TK_ADD, "+"),
			lex(token.TK_NUMBER, "2"),
		},
	}

	// WHEN grouping all lexemes into statements
	act, e := GroupAll(r)

	// THEN the one statments is returned without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}
