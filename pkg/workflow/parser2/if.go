package parser2

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// == If
func isIfStatement(r BufReaderOfTokens) bool {
	return peekType(r) == token.If
}

// IF := If EXPRESSION TERM STATEMENT_BLOCK
func parseIfStatement(r BufReaderOfTokens) ast.If {
	n := ast.If{}

	expectType(r, token.If)
	n.Condition = parseExpression(r)
	expectEndOfStmt(r)
	n.Body = parseStatementBlock(r)

	return n
}
