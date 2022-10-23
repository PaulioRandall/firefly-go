package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast/asttest"
	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func Test_parseExpr_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// 1
	given := []token.Token{
		gen(token.Number, "1"),
		gen(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		asttest.Literal(given[0]),
	}

	assert(t, given, exp)
}

func Test_parseExpr_2(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// "abc"
	given := []token.Token{
		gen(token.String, "abc"),
		gen(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		asttest.Literal(given[0]),
	}

	assert(t, given, exp)
}

func Test_parseExpr_3(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// true
	given := []token.Token{
		gen(token.True, "true"),
		gen(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		asttest.Literal(given[0]),
	}

	assert(t, given, exp)
}

func Test_parseExpr_4(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// false
	given := []token.Token{
		gen(token.False, "false"),
		gen(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		asttest.Literal(given[0]),
	}

	assert(t, given, exp)
}

func Test_parseExpr_5(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// 1 + 2
	given := []token.Token{
		gen(token.Number, "1"),
		gen(token.Add, "+"),
		gen(token.Number, "2"),
		gen(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		asttest.BinaryOperation(
			asttest.Literal(given[0]),
			given[1],
			asttest.Literal(given[2]),
		),
	}

	assert(t, given, exp)
}

func Test_parseExpr_6(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// 1 + 2 + 3
	// (1 + 2) + 3
	given := []token.Token{
		gen(token.Number, "1"),
		gen(token.Add, "+"),
		gen(token.Number, "2"),
		gen(token.Add, "+"),
		gen(token.Number, "3"),
		gen(token.Terminator, "\n"),
	}

	// 1 + 2
	a := asttest.BinaryOperation(
		asttest.Literal(given[0]),
		given[1],
		asttest.Literal(given[2]),
	)

	exp := []ast.Node{
		// a + 3
		asttest.BinaryOperation(a, given[3], asttest.Literal(given[4])),
	}

	assert(t, given, exp)
}

func Test_parseExpr_7(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// 1 + 2 * 3
	// 1 + (2 * 3)
	given := []token.Token{
		gen(token.Number, "1"),
		gen(token.Add, "+"),
		gen(token.Number, "2"),
		gen(token.Mul, "*"),
		gen(token.Number, "3"),
		gen(token.Terminator, "\n"),
	}

	// 2 * 3
	a := asttest.BinaryOperation(
		asttest.Literal(given[2]),
		given[3],
		asttest.Literal(given[4]),
	)

	// 1 + a
	exp := []ast.Node{
		asttest.BinaryOperation(asttest.Literal(given[0]), given[1], a),
	}

	assert(t, given, exp)
}

func Test_parseExpr_8(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// 1 * 2 + 3
	// (1 * 2) + 3
	given := []token.Token{
		gen(token.Number, "1"),
		gen(token.Mul, "*"),
		gen(token.Number, "2"),
		gen(token.Add, "+"),
		gen(token.Number, "3"),
		gen(token.Terminator, "\n"),
	}

	// 1 * 2
	a := asttest.BinaryOperation(
		asttest.Literal(given[0]),
		given[1],
		asttest.Literal(given[2]),
	)

	exp := []ast.Node{
		// a + 3
		asttest.BinaryOperation(a, given[3], asttest.Literal(given[4])),
	}

	assert(t, given, exp)
}

func Test_parseExpr_9(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// 1 * 2 + 3 * 4
	// (1 * 2) + (3 * 4)
	given := []token.Token{
		gen(token.Number, "1"),
		gen(token.Mul, "*"),
		gen(token.Number, "2"),
		gen(token.Add, "+"),
		gen(token.Number, "3"),
		gen(token.Mul, "*"),
		gen(token.Number, "4"),
		gen(token.Terminator, "\n"),
	}

	// 1 * 2
	a := asttest.BinaryOperation(
		asttest.Literal(given[0]),
		given[1],
		asttest.Literal(given[2]),
	)

	// 3 * 4
	b := asttest.BinaryOperation(
		asttest.Literal(given[4]),
		given[5],
		asttest.Literal(given[6]),
	)

	exp := []ast.Node{
		// a + b
		asttest.BinaryOperation(a, given[3], b),
	}

	assert(t, given, exp)
}
