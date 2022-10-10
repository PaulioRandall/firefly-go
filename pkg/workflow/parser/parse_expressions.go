package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func expectExpressions(a *auditor) []ast.Expr {
	var nodes []ast.Expr

	v := expectExpression(a)
	nodes = append(nodes, v)

	for a.accept(token.Comma) {
		v := expectExpression(a)
		nodes = append(nodes, v)
	}

	return nodes
}

func expectExpression(a *auditor) ast.Expr {
	return expectLiteral(a)
}

func expectLiteral(a *auditor) ast.Expr {
	return ast.Literal{
		Token: a.expectIf(token.IsLiteral, "literal"),
	}
}
