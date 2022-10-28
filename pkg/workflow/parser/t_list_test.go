package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func Test_list_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// [1]
	given := []token.Token{
		gen(token.BracketOpen, "["),  // 0
		gen(token.Number, "1"),       // 1
		gen(token.BracketClose, "]"), // 2
		gen(token.Newline, "\n"),
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

func Test_list_2(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// [1, 2]
	given := []token.Token{
		gen(token.BracketOpen, "["),  // 0
		gen(token.Number, "1"),       // 1
		gen(token.Comma, ","),        // 2
		gen(token.Number, "2"),       // 3
		gen(token.BracketClose, "]"), // 4
		gen(token.Newline, "\n"),
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

func Test_list_3(t *testing.T) {
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
		gen(token.Newline, "\n"),
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

func Test_list_4(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// [1, true,]
	given := []token.Token{
		gen(token.BracketOpen, "["),  // 0
		gen(token.Number, "1"),       // 1
		gen(token.Comma, ","),        // 2
		gen(token.True, "true"),      // 3
		gen(token.Comma, ","),        // 4
		gen(token.BracketClose, "]"), // 5
		gen(token.Newline, "\n"),
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

func Test_list_5(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// []
	given := []token.Token{
		gen(token.BracketOpen, "["),  // 0
		gen(token.BracketClose, "]"), // 1
		gen(token.Newline, "\n"),
	}

	exp := listExpr(
		given[0],
		nil,
		given[1],
	)

	doParseTest(t, given, exp)
}

func Test_list_6(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// [1
	given := []token.Token{
		gen(token.BracketOpen, "["), // 0
		gen(token.Number, "1"),      // 1
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrBadExpr,
		ErrMissingBracketClose,
		ErrBadList,
		ErrBadExpr,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_list_7(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// [1, true
	given := []token.Token{
		gen(token.BracketOpen, "["), // 0
		gen(token.Number, "1"),      // 1
		gen(token.Comma, ","),       // 2
		gen(token.True, "true"),     // 3
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrBadExpr,
		ErrMissingBracketClose,
		ErrBadList,
		ErrBadExpr,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_list_8(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// [1,,]
	given := []token.Token{
		gen(token.BracketOpen, "["),  // 0
		gen(token.Number, "1"),       // 1
		gen(token.Comma, ","),        // 2
		gen(token.Comma, ","),        // 3
		gen(token.BracketClose, "]"), // 4
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrBadExpr,
		ErrBadList,
		ErrBadExpr,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_list_9(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// [,]
	given := []token.Token{
		gen(token.BracketOpen, "["),  // 0
		gen(token.Comma, ","),        // 1
		gen(token.BracketClose, "]"), // 2
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrBadExpr,
		ErrBadList,
		ErrBadExpr,
		ErrBadStmt,
		ErrParsing,
	)
}
