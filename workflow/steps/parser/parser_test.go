package parser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/ast"
	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/token"

	"github.com/PaulioRandall/firefly-go/workflow/ast/asttest"
	"github.com/PaulioRandall/firefly-go/workflow/token/tokentest"
)

func parseTok(tt token.TokenType, v string) token.Token {
	return tokentest.Tok(tt, v)
}

func literal(tt token.TokenType, v string) ast.Node {
	return asttest.Literal(parseTok(tt, v))
}

func assert(t *testing.T, given []token.Token, exp []ast.Node) {
	r := inout.NewListReader(given)
	w := inout.NewListWriter[ast.Node]()

	e := Parse(r, w)

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, w.List())
}

func Test_1_Compile(t *testing.T) {
	given := []token.Token{
		parseTok(token.Number, "0"),
		parseTok(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		literal(token.Number, "0"),
	}

	assert(t, given, exp)
}
