package executor

import (
	"testing"

	"github.com/stretchr/testify/require"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
)

func Test_invokeSpell_1(t *testing.T) {

	given := ast.SpellCall{
		Name:   "println",
		Params: nil,
	}

	state := newState()

	exp := []any(nil)

	act := invokeSpell(state, given)

	require.Equal(t, exp, act)
}
