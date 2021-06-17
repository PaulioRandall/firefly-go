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

	// GIVEN a single digit number statement
	lr := token.NewProgramReader(
		// 9
		token.Program{
			token.Statement{
				lex(token.TokenNumber, "9"),
			},
		},
	)

	exp := []ast.Node{
		ast.Number{
			Value: 9,
		},
	}

	// WHEN parsing all statements
	act, e := ParseAll(lr)

	// THEN the number is parsed and returned without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func TestParseAll_2(t *testing.T) {

	// GIVEN a multi-digit number statement
	lr := token.NewProgramReader(
		// 99
		token.Program{
			token.Statement{
				lex(token.TokenNumber, "99"),
			},
		},
	)

	exp := []ast.Node{
		ast.Number{
			Value: 99,
		},
	}

	// WHEN parsing all statements
	act, e := ParseAll(lr)

	// THEN the number is parsed and returned without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func TestParseAll_3(t *testing.T) {

	// GIVEN an operation statement
	lr := token.NewProgramReader(
		// 1 + 2
		token.Program{
			token.Statement{
				lex(token.TokenNumber, "1"),
				lex(token.TokenAdd, "+"),
				lex(token.TokenNumber, "2"),
			},
		},
	)

	exp := []ast.Node{
		ast.Add{
			InfixOperation: ast.InfixOperation{
				Left:  ast.Number{Value: 1},
				Right: ast.Number{Value: 2},
			},
		},
	}

	// WHEN parsing all statements
	act, e := ParseAll(lr)

	// THEN the number is parsed and returned without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}
