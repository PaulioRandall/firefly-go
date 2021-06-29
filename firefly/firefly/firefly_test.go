package firefly

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/firefly/ast"
	"github.com/PaulioRandall/firefly-go/firefly/token"
)

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

func TestParseFile_1(t *testing.T) {

	// GIVEN an empty scroll
	file := "testdata/empty.scroll"

	// WHEN parsing the scroll
	act, e := ParseFile(file)

	var exp ast.Block

	//  THEN no errors are returned
	// AND the resultant program will contain no ASTs
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func TestParseFile_2(t *testing.T) {

	// GIVEN a scroll with a simple expression
	file := "testdata/simple.scroll"

	// WHEN parsing the scroll
	act, e := ParseFile(file)

	exp := ast.Block{
		infix(ast.NODE_ADD, num(2), num(2)),
	}

	// THEN no errors are returned
	// AND the resultant program contains the parsed AST
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func TestParseFile_3(t *testing.T) {

	// GIVEN a scroll with a multiple lines of expressions
	file := "testdata/multiline.scroll"

	// WHEN parsing the scroll
	act, e := ParseFile(file)

	exp := ast.Block{
		infix(ast.NODE_SUB, num(9), num(1)),
		infix(ast.NODE_ADD,
			infix(ast.NODE_DIV, num(32), num(8)),
			infix(ast.NODE_MUL,
				infix(ast.NODE_SUB, num(0), num(4)),
				num(2),
			),
		),
	}

	// THEN no errors are returned
	// AND the resultant program contains the parsed ASTs in the order specified
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}
