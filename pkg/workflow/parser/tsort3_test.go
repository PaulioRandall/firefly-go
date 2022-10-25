package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func Test_parseFor_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// for true
	// end
	given := []token.Token{
		gen(token.For, "for"),    // 0
		gen(token.True, "true"),  // 1
		gen(token.Newline, "\n"), // 2

		gen(token.End, "end"), // 3
		gen(token.Newline, "\n"),
	}

	exp := forStmt(
		given[0],
		nil,
		lit(given[1]),
		nil,
		nil,
		given[3],
	)

	doParseTest(t, given, exp)
}

func Test_parseFor_2(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// for i = 0; i < 5; i = i+1
	// end
	given := []token.Token{
		gen(token.For, "for"), // 0

		gen(token.Identifier, "i"), // 1
		gen(token.Assign, "="),     // 2
		gen(token.Number, "0"),     // 3
		gen(token.Terminator, ";"), // 4

		gen(token.Identifier, "i"), // 5
		gen(token.LT, "<"),         // 6
		gen(token.Number, "5"),     // 7
		gen(token.Terminator, ";"), // 8

		gen(token.Identifier, "i"), // 9
		gen(token.Assign, "="),     // 10
		gen(token.Identifier, "i"), // 11
		gen(token.Add, "+"),        // 12
		gen(token.Number, "1"),     // 13
		gen(token.Newline, "\n"),   // 14

		gen(token.End, "end"), // 15
		gen(token.Newline, "\n"),
	}

	initaliser := assStmt(
		vars(given[1]),
		given[2],
		lits(given[3]),
	)

	condition := binOp(
		varExpr(given[5]),
		given[6],
		lit(given[7]),
	)

	advancement := assStmt(
		vars(given[9]),
		given[10],
		[]ast.Expr{
			binOp(
				varExpr(given[11]),
				given[12],
				lit(given[13]),
			),
		},
	)

	exp := forStmt(
		given[0],
		initaliser,
		condition,
		advancement,
		nil,
		given[15],
	)

	doParseTest(t, given, exp)
}
