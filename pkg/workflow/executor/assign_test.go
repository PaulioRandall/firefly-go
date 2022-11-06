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

	exp := newState()
	exp.setVariable("x", float64(1))

	act := newState()
	exeAssign(act, given)

	require.Equal(t, exp, act)
}

func Test_exeAssign_2(t *testing.T) {

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

func Test_exeAssign_3(t *testing.T) {

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

func Test_exeAssign_4(t *testing.T) {

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
