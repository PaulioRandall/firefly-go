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

	// GIVEN a statement with redundant tokens
	lr := token.NewSliceStmtReader(
		// 1 + 2
		[]token.Statement{
			token.Statement{
				lex(token.TokenNumber, "1"),
				lex(token.TokenSpace, " "),
				lex(token.TokenOperator, "+"),
				lex(token.TokenSpace, " "),
				lex(token.TokenNumber, "2"),
			},
		},
	)

	exp := []token.Statement{
		token.Statement{
			lex(token.TokenNumber, "1"),
			lex(token.TokenOperator, "+"),
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
