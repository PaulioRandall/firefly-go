package workflow

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast/asttest"
	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
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
		varTk  = gen(token.Identifier, "a")
		_      = gen(token.Space, " ")
		assSym = gen(token.Assign, "=")
		_      = gen(token.Space, " ")
		exprTk = gen(token.Number, "0")
	)

	exp := []ast.Node{
		ast.MakeAssign(
			asttest.Vars(varTk),
			assSym,
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
		varTk1  = gen(token.Identifier, "a")
		_       = gen(token.Comma, ",")
		_       = gen(token.Space, " ")
		varTk2  = gen(token.Identifier, "b")
		_       = gen(token.Comma, ",")
		_       = gen(token.Space, " ")
		varTk3  = gen(token.Identifier, "c")
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
			asttest.Vars(varTk1, varTk2, varTk3),
			assSym,
			asttest.LitExprs(exprTk1, exprTk2, exprTk3),
		),
	}

	e := Parse(r, w)

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, w.List())
}
