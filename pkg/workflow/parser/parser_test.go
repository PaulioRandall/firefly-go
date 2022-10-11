package parser

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast/asttest"
	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func tok1(tt token.TokenType, v string) token.Token {
	return tokentest.Tok(tt, v)
}

func literal(tt token.TokenType, v string) ast.Literal {
	return ast.MakeLiteral(tok1(tt, v))
}

func variable(tt token.TokenType, v string) ast.Variable {
	return ast.MakeVariable(tok1(tt, v))
}

func assert(t *testing.T, given []token.Token, exp []ast.Node) {
	r := inout.NewListReader(given)
	w := inout.NewListWriter[ast.Node]()

	e := Parse(r, w)

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, w.List())
}

func assertError(t *testing.T, given []token.Token, exp error) {
	r := inout.NewListReader(given)
	w := inout.NewListWriter[ast.Node]()

	e := Parse(r, w)

	require.True(
		t,
		errors.Is(e, exp),
		"Want error %q but got error %q", exp, e,
	)
}

func Test_1(t *testing.T) {
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

func Test_2(t *testing.T) {
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

func Test_3(t *testing.T) {
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

	assertError(t, given, UnexpectedToken)
}

func Test_4(t *testing.T) {
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

func Test_5(t *testing.T) {
	// a, b = 0 1

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

func Test_6(t *testing.T) {
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

func Test_7(t *testing.T) {
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

	assertError(t, given, UnexpectedToken)
}

func Test_8(t *testing.T) {
	// a, b, c := false, 0, ""

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

func Test_9(t *testing.T) {
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

func Test_10(t *testing.T) {
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

func Test_11(t *testing.T) {
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

	assertError(t, given, MissingEnd)
}
