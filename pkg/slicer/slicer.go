package slicer

import (
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

// NextStatement is a recursion based function that returns the next slice of
// tokens that represent a statement. On error or while obtaining the last
// statement, the function will be nil.
type NextStatement func() ([]token.Lexeme, NextStatement, error)

// Begin returns a new NextStatement function.
func Begin(lr LexemeReader) NextStatement {
	if lr.More() {
		return slice(lr)
	}
	return nil
}

// SliceAll converts all tokens into a group of statements.
func SliceAll(lr LexemeReader) ([][]token.Lexeme, error) {

	var (
		result [][]token.Lexeme
		lx     []token.Lexeme
		f      = Begin(lr)
		e      error
	)

	for f != nil {
		lx, f, e = f()
		if e != nil {
			return nil, e
		}
		result = append(result, lx)
	}

	return result, nil
}

func slice(lr LexemeReader) NextStatement {
	return func() ([]token.Lexeme, NextStatement, error) {

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

func sliceStmt(lr LexemeReader) ([]token.Lexeme, error) {

	var stmt []token.Lexeme

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
