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

func numNode(n int64) ast.Node {
	return ast.Number{
		Value: n,
	}
}

func addNode(left, right ast.Node) ast.Node {
	return ast.Add{
		InfixOperation: ast.InfixOperation{
			Left:  left,
			Right: right,
		},
	}
}

func subNode(left, right ast.Node) ast.Node {
	return ast.Sub{
		InfixOperation: ast.InfixOperation{
			Left:  left,
			Right: right,
		},
	}
}

func mulNode(left, right ast.Node) ast.Node {
	return ast.Mul{
		InfixOperation: ast.InfixOperation{
			Left:  left,
			Right: right,
		},
	}
}

func divNode(left, right ast.Node) ast.Node {
	return ast.Div{
		InfixOperation: ast.InfixOperation{
			Left:  left,
			Right: right,
		},
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
		numNode(9),
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
		numNode(99),
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
		addNode(
			numNode(1),
			numNode(2),
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
		subNode(
			addNode(
				numNode(1),
				numNode(2),
			),
			numNode(3),
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
		addNode(
			numNode(1),
			mulNode(
				numNode(2),
				numNode(3),
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
		addNode(
			divNode(
				numNode(9),
				numNode(3),
			),
			mulNode(
				numNode(2),
				numNode(3),
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

	// 8 + [4 / 3] * 3 - 2 * 5
	n1 := divNode(
		numNode(4),
		numNode(3),
	)

	// 8 + [4 / 3 * 3] - 2 * 5
	n2 := mulNode(
		n1,
		numNode(3),
	)

	// [8 + 4 / 3 * 3] - 2 * 5
	n3 := addNode(
		numNode(8),
		n2,
	)

	// 8 + 4 / 3 * 3 - [2 * 5]
	n4 := mulNode(
		numNode(2),
		numNode(5),
	)

	// [8 + 4 / 3 * 3 - 2 * 5]
	exp := []ast.Node{
		subNode(n3, n4),
	}

	// WHEN parsing all statements
	act, e := ParseAll(lr)

	println(ast.String(act[0]))

	// THEN the expression is parsed
	// AND returned without error
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}
