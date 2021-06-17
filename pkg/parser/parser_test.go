package parser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/ast"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func lex(tk token.Token, v string) token.Lexeme {
	return token.Lexeme{
		Token: tk,
		Value: v,
	}
}

func TestParseAll_1(t *testing.T) {

	// GIVEN a number statement
	lr := token.NewProgramReader(
		// 1
		token.Program{
			token.Statement{
				lex(token.TokenNumber, "1"),
			},
		},
	)

	exp := []ast.Tree{
		ast.Number{
			Value: 1,
		},
	}

	// WHEN parsing all statements
	act, e := ParseAll(lr)

	// THEN the number is parsed and returned without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}
