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
		Debug() string
	}

	NumberNode struct{ Value int64 }

	InfixExprNode struct {
		AST
		Left  Node
		Right Node
	}
)

func (n NumberNode) Type() AST    { return AstNumber }
func (n InfixExprNode) Type() AST { return n.AST }

func (n NumberNode) String() string {
	return strconv.FormatInt(n.Value, 10)
}

func (n InfixExprNode) String() string {
	// TODO
	panic("Not yet implemented!")
}

func (n NumberNode) Debug() string {
	sb := &strings.Builder{}
	writeText(sb, 0, n.Type().String()+" "+n.String())
	return sb.String()
}

func (n InfixExprNode) Debug() string {
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
