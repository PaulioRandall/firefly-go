package slicer

import (
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

// Statement in the form of a slice of lexemes
type Statement []token.Lexeme

// NextStatement is a recursion based function that returns the next slice of
// lexemes that represent a statement. On error or while obtaining the last
// statement, the function will be nil.
type NextStatement func() (Statement, NextStatement, error)

// Begin returns a new NextStatement function.
func Begin(lr LexemeReader) NextStatement {
	if lr.More() {
		return slice(lr)
	}
	return nil
}

// SliceAll converts all tokens into a group of statements.
func SliceAll(lr LexemeReader) ([]Statement, error) {

	var (
		stmt []Statement
		lx   Statement
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

func slice(lr LexemeReader) NextStatement {
	return func() (Statement, NextStatement, error) {

		stmt, e := sliceStmt(lr)
		if e != nil {
			return stmt, nil, e
		}

		if lr.More() {
			return stmt, slice(lr), nil
		}

		return stmt, nil, nil
	}
}

func sliceStmt(lr LexemeReader) (Statement, error) {

	var stmt Statement

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
