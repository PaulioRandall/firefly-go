package parser2

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// EXPRESSIONS := [EXPRESSION {Comma EXPRESSION}]
func parseExpressions(r BufReaderOfTokens) []ast.Expr {
	var result []ast.Expr

	if !isExpression(r) {
		return result
	}

	result = append(result, parseExpression(r))
	for acceptType(r, token.Comma) {
		result = append(result, parseExpression(r))
	}

	return result
}

// == LITERAL
func isExpression(r BufReaderOfTokens) bool {
	return isLiteral(r)
}

// EXPRESSION := LITERAL
func parseExpression(r BufReaderOfTokens) ast.Expr {
	switch {
	case isLiteral(r):
		return parseLiteral(r)
	default:
		panic(ErrParsing.Track("Expected expression"))
	}
}
