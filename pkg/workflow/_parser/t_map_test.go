package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func Test_map_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {}
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.BraceClose, "}"), // 1
		gen(token.Newline, "\n"),
	}

	exp := mapExpr(
		given[0],
		nil,
		given[1],
	)

	doParseTest(t, given, exp)
}

func Test_map_2(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {"one": 1}
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.String, `"one"`), // 1
		gen(token.Colon, ":"),      // 2
		gen(token.Number, "1"),     // 3
		gen(token.BraceClose, "}"), // 4
		gen(token.Newline, "\n"),
	}

	entries := mapEntries(
		mapEntry(given[1], given[3]),
	)

	exp := mapExpr(
		given[0],
		entries,
		given[4],
	)

	doParseTest(t, given, exp)
}

func Test_map_3(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {"one": 1, "two": 2}
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.String, `"one"`), // 1
		gen(token.Colon, ":"),      // 2
		gen(token.Number, "1"),     // 3
		gen(token.Comma, ","),      // 4
		gen(token.String, `"two"`), // 5
		gen(token.Colon, ":"),      // 6
		gen(token.Number, "2"),     // 7
		gen(token.BraceClose, "}"), // 8
		gen(token.Newline, "\n"),
	}

	entries := mapEntries(
		mapEntry(given[1], given[3]),
		mapEntry(given[5], given[7]),
	)

	exp := mapExpr(
		given[0],
		entries,
		given[8],
	)

	doParseTest(t, given, exp)
}

func Test_map_4(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {1: "one", 2: "two"}
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.Number, "1"),     // 1
		gen(token.Colon, ":"),      // 2
		gen(token.String, `"one"`), // 3
		gen(token.Comma, ","),      // 4
		gen(token.Number, "2"),     // 5
		gen(token.Colon, ":"),      // 6
		gen(token.String, `"two"`), // 7
		gen(token.BraceClose, "}"), // 8
		gen(token.Newline, "\n"),
	}

	entries := mapEntries(
		mapEntry(given[1], given[3]),
		mapEntry(given[5], given[7]),
	)

	exp := mapExpr(
		given[0],
		entries,
		given[8],
	)

	doParseTest(t, given, exp)
}

func Test_map_5(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {1: "one", 2: true, 3: 3, 4: x}
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.Number, "1"),     // 1
		gen(token.Colon, ":"),      // 2
		gen(token.String, `"one"`), // 3
		gen(token.Comma, ","),      // 4
		gen(token.Number, "2"),     // 5
		gen(token.Colon, ":"),      // 6
		gen(token.Bool, "true"),    // 7
		gen(token.Comma, ","),      // 8
		gen(token.Number, "3"),     // 9
		gen(token.Colon, ":"),      // 10
		gen(token.Number, "3"),     // 11
		gen(token.Comma, ","),      // 12
		gen(token.Number, "4"),     // 13
		gen(token.Colon, ":"),      // 14
		gen(token.Identifier, "a"), // 15
		gen(token.BraceClose, "}"), // 16
		gen(token.Newline, "\n"),
	}

	entries := mapEntries(
		mapEntry(given[1], given[3]),
		mapEntry(given[5], given[7]),
		mapEntry(given[9], given[11]),
		mapEntry(given[13], given[15]),
	)

	exp := mapExpr(
		given[0],
		entries,
		given[16],
	)

	doParseTest(t, given, exp)
}

func Test_map_6(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {"one": 1,}
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.String, `"one"`), // 1
		gen(token.Colon, ":"),      // 2
		gen(token.Number, "1"),     // 3
		gen(token.Comma, ","),      // 4
		gen(token.BraceClose, "}"), // 5
		gen(token.Newline, "\n"),
	}

	entries := mapEntries(
		mapEntry(given[1], given[3]),
	)

	exp := mapExpr(
		given[0],
		entries,
		given[5],
	)

	doParseTest(t, given, exp)
}

func Test_map_7(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {"one" 1}
	//
	// Missing colon in map entry
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.String, `"one"`), // 1
		gen(token.Number, "1"),     // 2
		gen(token.BraceClose, "}"), // 3
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrBadMapEntry,
		ErrBadMap,
		ErrBadExpr,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_map_8(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {"one": 1,,}
	//
	// Duplicate comma
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.String, `"one"`), // 1
		gen(token.Colon, ":"),      // 2
		gen(token.Number, "1"),     // 3
		gen(token.Comma, ","),      // 4
		gen(token.Comma, ","),      // 5
		gen(token.BraceClose, "}"), // 6
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrBadMapEntry,
		ErrBadMap,
		ErrBadExpr,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_map_9(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {: 1}
	//
	// Missing key
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.Colon, ":"),      // 2
		gen(token.Number, "1"),     // 3
		gen(token.BraceClose, "}"), // 4
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrBadMapEntry,
		ErrBadMap,
		ErrBadExpr,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_map_10(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {"one":}
	//
	// Missing value
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.String, `"one"`), // 1
		gen(token.Colon, ":"),      // 2
		gen(token.BraceClose, "}"), // 3
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrBadMapEntry,
		ErrBadMap,
		ErrBadExpr,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_map_11(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {"one": 1
	//
	// Missing closing brace
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.String, `"one"`), // 1
		gen(token.Colon, ":"),      // 2
		gen(token.Number, "1"),     // 3
		gen(token.Newline, "\n"),
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrMissingBraceClose,
		ErrBadMap,
		ErrBadExpr,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_map_12(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {,}
	//
	// Missing map entry
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.Comma, ","),      // 1
		gen(token.BraceClose, "}"), // 2
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrBadMapEntry,
		ErrBadMap,
		ErrBadExpr,
		ErrBadStmt,
		ErrParsing,
	)
}

func Test_map_13(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {"one" :: 1}
	//
	// Duplicate colon
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.String, `"one"`), // 1
		gen(token.Colon, ":"),      // 2
		gen(token.Colon, ":"),      // 3
		gen(token.Number, "1"),     // 4
		gen(token.BraceClose, "}"), // 5
		gen(token.Newline, "\n"),
	}

	doErrorTest(t, given,
		ErrUnexpectedToken,
		ErrBadMapEntry,
		ErrBadMap,
		ErrBadExpr,
		ErrBadStmt,
		ErrParsing,
	)
}
