package ast

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// Term is an Expr that is not an operation.
type Term interface {
	Expr
	term()
}

// Variable represents a variable value referenced using an identifier
type Variable struct {
	Identifier token.Token
}

func (n Variable) node() {}
func (n Variable) stmt() {}
func (n Variable) proc() {}
func (n Variable) expr() {}
func (n Variable) term() {}

func (n Variable) Precedence() int {
	return n.Identifier.Precedence()
}

func (n Variable) Where() (from, to pos.Pos) {
	return n.Identifier.Where()
}

// Literal represents one of the following literal tokens:
// - True
// - False
// - Number
// - String
type Literal struct {
	Token token.Token
}

func (n Literal) node() {}
func (n Literal) stmt() {}
func (n Literal) proc() {}
func (n Literal) expr() {}
func (n Literal) term() {}

func (n Literal) Precedence() int {
	return n.Token.Precedence()
}

func (n Literal) Where() (from, to pos.Pos) {
	return n.Token.Where()
}

// List represents an array or ordered list of values.
type List struct {
	Opener token.Token
	Values []Expr
	Closer token.Token
}

func (n List) node() {}
func (n List) stmt() {}
func (n List) proc() {}
func (n List) expr() {}
func (n List) term() {}

func (n List) Precedence() int {
	return 0
}

func (n List) Where() (from, to pos.Pos) {
	from, _ = n.Opener.Where()
	_, to = n.Closer.Where()
	return from, to
}

// Map represents an associative array or unordered set of key value pairs.
type Map struct {
	Opener  token.Token
	Entries []MapEntry
	Closer  token.Token
}

func (n Map) node() {}
func (n Map) stmt() {}
func (n Map) proc() {}
func (n Map) expr() {}
func (n Map) term() {}

func (n Map) Precedence() int {
	return 0
}

func (n Map) Where() (from, to pos.Pos) {
	from, _ = n.Opener.Where()
	_, to = n.Closer.Where()
	return from, to
}

type MapEntry struct {
	Key   Expr
	Value Expr
}

func (n MapEntry) node() {}
func (n MapEntry) stmt() {}
func (n MapEntry) proc() {}
func (n MapEntry) expr() {}
func (n MapEntry) term() {}

func (n MapEntry) Precedence() int {
	return 0
}

func (n MapEntry) Where() (from, to pos.Pos) {
	from, _ = n.Key.Where()
	_, to = n.Value.Where()
	return from, to
}
