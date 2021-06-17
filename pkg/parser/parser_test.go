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

	// GIVEN a single digit number
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

	// GIVEN a multi-digit number
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

	// GIVEN a basic expression
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

	// THEN the number is parsed
	// AND returned without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func TestParseAll_4(t *testing.T) {

	// GIVEN a compound expression with equal operator precedence
	lr := token.NewProgramReader(
		// 1 + 2 - 3
		token.Program{
			token.Statement{
				lex(token.TokenNumber, "1"),
				lex(token.TokenAdd, "+"),
				lex(token.TokenNumber, "2"),
				lex(token.TokenSub, "-"),
				lex(token.TokenNumber, "3"),
			},
		},
	)

	exp := []ast.Node{
		ast.Sub{
			InfixOperation: ast.InfixOperation{
				Left: ast.Add{
					InfixOperation: ast.InfixOperation{
						Left:  ast.Number{Value: 1},
						Right: ast.Number{Value: 2},
					},
				},
				Right: ast.Number{Value: 3},
			},
		},
	}

	// WHEN parsing all statements
	act, e := ParseAll(lr)

	// THEN the number is parsed
	// AND returned without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func TestParseAll_5(t *testing.T) {

	// GIVEN a compound expression
	// AND the latter operator having a higher precedence
	lr := token.NewProgramReader(
		// 1 + 2 * 3
		// 1 + (2 * 3)
		token.Program{
			token.Statement{
				lex(token.TokenNumber, "1"),
				lex(token.TokenAdd, "+"),
				lex(token.TokenNumber, "2"),
				lex(token.TokenMul, "*"),
				lex(token.TokenNumber, "3"),
			},
		},
	)

	exp := []ast.Node{
		ast.Add{
			InfixOperation: ast.InfixOperation{
				Left: ast.Number{Value: 1},
				Right: ast.Mul{
					InfixOperation: ast.InfixOperation{
						Left:  ast.Number{Value: 2},
						Right: ast.Number{Value: 3},
					},
				},
			},
		},
	}

	// WHEN parsing all statements
	act, e := ParseAll(lr)

	// THEN the number is parsed
	// AND returned without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func TestParseAll_6(t *testing.T) {

	// GIVEN a long compound expression
	lr := token.NewProgramReader(
		// 9 / 3 + 2 * 3
		// (9 / 3) + (2 * 3)
		token.Program{
			token.Statement{
				lex(token.TokenNumber, "9"),
				lex(token.TokenDiv, "/"),
				lex(token.TokenNumber, "3"),
				lex(token.TokenAdd, "+"),
				lex(token.TokenNumber, "2"),
				lex(token.TokenMul, "*"),
				lex(token.TokenNumber, "3"),
			},
		},
	)

	exp := []ast.Node{
		ast.Add{
			InfixOperation: ast.InfixOperation{
				Left: ast.Div{
					InfixOperation: ast.InfixOperation{
						Left:  ast.Number{Value: 9},
						Right: ast.Number{Value: 3},
					},
				},
				Right: ast.Mul{
					InfixOperation: ast.InfixOperation{
						Left:  ast.Number{Value: 2},
						Right: ast.Number{Value: 3},
					},
				},
			},
		},
	}

	// WHEN parsing all statements
	act, e := ParseAll(lr)

	// THEN the number is parsed
	// AND returned without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}
