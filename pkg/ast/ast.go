// Package ast (Abstract Syntax Tree) defines the set of AST types that are
// used to drive interpretation. Each tree represents a program statement with
// some containing sub statements.
//
// This package also defines the nodes within an AST.
package ast

// AST is the type of an AST.
type AST int

const (
	AstUndefined AST = iota
	AstEmpty
	AstNumber
	AstAdd
	AstSub
	AstMul
	AstDiv
)

var astNames = map[AST]string{
	AstEmpty:  "EMPTY",
	AstNumber: "NUMBER",
	AstAdd:    "ADD",
	AstSub:    "SUBTRACT",
	AstMul:    "MULTIPLY",
	AstDiv:    "DIVIDE",
}

// String returns the string representation of the AST.
func (a AST) String() string {
	return astNames[a]
}
