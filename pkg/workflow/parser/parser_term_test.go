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

func Test_Parse_literal_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// 1
	given := []token.Token{
		gen(token.Number, "1"),
		gen(token.Newline, "\n"),
	}

	doParseTest(t, given, lit(given[0]))
}

func Test_Parse_literal_2(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// "abc"
	given := []token.Token{
		gen(token.String, `"abc"`),
		gen(token.Newline, "\n"),
	}

	doParseTest(t, given, lit(given[0]))
}

func Test_Parse_literal_3(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// true
	given := []token.Token{
		gen(token.True, "true"),
		gen(token.Newline, "\n"),
	}

	doParseTest(t, given, lit(given[0]))
}

func Test_Parse_literal_4(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// false
	given := []token.Token{
		gen(token.False, "false"),
		gen(token.Newline, "\n"),
	}

	doParseTest(t, given, lit(given[0]))
}
