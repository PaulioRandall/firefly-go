package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/auditor"
)

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
		Operator: a.ExpectFunc("literal", token.IsLiteral),
	}
}
