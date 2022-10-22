package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func acceptExpressions(a auditor) []ast.Expr {
	var nodes []ast.Expr

	for a.More() {
		v := acceptExpression(a)
		if v == nil {
			break
		}

		nodes = append(nodes, v)
	}

	return nodes
}

func acceptExpression(a auditor) ast.Expr {
	left := acceptLiteral(a)
	return operation(a, left)
}

func acceptLiteral(a auditor) ast.Expr {
	tk, ok := a.acquireIf(token.IsLiteral)
	if !ok {
		return nil
	}

	return ast.Literal{
		Token: tk,
	}
}

func expectExpressions(a auditor) []ast.Expr {
	var nodes []ast.Expr

	v := expectExpression(a)
	nodes = append(nodes, v)

	for a.accept(token.Comma) {
		v := expectExpression(a)
		nodes = append(nodes, v)
	}

	return nodes
}

func expectExpression(a auditor) ast.Expr {
	left := acceptLiteral(a)
	return operation(a, left)
}

func expectLiteral(a auditor) ast.Expr {
	return ast.Literal{
		Token: a.expectFor("literal", token.IsLiteral),
	}
}

func operation(a auditor, left ast.Expr) ast.Expr {
	_, ok := a.acquireIf(token.IsBinaryOperator)
	if !ok {
		return left
	}

	// 1
	// + 1
	/*
		if acceptOperator() {

		}
	*/
	// TODO: continue

	return nil
}
