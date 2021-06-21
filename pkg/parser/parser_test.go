package parser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/ast"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func happyTest(t *testing.T, p token.Program, exp ast.Program) {
	pr := token.NewProgramReader(p)
	act, e := ParseAll(pr)
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func lex(tk token.Token, v string) token.Lexeme {
	return token.Lexeme{
		Token: tk,
		Value: v,
	}
}

func num(n int64) ast.NumberNode {
	return ast.NumberNode{
		Value: n,
	}
}

func infix(t ast.AST, left, right ast.Node) ast.InfixNode {
	return ast.InfixNode{
		AST:   t,
		Left:  left,
		Right: right,
	}
}

func TestParseAll_0(t *testing.T) {

	// GIVEN an empty statement
	p := token.Program{
		token.Statement{},
	}

	exp := ast.Program{
		ast.EmptyNode{},
	}

	// WHEN parsing all statements
	// THEN an empty node is returned without error
	happyTest(t, p, exp)
}

func TestParseAll_1(t *testing.T) {

	// GIVEN a single digit number
	// 9
	p := token.Program{
		token.Statement{
			lex(token.TokenNumber, "9"),
		},
	}

	exp := ast.Program{
		num(9),
	}

	// WHEN parsing all statements
	// THEN the number is parsed and returned without error
	happyTest(t, p, exp)
}

func TestParseAll_2(t *testing.T) {

	// GIVEN a multi-digit number
	// 99
	p := token.Program{
		token.Statement{
			lex(token.TokenNumber, "99"),
		},
	}

	exp := ast.Program{
		num(99),
	}

	// WHEN parsing all statements
	// THEN the number is parsed and returned without error
	happyTest(t, p, exp)
}

func TestParseAll_3(t *testing.T) {

	// GIVEN a basic expression
	// 1 + 2
	p := token.Program{
		token.Statement{
			lex(token.TokenNumber, "1"),
			lex(token.TokenAdd, "+"),
			lex(token.TokenNumber, "2"),
		},
	}

	exp := ast.Program{
		infix(ast.AstAdd, num(1), num(2)),
	}

	// WHEN parsing all statements
	// THEN the number is parsed
	// AND returned without error
	happyTest(t, p, exp)
}

func TestParseAll_4(t *testing.T) {

	// GIVEN a compound expression with equal operator precedence
	// 1 + 2 - 3
	p := token.Program{
		token.Statement{
			lex(token.TokenNumber, "1"),
			lex(token.TokenAdd, "+"),
			lex(token.TokenNumber, "2"),
			lex(token.TokenSub, "-"),
			lex(token.TokenNumber, "3"),
		},
	}

	exp := ast.Program{
		infix(ast.AstSub,
			infix(ast.AstAdd, num(1), num(2)),
			num(3),
		),
	}

	// WHEN parsing all statements
	// THEN the number is parsed
	// AND returned without error
	happyTest(t, p, exp)
}

func TestParseAll_5(t *testing.T) {

	// GIVEN a compound expression
	// AND the latter operator having a higher precedence

	// 1 + 2 * 3
	// 1 + (2 * 3)
	p := token.Program{
		token.Statement{
			lex(token.TokenNumber, "1"),
			lex(token.TokenAdd, "+"),
			lex(token.TokenNumber, "2"),
			lex(token.TokenMul, "*"),
			lex(token.TokenNumber, "3"),
		},
	}

	exp := ast.Program{
		infix(ast.AstAdd,
			num(1),
			infix(ast.AstMul, num(2), num(3)),
		),
	}

	// WHEN parsing all statements
	// THEN the number is parsed
	// AND returned without error
	happyTest(t, p, exp)
}

func TestParseAll_6(t *testing.T) {

	// GIVEN a long compound expression
	// 9 / 3 + 2 * 3
	// (9 / 3) + (2 * 3)
	p := token.Program{
		token.Statement{
			lex(token.TokenNumber, "9"),
			lex(token.TokenDiv, "/"),
			lex(token.TokenNumber, "3"),
			lex(token.TokenAdd, "+"),
			lex(token.TokenNumber, "2"),
			lex(token.TokenMul, "*"),
			lex(token.TokenNumber, "3"),
		},
	}

	exp := ast.Program{
		infix(ast.AstAdd,
			infix(ast.AstDiv, num(9), num(3)),
			infix(ast.AstMul, num(2), num(3)),
		),
	}

	// WHEN parsing all statements
	// THEN the number is parsed
	// AND returned without error
	happyTest(t, p, exp)
}

func TestParseAll_7(t *testing.T) {

	// GIVEN a long compound expression
	// 8 + 4 / 3 * 3 - 2 * 5
	// (8 + ((4 / 3) * 3)) - (2 * 5)
	p := token.Program{
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
	}

	// 4 / 3
	ex1 := infix(ast.AstDiv, num(4), num(3))

	// (4 / 3) * 3
	ex2 := infix(ast.AstMul, ex1, num(3))

	// 8 + (4 / 3 * 3)
	ex3 := infix(ast.AstAdd, num(8), ex2)

	// 2 * 5
	ex4 := infix(ast.AstMul, num(2), num(5))

	// (8 + 4 / 3 * 3) - (2 * 5)
	exp := ast.Program{
		infix(ast.AstSub, ex3, ex4),
	}

	// WHEN parsing all statements
	// THEN the expression is parsed
	// AND returned without error
	happyTest(t, p, exp)
}

func TestParseAll_8(t *testing.T) {

	// GIVEN an expression with parentheses
	// (9)
	p := token.Program{
		token.Statement{
			lex(token.TokenParenOpen, "("),
			lex(token.TokenNumber, "9"),
			lex(token.TokenParenClose, ")"),
		},
	}

	exp := ast.Program{
		num(9),
	}

	// WHEN parsing all statements
	// THEN the number is parsed
	// AND returned without error
	happyTest(t, p, exp)
}

func TestParseAll_9(t *testing.T) {

	// GIVEN a long compound expression with parentheses
	// 8 + (4 / 3 * (3 - 2)) * 5
	// 8 + (((4 / 3) * (3 - 2)) * 5)
	p := token.Program{
		token.Statement{
			lex(token.TokenNumber, "8"),
			lex(token.TokenAdd, "+"),
			lex(token.TokenParenOpen, "("),
			lex(token.TokenNumber, "4"),
			lex(token.TokenDiv, "/"),
			lex(token.TokenNumber, "3"),
			lex(token.TokenMul, "*"),
			lex(token.TokenParenOpen, "("),
			lex(token.TokenNumber, "3"),
			lex(token.TokenSub, "-"),
			lex(token.TokenNumber, "2"),
			lex(token.TokenParenClose, ")"),
			lex(token.TokenParenClose, ")"),
			lex(token.TokenMul, "*"),
			lex(token.TokenNumber, "5"),
		},
	}

	// 4 / 3
	ex1 := infix(ast.AstDiv, num(4), num(3))

	// 3 - 2)
	ex2 := infix(ast.AstSub, num(3), num(2))

	// (4 / 3) * (3 - 2)
	ex3 := infix(ast.AstMul, ex1, ex2)

	// ... * 5
	ex4 := infix(ast.AstMul, ex3, num(5))

	// 8 + ...
	exp := ast.Program{
		infix(ast.AstAdd, num(8), ex4),
	}

	// WHEN parsing all statements
	// THEN the expression is parsed
	// AND returned without error
	happyTest(t, p, exp)
}
