package firefly

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

func TestParseFile_1(t *testing.T) {

	// GIVEN an empty scroll
	file := "testdata/empty.scroll"

	// WHEN parsing the scroll
	act, e := ParseFile(file)

	exp := ast.Program{}

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

	exp := ast.Program{
		infix(ast.AstAdd, num(2), num(2)),
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

	exp := ast.Program{
		infix(ast.AstSub, num(9), num(1)),
		infix(ast.AstAdd,
			infix(ast.AstDiv, num(32), num(8)),
			infix(ast.AstMul,
				infix(ast.AstSub, num(0), num(4)),
				num(2),
			),
		),
	}

	// THEN no errors are returned
	// AND the resultant program contains the parsed ASTs in the order specified
	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}
