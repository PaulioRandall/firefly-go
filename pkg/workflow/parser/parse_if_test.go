package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast/asttest"
	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func Test_parseIf_1(t *testing.T) {
	// if true
	// end

	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(token.If, "if"),
		gen(token.True, "true"),
		gen(token.Terminator, "\n"),
		gen(token.End, "end"),
		gen(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		asttest.If(
			given[0],
			asttest.Literal(given[1]),
			nil,
			given[3],
		),
	}

	doParseTest(t, given, exp)
}

func Test_parseIf_2(t *testing.T) {
	// if true
	//   a = 0
	// end

	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(token.If, "if"),
		gen(token.True, "true"),
		gen(token.Terminator, "\n"),
		gen(token.Identifier, "a"),
		gen(token.Assign, "="), // 4
		gen(token.Number, "0"),
		gen(token.Terminator, "\n"),
		gen(token.End, "end"),
		gen(token.Terminator, "\n"), // 8
	}

	body := []ast.Stmt{
		asttest.Assign(
			asttest.Variables(given[3]),
			given[4],
			asttest.ExprSet(asttest.Expressions(given[5])...),
		),
	}

	exp := []ast.Node{
		asttest.If(
			given[0],
			asttest.Literal(given[1]),
			body,
			given[7],
		),
	}

	doParseTest(t, given, exp)
}

func Test_parseIf_3(t *testing.T) {
	// if true
	//   a = 0

	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(token.If, "if"),
		gen(token.True, "true"),
		gen(token.Terminator, "\n"),
		gen(token.Identifier, "a"),
		gen(token.Assign, "="),
		gen(token.Number, "0"),
		gen(token.Terminator, "\n"),
	}

	doErrorTest(t, given, UnexpectedEOF)
}
