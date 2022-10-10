package workflow

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/ast"
	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/token"

	"github.com/PaulioRandall/firefly-go/workflow/ast/asttest"
	"github.com/PaulioRandall/firefly-go/workflow/token/tokentest"
)

func Test_1_Workflow(t *testing.T) {
	r := inout.NewListReader([]rune(""))
	w := inout.NewListWriter[ast.Node]()

	e := Parse(r, w)

	require.Nil(t, e, "%+v", e)
	require.Empty(t, w.List())
}

func Test_2_Workflow(t *testing.T) {
	// a = 0

	r := inout.NewListReader([]rune("a = 0\n"))
	w := inout.NewListWriter[ast.Node]()

	var (
		gen    = tokentest.NewTokenGenerator()
		varTk  = gen(token.Var, "a")
		_      = gen(token.Space, " ")
		assSym = gen(token.Assign, "=")
		_      = gen(token.Space, " ")
		exprTk = gen(token.Number, "0")
	)

	exp := []ast.Node{
		ast.MakeAssign(
			assSym,
			asttest.Vars(varTk),
			asttest.LitExprs(exprTk),
		),
	}

	e := Parse(r, w)

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, w.List())
}

func Test_3_Workflow(t *testing.T) {
	// a, b, c = 0, 1, 2

	r := inout.NewListReader([]rune(`a, b, c = 0, 1, 2` + "\n"))
	w := inout.NewListWriter[ast.Node]()

	var (
		gen     = tokentest.NewTokenGenerator()
		varTk1  = gen(token.Var, "a")
		_       = gen(token.Comma, ",")
		_       = gen(token.Space, " ")
		varTk2  = gen(token.Var, "b")
		_       = gen(token.Comma, ",")
		_       = gen(token.Space, " ")
		varTk3  = gen(token.Var, "c")
		_       = gen(token.Space, " ")
		assSym  = gen(token.Assign, "=")
		_       = gen(token.Space, " ")
		exprTk1 = gen(token.Number, "0")
		_       = gen(token.Comma, ",")
		_       = gen(token.Space, " ")
		exprTk2 = gen(token.Number, "1")
		_       = gen(token.Comma, ",")
		_       = gen(token.Space, " ")
		exprTk3 = gen(token.Number, "2")
	)

	exp := []ast.Node{
		ast.MakeAssign(
			assSym,
			asttest.Vars(varTk1, varTk2, varTk3),
			asttest.LitExprs(exprTk1, exprTk2, exprTk3),
		),
	}

	e := Parse(r, w)

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, w.List())
}
