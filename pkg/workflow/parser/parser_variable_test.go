package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func Test_Parse_variable_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// x
	given := []token.Token{
		gen(token.Identifier, "x"),
		gen(token.Newline, "\n"),
	}

	doParseTest(t, given, varExpr(given[0]))
}
