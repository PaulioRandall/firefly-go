package ast

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

type whereNode struct {
	from, to pos.Pos
}

func (n whereNode) node() {}
func (n whereNode) Where() (from, to pos.Pos) {
	return n.from, n.to
}

func assertWhere(t *testing.T, exp, act Node) {
	expFrom, expTo := exp.Where()
	actFrom, actTo := act.Where()
	require.Equal(t, expFrom, actFrom)
	require.Equal(t, expTo, actTo)
}

func Test_enforceTypes(t *testing.T) {
	_ = Stmt(If{})
	_ = Stmt(For{})
	_ = Stmt(When{})
	_ = Stmt(WhenCase{})
	_ = Stmt(Watch{})
	_ = Stmt(Assign{})

	_ = Proc(SeriesOfExpr{})

	_ = Expr(BinaryOperation{})
	_ = Expr(Is{})

	_ = Term(Literal{})
	_ = Term(Variable{})
	_ = Term(List{})
	_ = Term(Map{})
}

func Test_1_If(t *testing.T) {
	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(token.If, "if"),         // 2, 0, 2
		gen(token.True, "true"),     // 6, 0, 6
		gen(token.Terminator, "\n"), // 7, 1, 0
		gen(token.End, "end"),       // 10, 1, 3
		gen(token.Terminator, "\n"),
	}

	act := If{
		Keyword: given[0],
		Condition: Literal{
			Token: given[1],
		},
		Body: nil,
		End:  given[3],
	}

	exp := whereNode{
		from: pos.At(0, 0, 0),
		to:   pos.At(10, 1, 3),
	}

	assertWhere(t, exp, act)
}

func Test_2_When(t *testing.T) {
	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(token.When, "when"),     // 4, 0, 4
		gen(token.Terminator, "\n"), // 5, 1, 0
		gen(token.End, "end"),       // 8, 1, 3
		gen(token.Terminator, "\n"),
	}

	act := When{
		Keyword: given[0],
		Subject: nil,
		Cases:   nil,
		End:     given[2],
	}

	exp := whereNode{
		from: pos.At(0, 0, 0),
		to:   pos.At(8, 1, 3),
	}

	assertWhere(t, exp, act)
}

func Test_3_BinaryOperation(t *testing.T) {
	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(token.Number, "1"), // 1, 0, 1
		gen(token.Add, "+"),    // 2, 0, 2
		gen(token.Number, "1"), // 3, 0, 3
	}

	act := BinaryOperation{
		Left: Literal{
			Token: given[0],
		},
		Operator: given[1],
		Right: Literal{
			Token: given[2],
		},
	}

	exp := whereNode{
		from: pos.At(0, 0, 0),
		to:   pos.At(3, 0, 3),
	}

	assertWhere(t, exp, act)
}
