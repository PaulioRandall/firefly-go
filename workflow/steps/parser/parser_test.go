package parser

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/ast"
	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/token"

	"github.com/PaulioRandall/firefly-go/workflow/token/tokentest"
)

func tok1(tt token.TokenType, v string) token.Token {
	return tokentest.Tok(tt, v)
}

func literal(tt token.TokenType, v string) ast.Literal {
	return ast.MakeLiteral(tok1(tt, v))
}

func variable(tt token.TokenType, v string) ast.Variable {
	return ast.MakeVariable(tok1(tt, v))
}

func assert(t *testing.T, given []token.Token, exp []ast.Node) {
	r := inout.NewListReader(given)
	w := inout.NewListWriter[ast.Node]()

	e := Parse(r, w)

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, w.List())
}

func assertError(t *testing.T, given []token.Token, exp error) {
	r := inout.NewListReader(given)
	w := inout.NewListWriter[ast.Node]()

	e := Parse(r, w)

	require.True(t, errors.Is(e, exp))
}

func Test_1(t *testing.T) {
	// a = 0

	given := []token.Token{
		tok1(token.Var, "a"),
		tok1(token.Assign, "="),
		tok1(token.Number, "0"),
		tok1(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		ast.MakeAssign(
			tok1(token.Assign, "="),
			[]ast.Variable{
				variable(token.Var, "a"),
			},
			[]ast.Expr{
				literal(token.Number, "0"),
			},
		),
	}

	assert(t, given, exp)
}

func Test_2(t *testing.T) {
	// a, b = 0, 1

	given := []token.Token{
		tok1(token.Var, "a"),
		tok1(token.Comma, ","),
		tok1(token.Var, "b"),
		tok1(token.Assign, "="),
		tok1(token.Number, "0"),
		tok1(token.Comma, ","),
		tok1(token.Number, "1"),
		tok1(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		ast.Assign{
			Token: tok1(token.Assign, "="),
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

func Test_3(t *testing.T) {
	// a b = 0, 1

	given := []token.Token{
		tok1(token.Var, "a"),
		tok1(token.Var, "b"),
		tok1(token.Assign, "="),
		tok1(token.Number, "0"),
		tok1(token.Comma, ","),
		tok1(token.Number, "1"),
		tok1(token.Terminator, "\n"),
	}

	assertError(t, given, UnexpectedToken)
}

// TODO: Test missing commas
// TODO: Test missing variable
// TODO: Test missing expression
