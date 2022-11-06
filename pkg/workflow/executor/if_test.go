package executor

import (
	"testing"

	"github.com/stretchr/testify/require"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
)

func Test_exeIf_1(t *testing.T) {

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

func Test_exeIf_2(t *testing.T) {

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

func Test_exeIf_3(t *testing.T) {

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
