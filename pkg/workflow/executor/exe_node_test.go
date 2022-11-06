package executor

import (
	"testing"

	"github.com/stretchr/testify/require"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
)

var emptyState = NewState()

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

	// x = 1
	given := ast.Assign{
		Dst: mockVariables("x"),
		Src: mockLiterals(float64(1)),
	}

	// x: 1
	exp := &exeState{
		variables: map[string]any{
			"x": float64(1),
		},
	}

	actState := NewState()
	exeNode(actState, given)

	require.Equal(t, exp, actState)
}

func Test_2_exeAssign(t *testing.T) {

	// x = "abc"
	given := ast.Assign{
		Dst: mockVariables("x"),
		Src: mockLiterals("abc"),
	}

	// x: "abc"
	exp := &exeState{
		variables: map[string]any{
			"x": "abc",
		},
	}

	actState := NewState()
	exeNode(actState, given)

	require.Equal(t, exp, actState)
}

func Test_3_exeAssign(t *testing.T) {

	// x = true
	given := ast.Assign{
		Dst: mockVariables("x"),
		Src: mockLiterals(true),
	}

	// x: true
	exp := &exeState{
		variables: map[string]any{
			"x": true,
		},
	}

	actState := NewState()
	exeNode(actState, given)

	require.Equal(t, exp, actState)
}

func Test_4_exeAssign(t *testing.T) {

	// x, y, z = 1, "abc", true
	given := ast.Assign{
		Dst: mockVariables("x", "y", "z"),
		Src: mockLiterals(float64(1), "abc", true),
	}

	// x: 1
	// y: "abc"
	// z: true
	exp := &exeState{
		variables: map[string]any{
			"x": float64(1),
			"y": "abc",
			"z": true,
		},
	}

	actState := NewState()
	exeNode(actState, given)

	require.Equal(t, exp, actState)
}

func Test_5_exeIf(t *testing.T) {

	// if true
	// end
	given := ast.If{
		Condition: ast.Literal{Value: true},
		Body:      nil,
	}

	exp := emptyState

	actState := NewState()
	exeNode(actState, given)

	require.Equal(t, exp, actState)
}

func Test_6_exeIf(t *testing.T) {

	// if true
	//   x = true
	// end
	given := ast.If{
		Condition: ast.Literal{Value: true},
		Body: []ast.Stmt{
			ast.Assign{
				Dst: mockVariables("x"),
				Src: mockLiterals(true),
			},
		},
	}

	// x: true
	exp := &exeState{
		variables: map[string]any{
			"x": true,
		},
	}

	actState := NewState()
	exeNode(actState, given)

	require.Equal(t, exp, actState)
}

func Test_7_exeIf(t *testing.T) {

	// if false
	//   x = true
	// end
	given := ast.If{
		Condition: ast.Literal{Value: false},
		Body: []ast.Stmt{
			ast.Assign{
				Dst: mockVariables("x"),
				Src: mockLiterals(true),
			},
		},
	}

	exp := emptyState

	actState := NewState()
	exeNode(actState, given)

	require.Equal(t, exp, actState)
}

func Test_8_exeIf(t *testing.T) {

	// if 1 == 1
	//   x = true
	// end
	given := ast.If{
		Condition: ast.BinaryOperation{
			Left:     ast.Literal{Value: 1},
			Operator: "==",
			Right:    ast.Literal{Value: 1},
		},
		Body: []ast.Stmt{
			ast.Assign{
				Dst: mockVariables("x"),
				Src: mockLiterals(true),
			},
		},
	}

	// x: true
	exp := &exeState{
		variables: map[string]any{
			"x": true,
		},
	}

	actState := NewState()
	exeNode(actState, given)

	require.Equal(t, exp, actState)
}

func Test_9_exeIf(t *testing.T) {

	// if 1 == 2
	//   x = true
	// end
	given := ast.If{
		Condition: ast.BinaryOperation{
			Left:     ast.Literal{Value: 1},
			Operator: "==",
			Right:    ast.Literal{Value: 2},
		},
		Body: []ast.Stmt{
			ast.Assign{
				Dst: mockVariables("x"),
				Src: mockLiterals(true),
			},
		},
	}

	exp := emptyState

	actState := NewState()
	exeNode(actState, given)

	require.Equal(t, exp, actState)
}

func Test_10_exeIf(t *testing.T) {

	// if 1 != 2
	//   x = true
	// end
	given := ast.If{
		Condition: ast.BinaryOperation{
			Left:     ast.Literal{Value: 1},
			Operator: "!=",
			Right:    ast.Literal{Value: 2},
		},
		Body: []ast.Stmt{
			ast.Assign{
				Dst: mockVariables("x"),
				Src: mockLiterals(true),
			},
		},
	}

	// x: true
	exp := &exeState{
		variables: map[string]any{
			"x": true,
		},
	}

	actState := NewState()
	exeNode(actState, given)

	require.Equal(t, exp, actState)
}

func Test_11_exeIf(t *testing.T) {

	// if 1 != 1
	//   x = true
	// end
	given := ast.If{
		Condition: ast.BinaryOperation{
			Left:     ast.Literal{Value: 1},
			Operator: "!=",
			Right:    ast.Literal{Value: 1},
		},
		Body: []ast.Stmt{
			ast.Assign{
				Dst: mockVariables("x"),
				Src: mockLiterals(true),
			},
		},
	}

	exp := emptyState

	actState := NewState()
	exeNode(actState, given)

	require.Equal(t, exp, actState)
}
