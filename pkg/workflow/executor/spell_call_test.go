package executor

import (
	"testing"

	"github.com/stretchr/testify/require"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/memory"
)

func Test_invokeSpell_1(t *testing.T) {

	given := ast.SpellCall{
		Name:   "meh",
		Params: nil,
	}

	mem := memory.NewMemory()

	exp := []any(nil)

	act := invokeSpell(mem, given)

	require.Equal(t, exp, act)
}

func Test_invokeSpell_2(t *testing.T) {

	given := ast.SpellCall{
		Name: "len",
		Params: []ast.Expr{
			mockString("abc"),
		},
	}

	mem := memory.NewMemory()

	exp := []any{
		float64(3),
	}

	act := invokeSpell(mem, given)

	require.Equal(t, exp, act)
}
