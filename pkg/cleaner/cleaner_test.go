package cleaner

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

func TestCleanAll_1(t *testing.T) {

	// GIVEN multiple statements with redundant tokens
	lr := token.NewProgramReader(
		// 1 + 1
		// 2 * 2
		token.Program{
			token.Statement{
				lex(token.TokenNumber, "1"),
				lex(token.TokenSpace, " "),
				lex(token.TokenOperator, "+"),
				lex(token.TokenSpace, " "),
				lex(token.TokenNumber, "1"),
			},
			token.Statement{
				lex(token.TokenNumber, "2"),
				lex(token.TokenSpace, " "),
				lex(token.TokenOperator, "*"),
				lex(token.TokenSpace, " "),
				lex(token.TokenNumber, "2"),
			},
		},
	)

	exp := token.Program{
		token.Statement{
			lex(token.TokenNumber, "1"),
			lex(token.TokenOperator, "+"),
			lex(token.TokenNumber, "1"),
		},
		token.Statement{
			lex(token.TokenNumber, "2"),
			lex(token.TokenOperator, "*"),
			lex(token.TokenNumber, "2"),
		},
	}

	// WHEN cleaning all statements
	act, e := CleanAll(lr)

	// THEN the redundant tokens are removed and the rest of the statement is
	// returned without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func TestCleanAll_2(t *testing.T) {

	// GIVEN a statement without redundant tokens
	lr := token.NewProgramReader(
		// 1 + 1
		token.Program{
			token.Statement{
				lex(token.TokenNumber, "1"),
				lex(token.TokenOperator, "+"),
				lex(token.TokenNumber, "1"),
			},
		},
	)

	exp := token.Program{
		token.Statement{
			lex(token.TokenNumber, "1"),
			lex(token.TokenOperator, "+"),
			lex(token.TokenNumber, "1"),
		},
	}

	// WHEN cleaning all statements
	act, e := CleanAll(lr)

	// THEN the statement is returned unchanged and without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}
