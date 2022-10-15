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
	_ = Stmt(When{})

	_ = Proc(Assign{})

	_ = Expr(Literal{})
	_ = Expr(Variable{})
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
