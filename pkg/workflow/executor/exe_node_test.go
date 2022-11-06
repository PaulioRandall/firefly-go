package executor

import (
	"testing"

	"github.com/stretchr/testify/require"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
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

	// x = 1
	given := ast.Assign{
		Dst: mockVariables("x"),
		Src: mockLiterals(float64(1)),
	}

	exp := newState()
	exp.setVariable("x", float64(1))

	act := newState()
	exeAssign(act, given)

	require.Equal(t, exp, act)
}

func Test_2_exeAssign(t *testing.T) {

	// x = "abc"
	given := ast.Assign{
		Dst: mockVariables("x"),
		Src: mockLiterals("abc"),
	}

	exp := newState()
	exp.setVariable("x", "abc")

	act := newState()
	exeAssign(act, given)

	require.Equal(t, exp, act)
}

func Test_3_exeAssign(t *testing.T) {

	// x = true
	given := ast.Assign{
		Dst: mockVariables("x"),
		Src: mockLiterals(true),
	}

	exp := newState()
	exp.setVariable("x", true)

	act := newState()
	exeAssign(act, given)

	require.Equal(t, exp, act)
}

func Test_4_exeAssign(t *testing.T) {

	// x, y, z = 1, "abc", true
	given := ast.Assign{
		Dst: mockVariables("x", "y", "z"),
		Src: mockLiterals(float64(1), "abc", true),
	}

	exp := newState()
	exp.setVariable("x", float64(1))
	exp.setVariable("y", "abc")
	exp.setVariable("z", true)

	act := newState()
	exeAssign(act, given)

	require.Equal(t, exp, act)
}

func Test_5_exeIf(t *testing.T) {

	// if true
	// end
	given := ast.If{
		Condition: ast.Literal{Value: true},
		Body:      nil,
	}

	exp := newState()

	act := newState()
	exeIf(act, given)

	require.Equal(t, exp, act)
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

	exp := newState()
	exp.setVariable("x", true)

	act := newState()
	exeIf(act, given)

	require.Equal(t, exp, act)
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

	exp := newState()

	act := newState()
	exeIf(act, given)

	require.Equal(t, exp, act)
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

	exp := newState()
	exp.setVariable("x", true)

	act := newState()
	exeIf(act, given)

	require.Equal(t, exp, act)
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

	exp := newState()

	act := newState()
	exeIf(act, given)

	require.Equal(t, exp, act)
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

	exp := newState()
	exp.setVariable("x", true)

	act := newState()
	exeIf(act, given)

	require.Equal(t, exp, act)
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

	exp := newState()

	act := newState()
	exeIf(act, given)

	require.Equal(t, exp, act)
}
