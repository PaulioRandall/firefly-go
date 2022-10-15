package debug

import (
	"fmt"
	"strings"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
)

var indent = strings.Repeat("  ", 1)

func Print(v any) {
	fmt.Print(String(v))
}

func Println(v any) {
	fmt.Println(String(v))
}

func String(v any) string {
	switch t := v.(type) {
	case error:
		return wrappedError(t)
	case ast.Node:
		return astNode(t)
	}

	return fmt.Sprintf("%+v", v)
}

func write(sb *strings.Builder, ss ...string) {
	for _, s := range ss {
		sb.WriteString(s)
	}
}

func writeLine(sb *strings.Builder, ss ...string) {
	write(sb, ss...)
	sb.WriteRune('\n')
}

func writeIndent(sb *strings.Builder, ss ...string) {
	sub := &strings.Builder{}
	write(sub, ss...)
	s := indentLines(sub.String(), 1, true)
	sb.WriteString(s)
}

func writeIndentLine(sb *strings.Builder, ss ...string) {
	writeIndent(sb, ss...)
	sb.WriteRune('\n')
}

func indentLines(s string, n int, includeFirst bool) string {
	indents := strings.Repeat("  ", n)
	s = strings.ReplaceAll(s, "\n", "\n"+indents)

	if includeFirst {
		return indents + s
	}
	return s
}
