package parser

import (
	"testing"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
	//"github.com/PaulioRandall/firefly-go/pkg/utilities/auditor"
	//"github.com/PaulioRandall/firefly-go/pkg/models/ast/asttest"
)

func Test_parseWhen_1(t *testing.T) {
	// when
	// end

	given := []token.Token{
		tok(token.When, "when"),
		tok(token.Terminator, "\n"),
		tok(token.End, "end"),
		tok(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		ast.MakeWhen(
			given[0],
			nil,
			nil,
			given[2],
		),
	}

	assert(t, given, exp)
}

func Test_parseWhen_2(t *testing.T) {
	// when 1
	// end

	given := []token.Token{
		tok(token.When, "when"),
		tok(token.Number, "1"),
		tok(token.Terminator, "\n"),
		tok(token.End, "end"),
		tok(token.Terminator, "\n"),
	}

	exp := []ast.Node{
		ast.MakeWhen(
			given[0],
			ast.MakeLiteral(tok(token.Number, "1")),
			nil,
			given[3],
		),
	}

	assert(t, given, exp)
}

func Test_parseWhen_3(t *testing.T) {
	// when
	//   true:
	// end
}

func Test_parseWhen_4(t *testing.T) {
	// when
	//   1 == 1:
	// end
}

func Test_parseWhen_5(t *testing.T) {
	// when
	//   a == 1:
	//   a == 2:
	//   a == 3:
	// end
}

func Test_parseWhen_6(t *testing.T) {
	// when a
	//   is 1:
	// end
}

func Test_parseWhen_7(t *testing.T) {
	// when a
	//   is 1:
	//   is 2:
	//   is 3:
	// end
}

func Test_parseWhen_8(t *testing.T) {
	// when a
	//   a == 1:
	//   a == 2:
	//   a == 3:
	// end
}

func Test_parseWhen_9(t *testing.T) {
	// when a
	//   is 1:
	//   a == 2:
	//   is 3:
	//   a == 4:
	// end
}

func Test_parseWhen_10(t *testing.T) {
	// when a
	//   is 1:
	//   a == 2:
	//   true:
	// end
}

func Test_parseWhen_11(t *testing.T) {
	// when
	//   is 1
	// end

	// Error!
}

func Test_parseWhen_12(t *testing.T) {
	// when
	//   a == 1
	// end

	// Error!
}

func Test_parseWhen_13(t *testing.T) {
	// when
	//   a == 1

	// Error!
}

func Test_parseWhen_14(t *testing.T) {
	// when a
	//   is:
	// end

	// Error!
}
