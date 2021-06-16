package token

import (
	"errors"
)

// Program in the form of a slice of statements formed of a slice of lexemes.
type Program []Statement

// Statement in the form of a slice of lexemes.
type Statement []Lexeme

// StmtReader is the interface for accessing token statement.
type StmtReader interface {

	// More returns true if there are unread statements.
	More() bool

	// Read returns the next statement and moves the read head to the next item.
	Read() (Statement, error)
}

// NewProgramReader wraps a slice of tokens in a Lexeme reader.
func NewProgramReader(p Program) *programReader {
	return &programReader{
		stmts: p,
	}
}

type programReader struct {
	idx   int
	stmts []Statement
}

func (ssr *programReader) More() bool {
	return len(ssr.stmts) > ssr.idx
}

func (ssr *programReader) Read() (Statement, error) {
	if !ssr.More() {
		return nil, errors.New("EOF")
	}
	stmt := ssr.stmts[ssr.idx]
	ssr.idx++
	return stmt, nil
}
