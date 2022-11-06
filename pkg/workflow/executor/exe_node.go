package executor

import (
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
)

var (
	ErrUnknownNode     = err.Trackable("Unknown Node")
	ErrUnknownStmtNode = err.Trackable("Unknown Stmt Node")
	ErrUnknownExprNode = err.Trackable("Unknown Expr Node")
)

func exeNodes(state *exeState, nodes []ast.Node) {
	for _, n := range nodes {
		exeNode(state, n)
	}
}

func exeNode(state *exeState, n ast.Node) {
	switch v := n.(type) {
	case ast.Stmt:
		exeStmt(state, v)
	default:
		panic(unknownNode(nil))
	}
}

func exeStmts(state *exeState, nodes []ast.Stmt) {
	for _, n := range nodes {
		exeNode(state, n)
	}
}

func exeStmt(state *exeState, n ast.Stmt) {
	switch v := n.(type) {
	case ast.Assign:
		exeAssign(state, v)
	case ast.If:
		exeIf(state, v)
	default:
		panic(unknownStmtNode())
	}
}

func exeExpr(state *exeState, n ast.Expr) any {
	switch v := n.(type) {
	case ast.Literal:
		return v.Value
	case ast.BinaryOperation:
		return exeBinaryOperation(state, v)
	default:
		panic(unknownExprNode())
	}
}

func exeAssign(state *exeState, n ast.Assign) {
	result := make([]any, len(n.Src))

	for i, v := range n.Src {
		result[i] = exeExpr(state, v)
	}

	for i, dst := range n.Dst {
		state.setVariable(dst.Name, result[i])
	}
}

func exeIf(state *exeState, n ast.If) {
	if exeExpr(state, n.Condition).(bool) {
		exeStmts(state, n.Body)
	}
}

func exeBinaryOperation(state *exeState, n ast.BinaryOperation) any {

	left := exeExpr(state, n.Left)
	right := exeExpr(state, n.Right)

	switch n.Operator {
	case "==":
		return left == right
	case "!=":
		return left != right
	}

	return nil
}

func unknownNode(e error) error {
	return ErrUnknownNode.Wrap(e, "Node type does not match any known executable type")
}

func unknownStmtNode() error {
	e := ErrUnknownStmtNode.Track("Stmt type does not match any known executable type")
	return unknownNode(e)
}

func unknownExprNode() error {
	e := ErrUnknownExprNode.Track("Expr type does not match any known executable type")
	return unknownNode(e)
}
