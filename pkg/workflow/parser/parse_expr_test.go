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
