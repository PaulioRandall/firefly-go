package ast

import (
	"strconv"
	"strings"
)

type (
	// Block represents a block of statements and may represent a whole program.
	Block []Node

	// Node is a common interface for all AST nodes.
	Node interface {

		// Type returns the AST type.
		Type() AST

		// String returns a human readable representation of the node.
		String() string

		// Debug returns a string suitable for debugging.
		Debug() string
	}

	// EmptyNode represents an empty statement. This should only be used as a
	// root node, never as member of another node.
	EmptyNode struct {
	}

	// NumberNode represents an 64 btit integer value.
	NumberNode struct {
		Value int64
	}

	// InfixNode represents an infix operation such as addition or multipication.
	InfixNode struct {
		AST
		Left  Node
		Right Node
	}
)

// enforceTypes should be ignored and never used. It only serves to generate
// compiler errors if a node struct does not fully implement the required
// interfaces within this package.
func enforceTypes() {
	var n Node

	n = EmptyNode{}
	n = NumberNode{}
	n = InfixNode{}

	_ = n
}

func (n EmptyNode) Type() AST  { return AstEmpty }
func (n NumberNode) Type() AST { return AstNumber }
func (n InfixNode) Type() AST  { return n.AST }

func (n EmptyNode) String() string  { return "" }
func (n NumberNode) String() string { return strconv.FormatInt(n.Value, 10) }
func (n InfixNode) String() string {
	op := n.AST.String()
	left := n.Left.String()
	right := n.Right.String()
	return "(" + left + " " + op + " " + right + ")"
}

func (n EmptyNode) Debug() string {
	return n.Type().String()
}

func (n NumberNode) Debug() string {
	return n.Type().String() + " " + n.String()
}

func (n InfixNode) Debug() string {
	sb := &strings.Builder{}

	writeLine(sb, 0, n.AST.String())

	writeText(sb, 1, "L: ")
	writeNode(sb, 0, n.Left)

	writeText(sb, 1, "R: ")
	writeNode(sb, 0, n.Right)

	return sb.String()
}

func writeNode(sb *strings.Builder, indent int, n Node) {
	writeLine(sb, 0, n.Type().String())
	indent++
	sb.WriteString(n.String())
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
