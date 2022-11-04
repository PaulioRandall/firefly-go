package executor

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/workflow/executor/ast"
)

func Test_1_exeAssign(t *testing.T) {
	given := ast.Assign{
		Dst: []ast.Variable{
			ast.Variable{Name: "x"},
		},
		Src: []ast.Expr{
			ast.Literal{Value: float64(1)},
		},
	}

	expState := &exeState{
		variables: map[string]any{
			"x": float64(1),
		},
	}

	actState := NewState()
	exeNode(actState, given)

	require.Equal(t, expState, actState)
}

func Test_2_exeAssign(t *testing.T) {
	given := ast.Assign{
		Dst: []ast.Variable{
			ast.Variable{Name: "x"},
			ast.Variable{Name: "y"},
			ast.Variable{Name: "z"},
		},
		Src: []ast.Expr{
			ast.Literal{Value: float64(1)},
			ast.Literal{Value: float64(2)},
			ast.Literal{Value: float64(3)},
		},
	}

	expState := &exeState{
		variables: map[string]any{
			"x": float64(1),
			"y": float64(2),
			"z": float64(3),
		},
	}

	actState := NewState()
	exeNode(actState, given)

	require.Equal(t, expState, actState)
}
