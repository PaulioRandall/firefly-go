package ast

import (
	"strconv"
	"strings"
)

type (
	Node interface{ Type() AST }

	Number struct{ Value int64 }

	InfixOperation struct {
		Left  Node
		Right Node
	}

	Add struct{ InfixOperation }
	Sub struct{ InfixOperation }
	Mul struct{ InfixOperation }
	Div struct{ InfixOperation }
)

func (t Number) Type() AST { return AstNumber }
func (t Add) Type() AST    { return AstAdd }
func (t Sub) Type() AST    { return AstSub }
func (t Mul) Type() AST    { return AstMul }
func (t Div) Type() AST    { return AstDiv }

func String(n Node) string {
	sb := strings.Builder{}
	writeNode(&sb, 0, n)
	return sb.String()
}

func writeNode(sb *strings.Builder, indent int, n Node) {

	writeLine(sb, 0, n.Type().String())

	indent++

	if v, ok := n.(Number); ok {
		num := strconv.FormatInt(v.Value, 10)
		writeLine(sb, indent, num)
		return
	}

	switch v := n.(type) {
	case Add:
		writeInfixOperation(sb, indent, v.InfixOperation)
	case Sub:
		writeInfixOperation(sb, indent, v.InfixOperation)
	case Mul:
		writeInfixOperation(sb, indent, v.InfixOperation)
	case Div:
		writeInfixOperation(sb, indent, v.InfixOperation)
		return
	}

}

func writeInfixOperation(sb *strings.Builder, indent int, n InfixOperation) {
	writeText(sb, indent, "L: ")
	writeNode(sb, indent, n.Left)

	writeText(sb, indent, "R: ")
	writeNode(sb, indent, n.Right)
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
