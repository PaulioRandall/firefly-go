package parser2

import (
	"testing"

	"github.com/stretchr/testify/require"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/debug"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

type literalType interface {
	float64 | string | bool
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

func mockLiterals[T literalType](values ...T) []ast.Expr {
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

func Test_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// x = 1
	given := []token.Token{
		gen(token.Identifier, "x"),
		gen(token.Assign, "="),
		gen(token.Number, "1"),
		gen(token.Newline, "\n"),
	}

	exp := ast.Assign{
		Dst: mockVariables("x"),
		Src: mockLiterals(float64(1)),
	}

	doParseTest(t, given, exp)
}

func Test_2(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// x = "s"
	given := []token.Token{
		gen(token.Identifier, "x"),
		gen(token.Assign, "="),
		gen(token.String, `"s"`),
		gen(token.Newline, "\n"),
	}

	exp := ast.Assign{
		Dst: mockVariables("x"),
		Src: mockLiterals("s"),
	}

	doParseTest(t, given, exp)
}
