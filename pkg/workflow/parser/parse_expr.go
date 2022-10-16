package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func acceptExpressions(r BufReaderOfTokens) []ast.Expr {
	var nodes []ast.Expr

	for r.More() {
		v := acceptExpression(r)
		if v == nil {
			break
		}

		nodes = append(nodes, v)
	}

	return nodes
}

func acceptExpression(r BufReaderOfTokens) ast.Expr {
	return acceptLiteral(r)
}

func acceptLiteral(r BufReaderOfTokens) ast.Expr {
	if !acceptFunc(r, token.IsLiteral) {
		return nil
	}

	return ast.Literal{
		Token: r.Prev(),
	}
}

func expectExpressions(r BufReaderOfTokens) []ast.Expr {
	var nodes []ast.Expr

	v := expectExpression(r)
	nodes = append(nodes, v)

	for accept(r, token.Comma) {
		v := expectExpression(r)
		nodes = append(nodes, v)
	}

	return nodes
}

func expectExpression(r BufReaderOfTokens) ast.Expr {
	return expectLiteral(r)
}

func expectLiteral(r BufReaderOfTokens) ast.Expr {
	return ast.Literal{
		Token: expectFunc(r, "literal", token.IsLiteral),
	}
}
