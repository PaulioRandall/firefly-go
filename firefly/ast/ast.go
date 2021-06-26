// Package ast (Abstract Syntax Tree) defines the set of ASTs that are used to
// drive interpretation or compilation.
//
// Each tree represents a program statement with many also capable of being
// sub trees.
package ast

import (
	"strconv"
	"strings"
)

type (
	// Block represents a block of statements and may represent a whole program.
	Block []Tree

	// Tree is a common interface for all ASTs.
	Tree interface {

		// Type returns the tree type.
		Type() Node

		// String returns a human readable representation of the tree.
		String() string

		// Debug returns a string suitable for debugging.
		Debug() string
	}

	// EmptyTree represents an empty statement. This should only be used on its
	// own, never as a sub tree.
	EmptyTree struct {
	}

	// NumberTree represents an 64 bit integer value.
	NumberTree struct {
		Value int64
	}

	// InfixTree represents an infix operation such as addition or multipication.
	InfixTree struct {
		Node
		Left  Tree
		Right Tree
	}
)

// enforceTypes should be ignored and never used. It only serves to generate
// compiler errors if an tree struct does not fully implement the required
// interfaces within this package.
func enforceTypes() {
	var tr Tree

	tr = EmptyTree{}
	tr = NumberTree{}
	tr = InfixTree{}

	_ = tr
}

func (tr EmptyTree) Type() Node  { return NODE_EMPTY }
func (tr NumberTree) Type() Node { return NODE_NUM }
func (tr InfixTree) Type() Node  { return tr.Node }

func (tr EmptyTree) String() string  { return "" }
func (tr NumberTree) String() string { return strconv.FormatInt(tr.Value, 10) }
func (tr InfixTree) String() string {
	op := tr.Node.String()
	left := tr.Left.String()
	right := tr.Right.String()
	return "(" + left + " " + op + " " + right + ")"
}

func (tr EmptyTree) Debug() string {
	return tr.Type().String()
}

func (tr NumberTree) Debug() string {
	return tr.Type().String() + " " + tr.String()
}

func (tr InfixTree) Debug() string {
	sb := &strings.Builder{}

	writeLine(sb, 0, tr.Node.String())

	writeText(sb, 1, "L: ")
	writeTree(sb, 0, tr.Left)

	writeText(sb, 1, "R: ")
	writeTree(sb, 0, tr.Right)

	return sb.String()
}

func writeTree(sb *strings.Builder, indent int, tr Tree) {
	writeLine(sb, 0, tr.Type().String())
	indent++
	sb.WriteString(tr.String())
}

func writeLine(sb *strings.Builder, indent int, text string) {
	writeText(sb, indent, text)
	sb.WriteRune('\n')
}

func writeText(sb *strings.Builder, indent int, text string) {
	writeIndent(sb, indent)
	sb.WriteString(text)
}

func writeIndent(sb *strings.Builder, n int) {
	for i := 0; i < n; i++ {
		sb.WriteRune(' ')
		sb.WriteRune(' ')
	}
}
