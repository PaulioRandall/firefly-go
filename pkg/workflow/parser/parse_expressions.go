package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func parseExpressions(a *auditor) []ast.Expr {
	var nodes []ast.Expr

	v := parseExpression(a)
	nodes = append(nodes, v)

	for a.accept(token.Comma) {
		v := parseExpression(a)
		nodes = append(nodes, v)
	}

	return nodes
}

func parseExpression(a *auditor) ast.Expr {
	return parseLiteral(a)
}

func parseLiteral(a *auditor) ast.Expr {
	return ast.Literal{
		Token: a.expectIf(token.IsLiteral, "literal"),
	}
}
