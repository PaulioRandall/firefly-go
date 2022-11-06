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

func doAssignmentTest(t *testing.T, given []token.Token, exp ast.Assign) {
	defer func() {
		if e := recover(); e != nil {
			require.Fail(t, debug.String(e))
		}
	}()

	r := inout.NewListReader(given)
	br := inout.NewBufReader[token.Token](r)
	act := parseAssignment(br)

	require.Equal(t, exp, act, debug.String(act))
}

func Test_parseAssignment_1(t *testing.T) {

	// x = 1
	gen := tokentest.NewTokenGenerator()
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

	doAssignmentTest(t, given, exp)
}

func Test_parseAssignment_2(t *testing.T) {

	// x = "s"
	gen := tokentest.NewTokenGenerator()
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

	doAssignmentTest(t, given, exp)
}

func Test_parseAssignment_3(t *testing.T) {

	// x = true
	gen := tokentest.NewTokenGenerator()
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

	doAssignmentTest(t, given, exp)
}

func Test_parseAssignment_4(t *testing.T) {

	// x, y, z = true, 1, "abc"
	gen := tokentest.NewTokenGenerator()
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

	doAssignmentTest(t, given, exp)
}
