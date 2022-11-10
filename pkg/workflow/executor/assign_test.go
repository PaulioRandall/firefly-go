package executor

import (
	"testing"

	"github.com/stretchr/testify/require"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
)

func Test_exeAssign_1(t *testing.T) {

	// x = 1
	given := ast.Assign{
		Dst: mockVariables("x"),
		Src: mockLiterals(float64(1)),
	}

	exp := NewMemory()
	exp.Variables["x"] = float64(1)

	act := NewMemory()
	exeAssign(act, given)

	require.Equal(t, exp, act)
}

func Test_exeAssign_2(t *testing.T) {

	// x = "abc"
	given := ast.Assign{
		Dst: mockVariables("x"),
		Src: mockLiterals("abc"),
	}

	exp := NewMemory()
	exp.Variables["x"] = "abc"

	act := NewMemory()
	exeAssign(act, given)

	require.Equal(t, exp, act)
}

func Test_exeAssign_3(t *testing.T) {

	// x = true
	given := ast.Assign{
		Dst: mockVariables("x"),
		Src: mockLiterals(true),
	}

	exp := NewMemory()
	exp.Variables["x"] = true

	act := NewMemory()
	exeAssign(act, given)

	require.Equal(t, exp, act)
}

func Test_exeAssign_4(t *testing.T) {

	// x, y, z = 1, "abc", true
	given := ast.Assign{
		Dst: mockVariables("x", "y", "z"),
		Src: mockLiterals(float64(1), "abc", true),
	}

	exp := NewMemory()
	exp.Variables["x"] = float64(1)
	exp.Variables["y"] = "abc"
	exp.Variables["z"] = true

	act := NewMemory()
	exeAssign(act, given)

	require.Equal(t, exp, act)
}
