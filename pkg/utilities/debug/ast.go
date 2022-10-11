package debug

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
)

func AstVariable(n ast.Variable) string {
	return fmt.Sprintf("Variable %q", n.Token.Value)
}

func AstLiteral(n ast.Literal) string {
	return fmt.Sprintf("Literal %q", n.Token.Value)
}

func AstAssign(n ast.Assign) string {
	return fmt.Sprintf("TODO: %q", n.Operator.Value)
}

func AstIf(n ast.If) string {
	return fmt.Sprintf("If %q", n.Keyword.Value)
}
