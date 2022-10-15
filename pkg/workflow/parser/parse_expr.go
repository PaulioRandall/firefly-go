package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/auditor"
)

func acceptExpressions(a *auditor.Auditor) []ast.Expr {
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

func acceptExpression(a *auditor.Auditor) ast.Expr {
	return acceptLiteral(a)
}

func acceptLiteral(a *auditor.Auditor) ast.Expr {
	if !a.AcceptFunc(token.IsLiteral) {
		return nil
	}

	return ast.Literal{
		Token: a.Prev(),
	}
}

func expectExpressions(a *auditor.Auditor) []ast.Expr {
	var nodes []ast.Expr

	v := expectExpression(a)
	nodes = append(nodes, v)

	for a.Accept(token.Comma) {
		v := expectExpression(a)
		nodes = append(nodes, v)
	}

	return nodes
}

func expectExpression(a *auditor.Auditor) ast.Expr {
	return expectLiteral(a)
}

func expectLiteral(a *auditor.Auditor) ast.Expr {
	return ast.Literal{
		Token: a.ExpectFunc("literal", token.IsLiteral),
	}
}
