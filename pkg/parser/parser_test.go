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

func num(n int64) ast.Number {
	return ast.Number{
		Value: n,
	}
}

func infix(t ast.AST, left, right ast.Node) ast.InfixExprNode {
	return ast.InfixExprNode{
		AST:   t,
		Left:  left,
		Right: right,
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
		num(9),
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
		num(99),
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
		infix(ast.AstAdd,
			num(1),
			num(2),
		),
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
		infix(ast.AstSub,
			infix(ast.AstAdd,
				num(1),
				num(2),
			),
			num(3),
		),
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
		infix(ast.AstAdd,
			num(1),
			infix(ast.AstMul,
				num(2),
				num(3),
			),
		),
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
		infix(ast.AstAdd,
			infix(ast.AstDiv,
				num(9),
				num(3),
			),
			infix(ast.AstMul,
				num(2),
				num(3),
			),
		),
	}

	// WHEN parsing all statements
	act, e := ParseAll(lr)

	// THEN the number is parsed
	// AND returned without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func TestParseAll_7(t *testing.T) {

	// GIVEN a long compound expression
	lr := token.NewProgramReader(
		// 8 + 4 / 3 * 3 - 2 * 5
		// (8 + ((4 / 3) * 3)) - (2 * 5)
		token.Program{
			token.Statement{
				lex(token.TokenNumber, "8"),
				lex(token.TokenAdd, "+"),
				lex(token.TokenNumber, "4"),
				lex(token.TokenDiv, "/"),
				lex(token.TokenNumber, "3"),
				lex(token.TokenMul, "*"),
				lex(token.TokenNumber, "3"),
				lex(token.TokenSub, "-"),
				lex(token.TokenNumber, "2"),
				lex(token.TokenMul, "*"),
				lex(token.TokenNumber, "5"),
			},
		},
	)

	// [4 / 3]
	ex1 := infix(ast.AstDiv,
		num(4),
		num(3),
	)

	// [4 / 3 * 3]
	ex2 := infix(ast.AstMul,
		ex1,
		num(3),
	)

	// [8 + 4 / 3 * 3]
	ex3 := infix(ast.AstAdd,
		num(8),
		ex2,
	)

	// [2 * 5]
	ex4 := infix(ast.AstMul,
		num(2),
		num(5),
	)

	// [8 + 4 / 3 * 3 - 2 * 5]
	exp := []ast.Node{
		infix(ast.AstSub, ex3, ex4),
	}

	// WHEN parsing all statements
	act, e := ParseAll(lr)

	// THEN the expression is parsed
	// AND returned without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}
