package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast/asttest"
	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func Test_parseExpr_1(t *testing.T) {
	// 1

	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(token.Number, "0"),
		gen(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		asttest.Literal(given[0]),
	}

	assert(t, given, exp)
}

func Test_parseExpr_2(t *testing.T) {
	// "abc"

	gen := tokentest.NewTokenGenerator()
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
	// true

	gen := tokentest.NewTokenGenerator()
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
	// false

	gen := tokentest.NewTokenGenerator()
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
	// 1 + 1

	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(token.Number, "1"),
		gen(token.Add, "+"),
		gen(token.Number, "1"),
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
	// 1 + 1 + 1

	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(token.Number, "1"),
		gen(token.Add, "+"),
		gen(token.Number, "1"),
		gen(token.Add, "+"),
		gen(token.Number, "1"),
		gen(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		asttest.BinaryOperation(
			asttest.BinaryOperation(
				asttest.Literal(given[0]),
				given[1],
				asttest.Literal(given[2]),
			),
			given[3],
			asttest.Literal(given[4]),
		),
	}

	assert(t, given, exp)
}

func Test_parseExpr_7(t *testing.T) {
	// 1 + 1 * 1

	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(token.Number, "1"),
		gen(token.Add, "+"),
		gen(token.Number, "1"),
		gen(token.Mul, "*"),
		gen(token.Number, "1"),
		gen(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		asttest.BinaryOperation(
			asttest.Literal(given[0]),
			given[1],
			asttest.BinaryOperation(
				asttest.Literal(given[2]),
				given[3],
				asttest.Literal(given[4]),
			),
		),
	}

	assert(t, given, exp)
}

func Test_parseExpr_8(t *testing.T) {
	// 1 * 1 + 1

	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(token.Number, "1"),
		gen(token.Mul, "*"),
		gen(token.Number, "1"),
		gen(token.Add, "+"),
		gen(token.Number, "1"),
		gen(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		asttest.BinaryOperation(
			asttest.BinaryOperation(
				asttest.Literal(given[0]),
				given[1],
				asttest.Literal(given[2]),
			),
			given[3],
			asttest.Literal(given[4]),
		),
	}

	assert(t, given, exp)
}

func Test_parseExpr_9(t *testing.T) {
	// 1 * 1 + 1 * 1

	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(token.Number, "1"),
		gen(token.Mul, "*"),
		gen(token.Number, "1"),
		gen(token.Add, "+"),
		gen(token.Number, "1"),
		gen(token.Mul, "*"),
		gen(token.Number, "1"),
		gen(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		asttest.BinaryOperation(
			asttest.BinaryOperation(
				asttest.Literal(given[0]),
				given[1],
				asttest.Literal(given[2]),
			),
			given[3],
			asttest.BinaryOperation(
				asttest.Literal(given[4]),
				given[5],
				asttest.Literal(given[6]),
			),
		),
	}

	assert(t, given, exp)
}
