package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func Test_parseList_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// [1]
	given := []token.Token{
		gen(token.BracketOpen, "["),  // 0
		gen(token.Number, "1"),       // 1
		gen(token.BracketClose, "]"), // 2
		gen(token.Terminator, "\n"),
	}

	values := exprs(
		given[1],
	)

	exp := listExpr(
		given[0],
		values,
		given[2],
	)

	doParseTest(t, given, exp)
}

func Test_parseList_2(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// [1, 2]
	given := []token.Token{
		gen(token.BracketOpen, "["),  // 0
		gen(token.Number, "1"),       // 1
		gen(token.Comma, ","),        // 2
		gen(token.Number, "2"),       // 3
		gen(token.BracketClose, "]"), // 4
		gen(token.Terminator, "\n"),
	}

	values := exprs(
		given[1],
		given[3],
	)

	exp := listExpr(
		given[0],
		values,
		given[4],
	)

	doParseTest(t, given, exp)
}

func Test_parseList_3(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// [1, true, "abc", x]
	given := []token.Token{
		gen(token.BracketOpen, "["),  // 0
		gen(token.Number, "1"),       // 1
		gen(token.Comma, ","),        // 2
		gen(token.True, "true"),      // 3
		gen(token.Comma, ","),        // 4
		gen(token.String, `"abc"`),   // 5
		gen(token.Comma, ","),        // 6
		gen(token.Identifier, "x"),   // 7
		gen(token.BracketClose, "]"), // 8
		gen(token.Terminator, "\n"),
	}

	values := exprs(
		given[1],
		given[3],
		given[5],
		given[7],
	)

	exp := listExpr(
		given[0],
		values,
		given[8],
	)

	doParseTest(t, given, exp)
}

func Test_parseList_4(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// [1, true,]
	given := []token.Token{
		gen(token.BracketOpen, "["),  // 0
		gen(token.Number, "1"),       // 1
		gen(token.Comma, ","),        // 2
		gen(token.True, "true"),      // 3
		gen(token.Comma, ","),        // 4
		gen(token.BracketClose, "]"), // 5
		gen(token.Terminator, "\n"),
	}

	values := exprs(
		given[1],
		given[3],
	)

	exp := listExpr(
		given[0],
		values,
		given[5],
	)

	doParseTest(t, given, exp)
}

func Test_parseList_5(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// [1
	given := []token.Token{
		gen(token.BracketOpen, "["), // 0
		gen(token.Number, "1"),      // 1
		gen(token.Terminator, "\n"),
	}

	doErrorTest(t, given, UnexpectedToken)
}

func Test_parseList_6(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// [1, true
	given := []token.Token{
		gen(token.BracketOpen, "["), // 0
		gen(token.Number, "1"),      // 1
		gen(token.Comma, ","),       // 2
		gen(token.True, "true"),     // 3
		gen(token.Terminator, "\n"),
	}

	doErrorTest(t, given, UnexpectedToken)
}

func Test_parseList_7(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// [1,,]
	given := []token.Token{
		gen(token.BracketOpen, "["),  // 0
		gen(token.Number, "1"),       // 1
		gen(token.Comma, ","),        // 2
		gen(token.Comma, ","),        // 3
		gen(token.BracketClose, "]"), // 4
		gen(token.Terminator, "\n"),
	}

	doErrorTest(t, given, UnexpectedToken)
}

func Test_parseList_8(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// [,]
	given := []token.Token{
		gen(token.BracketOpen, "["),  // 0
		gen(token.Comma, ","),        // 1
		gen(token.BracketClose, "]"), // 2
		gen(token.Terminator, "\n"),
	}

	doErrorTest(t, given, UnexpectedToken)
}
