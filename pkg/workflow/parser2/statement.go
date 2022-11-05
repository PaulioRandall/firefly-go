package parser2

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// STATEMENT_BLOCK := {STATEMENT} End
func parseStatementBlock(r BufReaderOfTokens) []ast.Stmt {
	var stmts []ast.Stmt

	acceptEmptyStatements(r)
	for !acceptType(r, token.End) {
		acceptEmptyStatements(r)
		n := parseTermStatement(r)
		stmts = append(stmts, n)
	}

	if r.Prev().TokenType != token.End {
		panic(ErrParsing.Track("Unterminated statement block"))
	}

	return stmts
}

// EMPTY_STATEMENTS := {TERM}
func acceptEmptyStatements(r BufReaderOfTokens) {
	for acceptEndOfStmt(r) {
	}
}

// TERM_STATEMENT := STATEMENT TERM
func parseTermStatement(r BufReaderOfTokens) ast.Stmt {
	n := parseStatement(r)
	expectEndOfStmt(r)
	return n
}

// STATEMENT := ASSIGNMENT
func parseStatement(r BufReaderOfTokens) ast.Stmt {
	switch {
	case isAssignment(r):
		return parseAssignment(r)
	case isIfStatement(r):
		return parseIfStatement(r)
	default:
		panic(ErrParsing.Track("Expected statement"))
	}
}

// := [TERM]
func acceptEndOfStmt(r BufReaderOfTokens) bool {
	return acceptType(r, token.Terminator) || acceptType(r, token.Newline)
}

// TERM := Terminator | Newline
func expectEndOfStmt(r BufReaderOfTokens) {
	if acceptEndOfStmt(r) {
		return
	}

	panic(ErrParsing.Track("Expected terminator or newline"))
}
