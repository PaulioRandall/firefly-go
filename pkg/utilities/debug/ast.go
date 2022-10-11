package debug

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
)

func astVariable(n ast.Variable) string {
	return fmt.Sprintf("Variable %q", n.Operator.Value)
}

func astLiteral(n ast.Literal) string {
	return fmt.Sprintf("Literal %q", n.Operator.Value)
}

func astAssign(n ast.Assign) string {
	// TODO
	return fmt.Sprintf("Assign: %q", n.Operator.Value)
}

func astIf(n ast.If) string {
	// TODO
	return fmt.Sprintf("If %q", n.Keyword.Value)
}
