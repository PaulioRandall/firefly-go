package compiler

import (
	"testing"

	//"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/workflow/ast"
	//"github.com/PaulioRandall/firefly-go/workflow/readers/tokenreader"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

func tok(tt token.TokenType, v string) token.Token {
	return token.MakeToken(tt, v, token.Range{})
}

func assertCompilation(t *testing.T, given []token.Token, exp []ast.Node) {

}

func Test_1_Compile(t *testing.T) {
	given := []token.Token{
		tok(token.Number, "0"),
	}

	exp := []ast.Node{
		ast.MakeLiteral(
			tok(token.Number, "0"),
		),
	}

	assertCompilation(t, given, exp)
}
