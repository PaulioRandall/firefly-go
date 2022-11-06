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

func Test_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// x = 1
	given := []token.Token{
		gen(token.Ident, "x"),
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
		gen(token.Ident, "x"),
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

func Test_3(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// x = true
	given := []token.Token{
		gen(token.Ident, "x"),
		gen(token.Assign, "="),
		gen(token.Bool, "true"),
		gen(token.Newline, "\n"),
	}

	exp := ast.Assign{
		Dst: mockVariables("x"),
		Src: mockLiterals(true),
	}

	doParseTest(t, given, exp)
}

func Test_4(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// x, y, z = true, 1, "abc"
	given := []token.Token{
		gen(token.Ident, "x"),
		gen(token.Comma, ","),
		gen(token.Ident, "y"),
		gen(token.Comma, ","),
		gen(token.Ident, "z"),
		gen(token.Assign, "="),
		gen(token.Bool, "true"),
		gen(token.Comma, ","),
		gen(token.Number, "1"),
		gen(token.Comma, ","),
		gen(token.String, `"abc"`),
		gen(token.Newline, "\n"),
	}

	exp := ast.Assign{
		Dst: mockVariables("x", "y", "z"),
		Src: mockLiterals(true, float64(1), "abc"),
	}

	doParseTest(t, given, exp)
}

func Test_5(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// if true
	// end
	given := []token.Token{
		gen(token.If, "if"),
		gen(token.Bool, "true"),
		gen(token.Newline, "\n"),
		gen(token.End, "end"),
		gen(token.Newline, "\n"),
	}

	exp := ast.If{
		Condition: ast.Literal{Value: true},
		Body:      nil,
	}

	doParseTest(t, given, exp)
}

func Test_6(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// if true
	//
	//
	// end
	given := []token.Token{
		gen(token.If, "if"),
		gen(token.Bool, "true"),
		gen(token.Newline, "\n"),
		gen(token.Newline, "\n"),
		gen(token.Newline, "\n"),
		gen(token.End, "end"),
		gen(token.Newline, "\n"),
	}

	exp := ast.If{
		Condition: ast.Literal{Value: true},
		Body:      nil,
	}

	doParseTest(t, given, exp)
}

func Test_7(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// if true
	//   x = 1
	// end
	given := []token.Token{
		gen(token.If, "if"),
		gen(token.Bool, "true"),
		gen(token.Newline, "\n"),
		gen(token.Ident, "x"),
		gen(token.Assign, "="),
		gen(token.Number, "1"),
		gen(token.Newline, "\n"),
		gen(token.End, "end"),
		gen(token.Newline, "\n"),
	}

	exp := ast.If{
		Condition: ast.Literal{Value: true},
		Body: []ast.Stmt{
			ast.Assign{
				Dst: mockVariables("x"),
				Src: mockLiterals(float64(1)),
			},
		},
	}

	doParseTest(t, given, exp)
}

func Test_8(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// if false
	//   x = 1
	// end
	given := []token.Token{
		gen(token.If, "if"),
		gen(token.Bool, "false"),
		gen(token.Newline, "\n"),
		gen(token.Ident, "x"),
		gen(token.Assign, "="),
		gen(token.Number, "1"),
		gen(token.Newline, "\n"),
		gen(token.End, "end"),
		gen(token.Newline, "\n"),
	}

	exp := ast.If{
		Condition: ast.Literal{Value: false},
		Body: []ast.Stmt{
			ast.Assign{
				Dst: mockVariables("x"),
				Src: mockLiterals(float64(1)),
			},
		},
	}

	doParseTest(t, given, exp)
}

func Test_9(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// if 1 == 1
	// end
	given := []token.Token{
		gen(token.If, "if"),
		gen(token.Number, "1"),
		gen(token.Equ, "=="),
		gen(token.Number, "1"),
		gen(token.Newline, "\n"),
		gen(token.End, "end"),
		gen(token.Newline, "\n"),
	}

	exp := ast.If{
		Condition: ast.BinaryOperation{
			Left:     ast.Literal{Value: float64(1)},
			Operator: "==",
			Right:    ast.Literal{Value: float64(1)},
		},
		Body: nil,
	}

	doParseTest(t, given, exp)
}

func Test_10(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// if 1 != 1
	// end
	given := []token.Token{
		gen(token.If, "if"),
		gen(token.Number, "1"),
		gen(token.Equ, "!="),
		gen(token.Number, "1"),
		gen(token.Newline, "\n"),
		gen(token.End, "end"),
		gen(token.Newline, "\n"),
	}

	exp := ast.If{
		Condition: ast.BinaryOperation{
			Left:     ast.Literal{Value: float64(1)},
			Operator: "!=",
			Right:    ast.Literal{Value: float64(1)},
		},
		Body: nil,
	}

	doParseTest(t, given, exp)
}
