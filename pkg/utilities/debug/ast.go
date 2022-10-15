package debug

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
)

func astVariable(n ast.Variable) string {
	return fmt.Sprintf("Variable %q", n.Identifier.Value)
}

func astLiteral(n ast.Literal) string {
	return fmt.Sprintf("Literal %q", n.Token.Value)
}

func astAssign(n ast.Assign) string {
	// TODO
	return fmt.Sprintf("Assign: %q", n.Operator.Value)
}

func astIf(n ast.If) string {
	// TODO
	return fmt.Sprintf("If %q", n.Keyword.Value)
}

func astWhen(n ast.When) string {
	// TODO
	return fmt.Sprintf("When %q", n.Keyword.Value)
}

func astWhenCase(n ast.WhenCase) string {
	// TODO
	return fmt.Sprintf("WhenCase")
}
