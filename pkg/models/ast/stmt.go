package ast

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// Stmt represents a statement or constraint for an executable statement
type Stmt interface {
	Node
	stmt()
}

// Assign represents an assignment with left being the target variables and
// right being the statement that determines the new or updated variable values
type Assign struct {
	Left     SeriesOfVar
	Operator token.Token
	Right    Stmt
}

func (n Assign) node() {}
func (n Assign) stmt() {}

func (n Assign) Where() (from, to pos.Pos) {
	from, _ = n.Left.Where()
	_, to = n.Right.Where()
	return from, to
}

// If represents a conditional block of statements
type If struct {
	Keyword   token.Token
	Condition Expr
	Body      []Stmt
	End       token.Token
}

func (n If) node() {}
func (n If) stmt() {}

func (n If) Where() (from, to pos.Pos) {
	from, _ = n.Keyword.Where()
	_, to = n.End.Where()
	return from, to
}

// For represents a looped block of statements
type For struct {
	Keyword     token.Token
	Initialiser Stmt
	Condition   Expr
	Advancement Stmt
	Body        []Stmt
	End         token.Token
}

func (n For) node() {}
func (n For) stmt() {}

func (n For) Where() (from, to pos.Pos) {
	from, _ = n.Keyword.Where()
	_, to = n.End.Where()
	return from, to
}

// For represents a looped block of statements
type ForEach struct {
	Keyword token.Token
	Index   Variable
	Item    Variable
	List    Expr
	Body    []Stmt
	End     token.Token
}

func (n ForEach) node() {}
func (n ForEach) stmt() {}

func (n ForEach) Where() (from, to pos.Pos) {
	from, _ = n.Keyword.Where()
	_, to = n.End.Where()
	return from, to
}

// When represents a match block or branch with multiple options
type When struct {
	Keyword token.Token
	Subject Expr
	Cases   []WhenCase
	End     token.Token
}

func (n When) node() {}
func (n When) stmt() {}

func (n When) Where() (from, to pos.Pos) {
	from, _ = n.Keyword.Where()
	_, to = n.End.Where()
	return from, to
}

// WhenCase represents a matchable case within a When block
//
// Basically a sophisticated switch case without fallthrough
type WhenCase struct {
	Condition Expr
	Statement Stmt
}

func (n WhenCase) node() {}
func (n WhenCase) stmt() {}

func (n WhenCase) Where() (from, to pos.Pos) {
	from, _ = n.Condition.Where()
	_, to = n.Statement.Where()
	return from, to
}

// Watch represents a statement block that monitors a specific variable
// exiting the block if the variable changes.
type Watch struct {
	Keyword  token.Token
	Variable Variable
	Body     []Stmt
	End      token.Token
}

func (n Watch) node() {}
func (n Watch) stmt() {}

func (n Watch) Precedence() int {
	return 0
}

func (n Watch) Where() (from, to pos.Pos) {
	from, _ = n.Keyword.Where()
	_, to = n.End.Where()
	return from, to
}
