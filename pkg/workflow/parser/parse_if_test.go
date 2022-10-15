package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast/asttest"
)

func Test_parseIf_1(t *testing.T) {
	// if true
	// end

	given := []token.Token{
		tok(token.If, "if"),
		tok(token.True, "true"),
		tok(token.Terminator, "\n"),
		tok(token.End, "end"),
		tok(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		asttest.If(
			given[0],
			asttest.Literal(given[1]),
			nil,
			given[3],
		),
	}

	assert(t, given, exp)
}

func Test_parseIf_2(t *testing.T) {
	// if true
	//   a = 0
	// end

	given := []token.Token{
		tok(token.If, "if"),
		tok(token.True, "true"),
		tok(token.Terminator, "\n"),
		tok(token.Identifier, "a"),
		tok(token.Assign, "="), // 4
		tok(token.Number, "0"),
		tok(token.Terminator, "\n"),
		tok(token.End, "end"),
		tok(token.Terminator, "\n"), // 8
	}

	body := []ast.Stmt{
		asttest.Assign(
			asttest.Variables(given[3]),
			given[4],
			asttest.Expressions(given[5]),
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

	assert(t, given, exp)
}

func Test_parseIf_3(t *testing.T) {
	// if true
	//   a = 0

	given := []token.Token{
		tok(token.If, "if"),
		tok(token.True, "true"),
		tok(token.Terminator, "\n"),
		tok(token.Identifier, "a"),
		tok(token.Assign, "="),
		tok(token.Number, "0"),
		tok(token.Terminator, "\n"),
	}

	assertError(t, given, UnexpectedEOF)
}
