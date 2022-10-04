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

	act, e := Parse(r)

	require.Nil(t, e, "%+v", e)
	require.Empty(t, act)
}

func Test_2_Workflow(t *testing.T) {
	r := inout.NewListReader([]rune("0\n"))

	act, e := Parse(r)

	gen := tokentest.NewTokenGenerator()
	exp := []ast.Node{
		asttest.Literal(gen(token.Number, "0")),
	}

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}
