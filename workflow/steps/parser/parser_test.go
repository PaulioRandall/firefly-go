package parser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/tokenreader"

	"github.com/PaulioRandall/firefly-go/workflow/ast"
	"github.com/PaulioRandall/firefly-go/workflow/token"

	"github.com/PaulioRandall/firefly-go/workflow/ast/asttest"
	"github.com/PaulioRandall/firefly-go/workflow/token/tokentest"
)

func tok(tt token.TokenType, v string) token.Token {
	return tokentest.Tok(tt, v)
}

func literal(tt token.TokenType, v string) ast.Node {
	return asttest.Literal(tok(tt, v))
}

func assert(t *testing.T, given []token.Token, exp []ast.Node) {
	tr := tokenreader.FromList(given...)
	act, e := Parse(tr)

	require.Nil(t, e, "%+v", e)
	require.Equal(t, exp, act)
}

func Test_1_Compile(t *testing.T) {
	given := []token.Token{
		tok(token.Number, "0"),
		tok(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		literal(token.Number, "0"),
	}

	assert(t, given, exp)
}
