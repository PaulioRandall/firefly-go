package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func acceptExpressions(a tokenAuditor) []ast.Expr {
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

func acceptExpression(a tokenAuditor) ast.Expr {
	return acceptLiteral(a)
}

func acceptLiteral(a tokenAuditor) ast.Expr {
	if !acceptFunc(a, token.IsLiteral) {
		return nil
	}

	return ast.Literal{
		Token: a.Prev(),
	}
}

func expectExpressions(a tokenAuditor) []ast.Expr {
	var nodes []ast.Expr

	v := expectExpression(a)
	nodes = append(nodes, v)

	for accept(a, token.Comma) {
		v := expectExpression(a)
		nodes = append(nodes, v)
	}

	return nodes
}

func expectExpression(a tokenAuditor) ast.Expr {
	return expectLiteral(a)
}

func expectLiteral(a tokenAuditor) ast.Expr {
	return ast.Literal{
		Token: expectFunc(a, "literal", token.IsLiteral),
	}
}
