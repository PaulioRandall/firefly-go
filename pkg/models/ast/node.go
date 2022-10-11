package ast

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// Node represents an AST, or put differently, a executable statement
//
// Nodes maybe nested and dependent upon others in an acyclic manner such that
// at runtime the dependees are executed first and their results used in the
// node's execution
type Node interface {
	node()
}

// Stmt is a constraint for an executable statement
type Stmt interface {
	Node
	stmt()
}

// Proc (Procedure) is a constraint for a resolvable expression that can return
// any number of output values including none at all
type Proc interface {
	Stmt
	proc()
}

// Expr is specific Proc constraint that only and always returns a single value
//
// All sub nodes (recursive) of an Expr will also be an Expr
type Expr interface {
	Proc
	expr()
}

type baseNode struct{}

func (n baseNode) node() {}

type baseStmt struct{ baseNode }

func (n baseStmt) stmt() {}

type baseProc struct{ baseStmt }

func (n baseProc) proc() {}

type baseExpr struct{ baseProc }

func (n baseExpr) expr() {}

type Literal struct {
	baseExpr
	Token token.Token
}

func MakeLiteral(tk token.Token) Literal {
	return Literal{
		Token: tk,
	}
}

type Variable struct {
	baseExpr
	Token token.Token
}

func MakeVariable(tk token.Token) Variable {
	return Variable{
		Token: tk,
	}
}

type Assign struct {
	baseProc
	Left     []Variable
	Operator token.Token
	Right    []Expr
}

func MakeAssign(left []Variable, op token.Token, right []Expr) Assign {
	return Assign{
		Left:     left,
		Operator: op,
		Right:    right,
	}
}

type If struct {
	baseExpr
	Keyword   token.Token
	Condition Expr
	Body      []Stmt
	End       token.Token
}

func MakeIf(
	keyword token.Token,
	condition Expr,
	body []Stmt,
	end token.Token,
) If {
	return If{
		Keyword:   keyword,
		Condition: condition,
		Body:      body,
		End:       end,
	}
}
