package ast

import (
	"strconv"
	"strings"
)

type (
	Program []Node

	Node interface {
		Type() AST
		String() string
	}

	Number struct{ Value int64 }

	InfixNode struct {
		AST
		Left  Node
		Right Node
	}
)

func (n Number) Type() AST    { return AstNumber }
func (n InfixNode) Type() AST { return n.AST }

func (n Number) String() string {
	sb := &strings.Builder{}
	writeText(sb, 0, n.Type().String())
	writeLine(sb, 0, strconv.FormatInt(n.Value, 10))
	return sb.String()
}

func (n InfixNode) String() string {
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
