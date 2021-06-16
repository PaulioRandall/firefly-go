package token

import (
	"errors"
)

// Statement in the form of a slice of lexemes.
type Statement []Lexeme

// StmtReader is the interface for accessing token statement.
type StmtReader interface {

	// More returns true if there are unread statements.
	More() bool

	// Read returns the next statement and moves the read head to the next item.
	Read() (Statement, error)
}

// NewSliceStmtReader wraps a slice of tokens in a Lexeme reader.
func NewSliceStmtReader(stmts []Statement) *sliceStmtReader {
	return &sliceStmtReader{
		stmts: stmts,
	}
}

type sliceStmtReader struct {
	idx   int
	stmts []Statement
}

func (ssr *sliceStmtReader) More() bool {
	return len(ssr.stmts) > ssr.idx
}

func (ssr *sliceStmtReader) Read() (Statement, error) {
	if !ssr.More() {
		return nil, errors.New("EOF")
	}
	stmt := ssr.stmts[ssr.idx]
	ssr.idx++
	return stmt, nil
}
