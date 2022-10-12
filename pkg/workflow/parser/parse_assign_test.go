package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/auditor"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast/asttest"
)

func Test_parseAssign_1(t *testing.T) {
	// a = 0

	given := []token.Token{
		tok1(token.Identifier, "a"),
		tok1(token.Assign, "="),
		tok1(token.Number, "0"),
		tok1(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		ast.MakeAssign(
			asttest.Vars(given[0]),
			given[1],
			asttest.LitExprs(given[2]),
		),
	}

	assert(t, given, exp)
}

func Test_parseAssign_2(t *testing.T) {
	// a, b = 0, 1

	given := []token.Token{
		tok1(token.Identifier, "a"),
		tok1(token.Comma, ","),
		tok1(token.Identifier, "b"),
		tok1(token.Assign, "="),
		tok1(token.Number, "0"),
		tok1(token.Comma, ","),
		tok1(token.Number, "1"),
		tok1(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		ast.MakeAssign(
			asttest.Vars(given[0], given[2]),
			given[3],
			asttest.LitExprs(given[4], given[6]),
		),
	}

	assert(t, given, exp)
}

func Test_parseAssign_3(t *testing.T) {
	// a b = 0, 1

	given := []token.Token{
		tok1(token.Identifier, "a"),
		tok1(token.Identifier, "b"),
		tok1(token.Assign, "="),
		tok1(token.Number, "0"),
		tok1(token.Comma, ","),
		tok1(token.Number, "1"),
		tok1(token.Terminator, "\n"),
	}

	assertError(t, given, auditor.UnexpectedToken)
}

func Test_parseAssign_4(t *testing.T) {
	// a, b = 0 1

	given := []token.Token{
		tok1(token.Identifier, "a"),
		tok1(token.Comma, ","),
		tok1(token.Identifier, "b"),
		tok1(token.Assign, "="),
		tok1(token.Number, "0"),
		tok1(token.Number, "1"),
		tok1(token.Terminator, "\n"),
	}

	assertError(t, given, MissingExpr)
}

func Test_parseAssign_5(t *testing.T) {
	// a, b = 0

	given := []token.Token{
		tok1(token.Identifier, "a"),
		tok1(token.Comma, ","),
		tok1(token.Identifier, "b"),
		tok1(token.Assign, "="),
		tok1(token.Number, "0"),
		tok1(token.Terminator, "\n"),
	}

	assertError(t, given, MissingExpr)
}

func Test_parseAssign_6(t *testing.T) {
	// a = 0, 1

	given := []token.Token{
		tok1(token.Identifier, "a"),
		tok1(token.Assign, "="),
		tok1(token.Number, "0"),
		tok1(token.Comma, ","),
		tok1(token.Number, "1"),
		tok1(token.Terminator, "\n"),
	}

	assertError(t, given, MissingVar)
}

func Test_parseAssign_7(t *testing.T) {
	// a, b 0, 1

	given := []token.Token{
		tok1(token.Identifier, "a"),
		tok1(token.Comma, ","),
		tok1(token.Identifier, "b"),
		tok1(token.Number, "0"),
		tok1(token.Comma, ","),
		tok1(token.Number, "1"),
		tok1(token.Terminator, "\n"),
	}

	assertError(t, given, auditor.UnexpectedToken)
}

func Test_parseAssign_8(t *testing.T) {
	// a, b, c = false, 0, ""

	given := []token.Token{
		tok1(token.Identifier, "a"),
		tok1(token.Comma, ","),
		tok1(token.Identifier, "b"),
		tok1(token.Comma, ","),
		tok1(token.Identifier, "c"),
		tok1(token.Assign, "="), // 5
		tok1(token.False, "false"),
		tok1(token.Comma, ","),
		tok1(token.Number, "0"),
		tok1(token.Comma, ","),
		tok1(token.String, `""`), // 10
		tok1(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		ast.MakeAssign(
			asttest.Vars(given[0], given[2], given[4]),
			given[5],
			asttest.LitExprs(given[6], given[8], given[10]),
		),
	}

	assert(t, given, exp)
}
