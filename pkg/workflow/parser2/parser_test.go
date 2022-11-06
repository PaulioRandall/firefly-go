package parser2

import (
	"testing"

	"github.com/stretchr/testify/require"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/debug"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

func mockBool(v bool) ast.Literal {
	return ast.Literal{Value: v}
}

func mockNumber(v float64) ast.Literal {
	return ast.Literal{Value: v}
}

func mockString(v string) ast.Literal {
	return ast.Literal{Value: v}
}

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

func doParseTest(t *testing.T, given []token.Token, exp ...ast.Node) {
	r := inout.NewListReader(given)
	w := inout.NewListWriter[ast.Node]()

	e := Parse(r, w)

	require.Nil(t, e, "%s", debug.String(e))
	require.Equal(t, exp, w.List(), debug.String(w.List()))
}
