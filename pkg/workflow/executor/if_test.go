package executor

import (
	"testing"

	"github.com/stretchr/testify/require"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/memory"
)

func Test_exeIf_1(t *testing.T) {

	// if true
	// end
	given := ast.If{
		Condition: ast.Literal{Value: true},
		Body:      nil,
	}

	exp := memory.NewMemory()

	act := memory.NewMemory()
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

	exp := memory.NewMemory()
	exp.Variables["x"] = true

	act := memory.NewMemory()
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

	exp := memory.NewMemory()

	act := memory.NewMemory()
	exeIf(act, given)

	require.Equal(t, exp, act)
}
