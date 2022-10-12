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
		tok1(token.If, "if"),
		tok1(token.True, "true"),
		tok1(token.Terminator, "\n"),
		tok1(token.End, "end"),
		tok1(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		ast.MakeIf(
			given[0],
			ast.MakeLiteral(given[1]),
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
		tok1(token.If, "if"),
		tok1(token.True, "true"),
		tok1(token.Terminator, "\n"),
		tok1(token.Identifier, "a"),
		tok1(token.Assign, "="), // 4
		tok1(token.Number, "0"),
		tok1(token.Terminator, "\n"),
		tok1(token.End, "end"),
		tok1(token.Terminator, "\n"), // 8
	}

	body := []ast.Stmt{
		ast.MakeAssign(
			asttest.Vars(given[3]),
			given[4],
			asttest.LitExprs(given[5]),
		),
	}

	exp := []ast.Node{
		ast.MakeIf(
			given[0],
			ast.MakeLiteral(given[1]),
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
		tok1(token.If, "if"),
		tok1(token.True, "true"),
		tok1(token.Terminator, "\n"),
		tok1(token.Identifier, "a"),
		tok1(token.Assign, "="),
		tok1(token.Number, "0"),
		tok1(token.Terminator, "\n"),
	}

	assertError(t, given, UnexpectedEOF)
}
