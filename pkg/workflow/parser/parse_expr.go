package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func acceptExpressions(r PosReaderOfTokens) []ast.Expr {
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

func acceptExpression(r PosReaderOfTokens) ast.Expr {
	return acceptLiteral(r)
}

func acceptLiteral(r PosReaderOfTokens) ast.Expr {
	if !acceptFunc(r, token.IsLiteral) {
		return nil
	}

	return ast.Literal{
		Token: r.Prev(),
	}
}

func expectExpressions(r PosReaderOfTokens) []ast.Expr {
	var nodes []ast.Expr

	v := expectExpression(r)
	nodes = append(nodes, v)

	for accept(r, token.Comma) {
		v := expectExpression(r)
		nodes = append(nodes, v)
	}

	return nodes
}

func expectExpression(r PosReaderOfTokens) ast.Expr {
	return expectLiteral(r)
}

func expectLiteral(r PosReaderOfTokens) ast.Expr {
	return ast.Literal{
		Token: expectFunc(r, "literal", token.IsLiteral),
	}
}
