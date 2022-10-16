package debug

import (
	"fmt"
	"strings"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
)

func astNodes[T ast.Node](nodes []T) string {
	sb := &strings.Builder{}

	writeLine(sb, "[")

	for _, n := range nodes {
		writeIndentLine(sb, astNode(n), ",")
	}

	write(sb, "]")
	return sb.String()
}

func astNode(n ast.Node) string {
	switch t := n.(type) {
	case ast.Stmt:
		return astStmt(t)
		//case ast.WhenCase:
		//  return astWhenCase(t)
	}

	return "Unknown AST Node"
}

func astStmt(n ast.Stmt) string {
	switch t := n.(type) {
	case ast.Proc:
		return astProc(t)
	case ast.If:
		return astIf(t)
	case ast.Assign:
		return astAssign(t)
	}

	return "Unknown AST Stmt"
}

func astProc(n ast.Stmt) string {
	switch t := n.(type) {
	case ast.Expr:
		return astExpr(t)
	case ast.ExprSet:
		return astExprSet(t)
	}

	return "Unknown AST Stmt"
}

func astExpr(n ast.Expr) string {
	switch t := n.(type) {
	case ast.Literal:
		return astLiteral(t)
	case ast.Variable:
		return astVariable(t)
	}

	return "Unknown AST Expr"
}

func astVariable(n ast.Variable) string {
	return fmt.Sprintf("Variable %v", n.Identifier)
}

func astLiteral(n ast.Literal) string {
	return fmt.Sprintf("Literal %v", n.Token)
}

func astAssign(n ast.Assign) string {
	sb := &strings.Builder{}

	writeLine(sb, "Assign:")
	writeIndentLine(sb, "Operator: ", n.Operator.String())
	writeIndentLine(sb, "Left: ", astNodes[ast.Variable](n.Left))
	writeIndent(sb, "Right: ", astNode(n.Right))

	return sb.String()
}

func astExprSet(n ast.ExprSet) string {
	return astNodes(n.Exprs)
}

func astIf(n ast.If) string {
	sb := &strings.Builder{}

	writeLine(sb, "If:")
	writeIndentLine(sb, "Keyword: ", n.Keyword.String())
	writeIndentLine(sb, "Condition: ", astNode(n.Condition))
	writeIndentLine(sb, "Body: ", astNodes[ast.Stmt](n.Body))
	writeIndent(sb, "End: ", n.End.String())

	return sb.String()
}

func astWhen(n ast.When) string {
	// TODO
	return fmt.Sprintf("When %q", n.Keyword.Value)
}

func astWhenCase(n ast.WhenCase) string {
	// TODO
	return fmt.Sprintf("WhenCase")
}
