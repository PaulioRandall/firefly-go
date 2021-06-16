package slicer

import (
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

// NextStatement is a recursion based function that returns the next slice of
// lexemes that represent a statement. On error or while obtaining the last
// statement, the function will be nil.
type NextStatement func() (token.Statement, NextStatement, error)

// Begin returns a new NextStatement function.
func Begin(lr LexemeReader) NextStatement {
	if lr.More() {
		return group(lr)
	}
	return nil
}

// GroupAll converts all tokens into a group of statements.
func GroupAll(lr LexemeReader) ([]token.Statement, error) {

	var (
		stmt []token.Statement
		lx   token.Statement
		f    = Begin(lr)
		e    error
	)

	for f != nil {
		lx, f, e = f()
		if e != nil {
			return nil, e
		}
		stmt = append(stmt, lx)
	}

	return stmt, nil
}

func group(lr LexemeReader) NextStatement {
	return func() (token.Statement, NextStatement, error) {

		stmt, e := sliceStmt(lr)
		if e != nil {
			return stmt, nil, e
		}

		if lr.More() {
			return stmt, group(lr), nil
		}

		return stmt, nil, nil
	}
}

func sliceStmt(lr LexemeReader) (token.Statement, error) {

	var stmt token.Statement

	for lr.More() {

		lx, e := lr.Read()
		if e != nil {
			return nil, e
		}

		if lx.Token == token.TokenNewline {
			break
		}

		stmt = append(stmt, lx)
	}

	return stmt, nil
}
