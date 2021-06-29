// Package cleaner removes redundant tokens and applies any replacement rules
// to a stream of token statements.
//
// To use, call the Begin function with a StmtReader to get the first
// CleanStatement function. Invoking it will return a cleaned token statement
// and the next CleanStatement function.
package cleaner

import (
	"github.com/PaulioRandall/firefly-go/firefly/token"
)

// CleanStatement is a recursion based function that returns the next token
// statement, list of tokens that form a statement or statement fragment. On
// error or while obtaining the last token statement, the function will be nil.
type CleanStatement func() (token.Statement, CleanStatement, error)

// StmtReader interface is for reading statements from a stream.
type StmtReader interface {

	// More returns true if there are unread statements.
	More() bool

	// Read returns the next statement and increments to the next item.
	Read() (token.Statement, error)
}

// Begin returns a new CleanStatement function from which to begin cleaning
// token statements. Nil is returned if the supplied reader has already reached
// the end of its stream.
func Begin(sr StmtReader) CleanStatement {
	if sr.More() {
		return clean(sr)
	}
	return nil
}

// CleanAll is a convenience function and example for cleaning all [remaining]
// token statements from a reader into a token block.
func CleanAll(sr StmtReader) (token.Block, error) {

	var (
		stmts token.Block
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

func clean(sr StmtReader) CleanStatement {
	return func() (token.Statement, CleanStatement, error) {

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

	for _, lx := range unclean {
		if !lx.Token.IsRedundant() {
			cleaned = append(cleaned, lx)
		}
	}

	return cleaned
}
