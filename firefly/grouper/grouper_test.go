package grouper

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/firefly/token"
)

func lex(tk token.Token, v string) token.Lexeme {
	return token.Lexeme{
		Token: tk,
		Value: v,
	}
}

func TestGroupAll_0(t *testing.T) {

	// GIVEN nothing
	r := token.NewLexReader(nil)

	var exp token.Block

	// WHEN grouping all lexemes into statements
	act, e := GroupAll(r)

	// THEN no error is returned
	require.Nil(t, e, "%+v", e)

	// AND the output is a nil slice of statements (token.block)
	require.Equal(t, exp, act)
}

func TestGroupAll_1(t *testing.T) {

	// GIVEN three statements in a lexeme reader
	r := token.NewLexReader(
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

	exp := token.Block{
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

	// THEN no error is returned
	require.Nil(t, e, "%+v", e)

	// AND three statments are separated and returned
	require.Equal(t, exp, act)
}

func TestGroupAll_2(t *testing.T) {

	// GIVEN a single statement in a lexeme reader
	r := token.NewLexReader(
		// 1 + 2
		token.Statement{
			lex(token.TK_NUMBER, "1"),
			lex(token.TK_ADD, "+"),
			lex(token.TK_NUMBER, "2"),
		},
	)

	exp := token.Block{
		token.Statement{
			lex(token.TK_NUMBER, "1"),
			lex(token.TK_ADD, "+"),
			lex(token.TK_NUMBER, "2"),
		},
	}

	// WHEN grouping all lexemes into statements
	act, e := GroupAll(r)

	// THEN no error is returned
	require.Nil(t, e, "%+v", e)

	// AND one statment is returned
	require.Equal(t, exp, act)
}
