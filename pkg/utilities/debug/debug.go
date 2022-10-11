package debug

import (
	"fmt"
	"strings"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
)

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

	case ast.Literal:
		return astLiteral(t)
	case ast.Variable:
		return astVariable(t)
	case ast.Assign:
		return astAssign(t)
	case ast.If:
		return astIf(t)
	}

	return fmt.Sprintf("%+v", v)
}

func indentLines(s string, n int) string {
	indents := strings.Repeat("  ", n)
	return strings.ReplaceAll(s, "\n", "\n"+indents)
}
