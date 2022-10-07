package parser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/ast"
	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/token"

	"github.com/PaulioRandall/firefly-go/workflow/ast/asttest"
	"github.com/PaulioRandall/firefly-go/workflow/token/tokentest"
)

func parseTok(tt token.TokenType, v string) token.Token {
	return tokentest.Tok(tt, v)
}

func literal(tt token.TokenType, v string) ast.Literal {
	return asttest.Literal(parseTok(tt, v))
}

func variable(tt token.TokenType, v string) ast.Variable {
	return asttest.Variable(parseTok(tt, v))
}

func assert(t *testing.T, given []token.Token, exp []ast.Node) {
	r := inout.NewListReader(given)
	w := inout.NewListWriter[ast.Node]()

	e := Parse(r, w)

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, w.List())
}

func Test_1(t *testing.T) {
	// a = 0

	given := []token.Token{
		parseTok(token.Var, "a"),
		parseTok(token.Assign, "="),
		parseTok(token.Number, "0"),
		parseTok(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		ast.Assign{
			Token: parseTok(token.Assign, "="),
			Left: []ast.Variable{
				variable(token.Var, "a"),
			},
			Right: []ast.Expr{
				literal(token.Number, "0"),
			},
		},
	}

	assert(t, given, exp)
}

func Test_2(t *testing.T) {
	// a, b = 0, 1

	given := []token.Token{
		parseTok(token.Var, "a"),
		parseTok(token.Comma, ","),
		parseTok(token.Var, "b"),
		parseTok(token.Assign, "="),
		parseTok(token.Number, "0"),
		parseTok(token.Comma, ","),
		parseTok(token.Number, "1"),
		parseTok(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		ast.Assign{
			Token: parseTok(token.Assign, "="),
			Left: []ast.Variable{
				variable(token.Var, "a"),
				variable(token.Var, "b"),
			},
			Right: []ast.Expr{
				literal(token.Number, "0"),
				literal(token.Number, "1"),
			},
		},
	}

	assert(t, given, exp)
}

// TODO: Test missing commas
// TODO: Test missing variable
// TODO: Test missing expression
