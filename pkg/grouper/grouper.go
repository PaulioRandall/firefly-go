package grouper

import (
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

// NextStatement is a recursion based function that returns the next slice of
// lexemes that represent a statement. On error or while obtaining the last
// statement, the function will be nil.
type NextStatement func() (token.Statement, NextStatement, error)

// Begin returns a new NextStatement function.
func Begin(r token.LexemeReader) NextStatement {
	if r.More() {
		return group(r)
	}
	return nil
}

// GroupAll converts all tokens into a group of statements.
func GroupAll(r token.LexemeReader) (token.Program, error) {

	var (
		prog = token.Program{}
		stmt token.Statement
		f    = Begin(r)
		e    error
	)

	for f != nil {
		stmt, f, e = f()
		if e != nil {
			return nil, e
		}
		prog = append(prog, stmt)
	}

	return prog, nil
}

func group(r token.LexemeReader) NextStatement {
	return func() (token.Statement, NextStatement, error) {

		stmt, e := sliceStmt(r)
		if e != nil {
			return stmt, nil, e
		}

		if r.More() {
			return stmt, group(r), nil
		}

		return stmt, nil, nil
	}
}

func sliceStmt(r token.LexemeReader) (token.Statement, error) {

	var stmt token.Statement

	for r.More() {

		lx, e := r.Read()
		if e != nil {
			return nil, e
		}

		if lx.Token == token.TK_NEWLINE {
			break
		}

		stmt = append(stmt, lx)
	}

	return stmt, nil
}
