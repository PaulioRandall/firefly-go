package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func Test_parseMap_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {}
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.BraceClose, "}"), // 1
		gen(token.Terminator, "\n"),
	}

	exp := mapExpr(
		given[0],
		nil,
		given[1],
	)

	doParseTest(t, given, exp)
}

func Test_parseMap_2(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {"one": 1}
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.String, `"one"`), // 1
		gen(token.Colon, ":"),      // 2
		gen(token.Number, "1"),     // 3
		gen(token.BraceClose, "}"), // 4
		gen(token.Terminator, "\n"),
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

func Test_parseMap_3(t *testing.T) {
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
		gen(token.Terminator, "\n"),
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

func Test_parseMap_4(t *testing.T) {
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
		gen(token.Terminator, "\n"),
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

func Test_parseMap_5(t *testing.T) {
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
		gen(token.True, "true"),    // 7
		gen(token.Comma, ","),      // 8
		gen(token.Number, "3"),     // 9
		gen(token.Colon, ":"),      // 10
		gen(token.Number, "3"),     // 11
		gen(token.Comma, ","),      // 12
		gen(token.Number, "4"),     // 13
		gen(token.Colon, ":"),      // 14
		gen(token.Identifier, "a"), // 15
		gen(token.BraceClose, "}"), // 16
		gen(token.Terminator, "\n"),
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

func Test_parseMap_6(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {"one": 1,}
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.String, `"one"`), // 1
		gen(token.Colon, ":"),      // 2
		gen(token.Number, "1"),     // 3
		gen(token.Comma, ","),      // 4
		gen(token.BraceClose, "}"), // 5
		gen(token.Terminator, "\n"),
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

func Test_parseMap_7(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {"one" 1}
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.String, `"one"`), // 1
		gen(token.Number, "1"),     // 2
		gen(token.BraceClose, "}"), // 3
		gen(token.Terminator, "\n"),
	}

	doErrorTest(t, given, UnexpectedToken)
}

func Test_parseMap_8(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {"one": 1,,}
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.String, `"one"`), // 1
		gen(token.Colon, ":"),      // 2
		gen(token.Number, "1"),     // 3
		gen(token.Comma, ","),      // 4
		gen(token.Comma, ","),      // 5
		gen(token.BraceClose, "}"), // 6
		gen(token.Terminator, "\n"),
	}

	doErrorTest(t, given, UnexpectedToken)
}

func Test_parseMap_9(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {: 1}
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.Colon, ":"),      // 2
		gen(token.Number, "1"),     // 3
		gen(token.BraceClose, "}"), // 4
		gen(token.Terminator, "\n"),
	}

	doErrorTest(t, given, UnexpectedToken)
}

func Test_parseMap_10(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {"one":}
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.String, `"one"`), // 1
		gen(token.Colon, ":"),      // 2
		gen(token.BraceClose, "}"), // 3
		gen(token.Terminator, "\n"),
	}

	doErrorTest(t, given, UnexpectedToken)
}

func Test_parseMap_11(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {"one": 1
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.String, `"one"`), // 1
		gen(token.Colon, ":"),      // 2
		gen(token.Number, "1"),     // 3
		gen(token.Terminator, "\n"),
	}

	doErrorTest(t, given, UnexpectedToken)
}

func Test_parseMap_12(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {,}
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.Comma, ","),      // 1
		gen(token.BraceClose, "}"), // 2
		gen(token.Terminator, "\n"),
	}

	doErrorTest(t, given, UnexpectedToken)
}

func Test_parseMap_13(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// {"one" :: 1}
	given := []token.Token{
		gen(token.BraceOpen, "{"),  // 0
		gen(token.String, `"one"`), // 1
		gen(token.Colon, ":"),      // 2
		gen(token.Colon, ":"),      // 3
		gen(token.Number, "1"),     // 4
		gen(token.BraceClose, "}"), // 5
		gen(token.Terminator, "\n"),
	}

	doErrorTest(t, given, UnexpectedToken)
}
