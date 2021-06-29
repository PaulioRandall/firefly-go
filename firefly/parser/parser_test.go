package parser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/firefly/ast"
	"github.com/PaulioRandall/firefly-go/firefly/token"
)

func happyTest(t *testing.T, b token.Block, exp ast.Block) {
	pr := token.NewStmtReader(b)
	act, e := ParseAll(pr)
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func unhappyTest(t *testing.T, b token.Block) {
	pr := token.NewStmtReader(b)
	_, e := ParseAll(pr)
	require.NotNil(t, e, "Expected error")
}

func lex(tk token.Token, v string) token.Lexeme {
	return token.Lexeme{
		Token: tk,
		Value: v,
	}
}

func num(n int64) ast.NumberTree {
	return ast.NumberTree{
		Value: n,
	}
}

func infix(n ast.Node, left, right ast.Tree) ast.InfixTree {
	return ast.InfixTree{
		Node:  n,
		Left:  left,
		Right: right,
	}
}

func TestParseAll_0(t *testing.T) {

	// GIVEN nothing
	var b token.Block

	exp := ast.Block{
		ast.EmptyTree{},
	}

	// WHEN parsing the statement
	// THEN an empty node is returned without error
	happyTest(t, b, exp)
}

func TestParseAll_1(t *testing.T) {

	// GIVEN a single digit number
	// 9
	b := token.Block{
		token.Statement{
			lex(token.TK_NUMBER, "9"),
		},
	}

	exp := ast.Block{
		num(9),
	}

	// WHEN parsing the statement
	// THEN the number is parsed and returned without error
	happyTest(t, b, exp)
}

func TestParseAll_2(t *testing.T) {

	// GIVEN a multi-digit number
	// 99
	b := token.Block{
		token.Statement{
			lex(token.TK_NUMBER, "99"),
		},
	}

	exp := ast.Block{
		num(99),
	}

	// WHEN parsing the statement
	// THEN the number is parsed and returned without error
	happyTest(t, b, exp)
}

func TestParseAll_3(t *testing.T) {

	// GIVEN a basic expression
	// 1 + 2
	b := token.Block{
		token.Statement{
			lex(token.TK_NUMBER, "1"),
			lex(token.TK_ADD, "+"),
			lex(token.TK_NUMBER, "2"),
		},
	}

	exp := ast.Block{
		infix(ast.NODE_ADD, num(1), num(2)),
	}

	// WHEN parsing the statement
	// THEN the number is parsed
	// AND returned without error
	happyTest(t, b, exp)
}

func TestParseAll_4(t *testing.T) {

	// GIVEN a compound expression with equal operator precedence
	// 1 + 2 - 3
	b := token.Block{
		token.Statement{
			lex(token.TK_NUMBER, "1"),
			lex(token.TK_ADD, "+"),
			lex(token.TK_NUMBER, "2"),
			lex(token.TK_SUB, "-"),
			lex(token.TK_NUMBER, "3"),
		},
	}

	exp := ast.Block{
		infix(ast.NODE_SUB,
			infix(ast.NODE_ADD, num(1), num(2)),
			num(3),
		),
	}

	// WHEN parsing the statement
	// THEN the number is parsed
	// AND returned without error
	happyTest(t, b, exp)
}

func TestParseAll_5(t *testing.T) {

	// GIVEN a compound expression
	// AND the latter operator having a higher precedence

	// 1 + 2 * 3
	// 1 + (2 * 3)
	b := token.Block{
		token.Statement{
			lex(token.TK_NUMBER, "1"),
			lex(token.TK_ADD, "+"),
			lex(token.TK_NUMBER, "2"),
			lex(token.TK_MUL, "*"),
			lex(token.TK_NUMBER, "3"),
		},
	}

	exp := ast.Block{
		infix(ast.NODE_ADD,
			num(1),
			infix(ast.NODE_MUL, num(2), num(3)),
		),
	}

	// WHEN parsing the statement
	// THEN the number is parsed
	// AND returned without error
	happyTest(t, b, exp)
}

func TestParseAll_6(t *testing.T) {

	// GIVEN a long compound expression
	// 9 / 3 + 2 * 3
	// (9 / 3) + (2 * 3)
	b := token.Block{
		token.Statement{
			lex(token.TK_NUMBER, "9"),
			lex(token.TK_DIV, "/"),
			lex(token.TK_NUMBER, "3"),
			lex(token.TK_ADD, "+"),
			lex(token.TK_NUMBER, "2"),
			lex(token.TK_MUL, "*"),
			lex(token.TK_NUMBER, "3"),
		},
	}

	exp := ast.Block{
		infix(ast.NODE_ADD,
			infix(ast.NODE_DIV, num(9), num(3)),
			infix(ast.NODE_MUL, num(2), num(3)),
		),
	}

	// WHEN parsing the statement
	// THEN the number is parsed
	// AND returned without error
	happyTest(t, b, exp)
}

func TestParseAll_7(t *testing.T) {

	// GIVEN a long compound expression
	// 8 + 4 / 3 * 3 - 2 * 5
	// (8 + ((4 / 3) * 3)) - (2 * 5)
	b := token.Block{
		token.Statement{
			lex(token.TK_NUMBER, "8"),
			lex(token.TK_ADD, "+"),
			lex(token.TK_NUMBER, "4"),
			lex(token.TK_DIV, "/"),
			lex(token.TK_NUMBER, "3"),
			lex(token.TK_MUL, "*"),
			lex(token.TK_NUMBER, "3"),
			lex(token.TK_SUB, "-"),
			lex(token.TK_NUMBER, "2"),
			lex(token.TK_MUL, "*"),
			lex(token.TK_NUMBER, "5"),
		},
	}

	// 4 / 3
	ex1 := infix(ast.NODE_DIV, num(4), num(3))

	// (4 / 3) * 3
	ex2 := infix(ast.NODE_MUL, ex1, num(3))

	// 8 + (4 / 3 * 3)
	ex3 := infix(ast.NODE_ADD, num(8), ex2)

	// 2 * 5
	ex4 := infix(ast.NODE_MUL, num(2), num(5))

	// (8 + 4 / 3 * 3) - (2 * 5)
	exp := ast.Block{
		infix(ast.NODE_SUB, ex3, ex4),
	}

	// WHEN parsing the statement
	// THEN the expression is parsed
	// AND returned without error
	happyTest(t, b, exp)
}

func TestParseAll_8(t *testing.T) {

	// GIVEN an expression with parentheses
	// (9)
	b := token.Block{
		token.Statement{
			lex(token.TK_PAREN_OPEN, "("),
			lex(token.TK_NUMBER, "9"),
			lex(token.TK_PAREN_CLOSE, ")"),
		},
	}

	exp := ast.Block{
		num(9),
	}

	// WHEN parsing the statement
	// THEN the number is parsed
	// AND returned without error
	happyTest(t, b, exp)
}

func TestParseAll_9(t *testing.T) {

	// GIVEN a long compound expression with parentheses
	// 8 + (4 / 3 * (3 - 2)) * 5
	// 8 + (((4 / 3) * (3 - 2)) * 5)
	b := token.Block{
		token.Statement{
			lex(token.TK_NUMBER, "8"),
			lex(token.TK_ADD, "+"),
			lex(token.TK_PAREN_OPEN, "("),
			lex(token.TK_NUMBER, "4"),
			lex(token.TK_DIV, "/"),
			lex(token.TK_NUMBER, "3"),
			lex(token.TK_MUL, "*"),
			lex(token.TK_PAREN_OPEN, "("),
			lex(token.TK_NUMBER, "3"),
			lex(token.TK_SUB, "-"),
			lex(token.TK_NUMBER, "2"),
			lex(token.TK_PAREN_CLOSE, ")"),
			lex(token.TK_PAREN_CLOSE, ")"),
			lex(token.TK_MUL, "*"),
			lex(token.TK_NUMBER, "5"),
		},
	}

	// 4 / 3
	ex1 := infix(ast.NODE_DIV, num(4), num(3))

	// 3 - 2)
	ex2 := infix(ast.NODE_SUB, num(3), num(2))

	// (4 / 3) * (3 - 2)
	ex3 := infix(ast.NODE_MUL, ex1, ex2)

	// ... * 5
	ex4 := infix(ast.NODE_MUL, ex3, num(5))

	// 8 + ...
	exp := ast.Block{
		infix(ast.NODE_ADD, num(8), ex4),
	}

	// WHEN parsing the statement
	// THEN the expression is parsed
	// AND returned without error
	happyTest(t, b, exp)
}

func TestParseAll_10(t *testing.T) {

	// GIVEN an expression with two consecutive operators
	// 1 + + 2
	b := token.Block{
		token.Statement{
			lex(token.TK_NUMBER, "1"),
			lex(token.TK_ADD, "+"),
			lex(token.TK_ADD, "+"),
			lex(token.TK_NUMBER, "2"),
		},
	}

	// WHEN parsing the statement
	// THEN an error is returned
	unhappyTest(t, b)
}

func TestParseAll_11(t *testing.T) {

	// GIVEN an expression with two consecutive operands
	// 1 2
	b := token.Block{
		token.Statement{
			lex(token.TK_NUMBER, "1"),
			lex(token.TK_NUMBER, "2"),
		},
	}

	// WHEN parsing the statement
	// THEN an error is returned
	unhappyTest(t, b)
}

func TestParseAll_12(t *testing.T) {

	// GIVEN an expression that begins with an invalid operator
	// +
	b := token.Block{
		token.Statement{
			lex(token.TK_ADD, "+"),
		},
	}

	// WHEN parsing the statement
	// THEN an error is returned
	unhappyTest(t, b)
}

func TestParseAll_13(t *testing.T) {

	// GIVEN an expression that ends with an invalid operator
	// 1 +
	b := token.Block{
		token.Statement{
			lex(token.TK_NUMBER, "1"),
			lex(token.TK_ADD, "+"),
		},
	}

	// WHEN parsing the statement
	// THEN an error is returned
	unhappyTest(t, b)
}

func TestParseAll_14(t *testing.T) {

	// GIVEN an expression without closing parenthesis
	// (1
	b := token.Block{
		token.Statement{
			lex(token.TK_PAREN_OPEN, "("),
			lex(token.TK_NUMBER, "1"),
		},
	}

	// WHEN parsing the statement
	// THEN an error is returned
	unhappyTest(t, b)
}

func TestParseAll_15(t *testing.T) {

	// GIVEN an expression with a closing parenthesis but no open one
	// 1)
	b := token.Block{
		token.Statement{
			lex(token.TK_NUMBER, "1"),
			lex(token.TK_PAREN_CLOSE, ")"),
		},
	}

	// WHEN parsing the statement
	// THEN an error is returned
	unhappyTest(t, b)
}

func TestParseAll_16(t *testing.T) {

	// GIVEN an expression with empty parentheses
	// ()
	b := token.Block{
		token.Statement{
			lex(token.TK_PAREN_OPEN, "("),
			lex(token.TK_PAREN_CLOSE, ")"),
		},
	}

	// WHEN parsing the statement
	// THEN an error is returned
	unhappyTest(t, b)
}

func TestParseAll_17(t *testing.T) {

	// GIVEN a number followed by an opening parenthesis
	// 9 (1)
	b := token.Block{
		token.Statement{
			lex(token.TK_NUMBER, "9"),
			lex(token.TK_PAREN_OPEN, "("),
			lex(token.TK_NUMBER, "1"),
			lex(token.TK_PAREN_CLOSE, ")"),
		},
	}

	// WHEN parsing the statement
	// THEN an error is returned
	unhappyTest(t, b)
}

func TestParseAll_18(t *testing.T) {

	// GIVEN multiple statements
	// 1
	// 2
	// 3
	b := token.Block{
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

	exp := ast.Block{
		num(1),
		num(2),
		num(3),
	}

	// WHEN parsing all statements
	// THEN the number is parsed and returned without error
	happyTest(t, b, exp)
}

func TestParseAll_19(t *testing.T) {

	// GIVEN a number that will fail number parsing
	// abc9
	b := token.Block{
		token.Statement{
			lex(token.TK_NUMBER, "abc9"),
		},
	}

	// WHEN parsing the statement
	// THEN an error is returned
	unhappyTest(t, b)
}
