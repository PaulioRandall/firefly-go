package slicer

import (
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

// TokenReader is the interface for accessing scanned tokens.
type TokenReader interface {

	// More returns true if there are unread tokens.
	More() bool

	// Read returns the next token and moves the read head to the next item.
	Read() (token.Lexeme, error)
}

// NextStatement is a recursion based function that returns the next slice of
// tokens that represent a statement. On error or while obtaining the last
// statement, the function will be nil.
type NextStatement func() ([]token.Lexeme, NextStatement, error)

// Begin returns a new NextStatement function.
func Begin(tr TokenReader) NextStatement {
	if tr.More() {
		return slice(tr)
	}
	return nil
}

// SliceAll converts all tokens into a group of statements.
func SliceAll(tr TokenReader) ([][]token.Lexeme, error) {

	var (
		result [][]token.Lexeme
		lx     []token.Lexeme
		f      = Begin(tr)
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

func slice(tr TokenReader) NextStatement {
	return func() ([]token.Lexeme, NextStatement, error) {

		stmt, e := sliceStmt(tr)
		if e != nil {
			return stmt, nil, e
		}

		if tr.More() {
			return stmt, slice(tr), nil
		}

		return stmt, nil, nil
	}
}

func sliceStmt(tr TokenReader) ([]token.Lexeme, error) {

	var stmt []token.Lexeme

	for tr.More() {

		lx, e := tr.Read()
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
