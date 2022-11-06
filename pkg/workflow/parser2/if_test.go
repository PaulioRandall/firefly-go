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

func doIfStatementTest(t *testing.T, given []token.Token, exp ast.If) {
	defer func() {
		if e := recover(); e != nil {
			require.Fail(t, debug.String(e))
		}
	}()

	r := inout.NewListReader(given)
	br := inout.NewBufReader[token.Token](r)
	act := parseIfStatement(br)

	require.Equal(t, exp, act, debug.String(act))
}

func Test_parseIfStatement_5(t *testing.T) {

	// if true
	// end
	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(token.If, "if"),
		gen(token.Bool, "true"),
		gen(token.Newline, "\n"),
		gen(token.End, "end"),
		gen(token.Newline, "\n"),
	}

	exp := ast.If{
		Condition: mockBool(true),
		Body:      nil,
	}

	doIfStatementTest(t, given, exp)
}

func Test_parseIfStatement_6(t *testing.T) {

	// if true
	//
	//
	// end
	gen := tokentest.NewTokenGenerator()
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
		Condition: mockBool(true),
		Body:      nil,
	}

	doIfStatementTest(t, given, exp)
}

func Test_parseIfStatement_7(t *testing.T) {

	// if true
	//   x = 1
	// end
	gen := tokentest.NewTokenGenerator()
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
		Condition: mockBool(true),
		Body: []ast.Stmt{
			ast.Assign{
				Dst: mockVariables("x"),
				Src: mockLiterals(float64(1)),
			},
		},
	}

	doIfStatementTest(t, given, exp)
}

func Test_parseIfStatement_8(t *testing.T) {

	// if true
	//   x = 1
	//   y = 2
	//   z = 3
	// end
	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(token.If, "if"),
		gen(token.Bool, "true"),
		gen(token.Newline, "\n"),

		gen(token.Ident, "x"),
		gen(token.Assign, "="),
		gen(token.Number, "1"),
		gen(token.Newline, "\n"),

		gen(token.Ident, "y"),
		gen(token.Assign, "="),
		gen(token.Number, "2"),
		gen(token.Newline, "\n"),

		gen(token.Ident, "z"),
		gen(token.Assign, "="),
		gen(token.Number, "3"),
		gen(token.Newline, "\n"),

		gen(token.End, "end"),
		gen(token.Newline, "\n"),
	}

	exp := ast.If{
		Condition: mockBool(true),
		Body: []ast.Stmt{
			ast.Assign{
				Dst: mockVariables("x"),
				Src: mockLiterals(float64(1)),
			},
			ast.Assign{
				Dst: mockVariables("y"),
				Src: mockLiterals(float64(2)),
			},
			ast.Assign{
				Dst: mockVariables("z"),
				Src: mockLiterals(float64(3)),
			},
		},
	}

	doIfStatementTest(t, given, exp)
}

func Test_parseIfStatement_9(t *testing.T) {

	// if 1 == 1
	// end
	gen := tokentest.NewTokenGenerator()
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
			Left:     mockNumber(1),
			Operator: "==",
			Right:    mockNumber(1),
		},
		Body: nil,
	}

	doIfStatementTest(t, given, exp)
}
