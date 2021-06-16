package cleaner

import (
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

// NextStatement is a recursion based function that returns the next slice of
// lexemes that represent a statement. On error or while obtaining the last
// statement, the function will be nil.
type NextStatement func() (token.Statement, NextStatement, error)

// Begin returns a new NextStatement function.
func Begin(sr token.StmtReader) NextStatement {
	if sr.More() {
		return clean(sr)
	}
	return nil
}

// CleanAll removes redundant tokens from a stream of statements.
func CleanAll(sr token.StmtReader) (token.Program, error) {

	var (
		stmts token.Program
		stmt  token.Statement
		f     = Begin(sr)
		e     error
	)

	for f != nil {
		stmt, f, e = f()
		if e != nil {
			return nil, e
		}
		stmts = append(stmts, stmt)
	}

	return stmts, nil
}

func clean(sr token.StmtReader) NextStatement {
	return func() (token.Statement, NextStatement, error) {

		unclean, e := sr.Read()
		if e != nil {
			return nil, nil, e
		}

		stmt := CleanStmt(unclean)

		if sr.More() {
			return stmt, clean(sr), nil
		}

		return stmt, nil, nil
	}
}

func CleanStmt(unclean token.Statement) token.Statement {

	cleaned := token.Statement{}
	isRedundant := func(tk token.Token) bool {
		return tk == token.TokenSpace
	}

	for _, lx := range unclean {
		if !isRedundant(lx.Token) {
			cleaned = append(cleaned, lx)
		}
	}

	return cleaned
}
