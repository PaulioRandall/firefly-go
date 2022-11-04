package executor

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/workflow/executor/ast"
)

func mockVariables(names ...string) []ast.Variable {
	n := make([]ast.Variable, len(names))

	for i, v := range names {
		n[i] = ast.Variable{
			Name: v,
		}
	}

	return n
}

func mockLiterals(values ...any) []ast.Expr {
	n := make([]ast.Expr, len(values))

	for i, v := range values {
		n[i] = ast.Literal{
			Value: v,
		}
	}

	return n
}

func Test_1_exeAssign(t *testing.T) {
	given := ast.Assign{
		Dst: mockVariables("x"),
		Src: mockLiterals(float64(1)),
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
		Dst: mockVariables("x"),
		Src: mockLiterals("abc"),
	}

	expState := &exeState{
		variables: map[string]any{
			"x": "abc",
		},
	}

	actState := NewState()
	exeNode(actState, given)

	require.Equal(t, expState, actState)
}

func Test_3_exeAssign(t *testing.T) {
	given := ast.Assign{
		Dst: mockVariables("x"),
		Src: mockLiterals(true),
	}

	expState := &exeState{
		variables: map[string]any{
			"x": true,
		},
	}

	actState := NewState()
	exeNode(actState, given)

	require.Equal(t, expState, actState)
}

func Test_4_exeAssign(t *testing.T) {
	given := ast.Assign{
		Dst: mockVariables("x", "y", "z"),
		Src: mockLiterals(float64(1), "abc", true),
	}

	expState := &exeState{
		variables: map[string]any{
			"x": float64(1),
			"y": "abc",
			"z": true,
		},
	}

	actState := NewState()
	exeNode(actState, given)

	require.Equal(t, expState, actState)
}
