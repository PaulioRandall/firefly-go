package parser

import (
	"github.com/PaulioRandall/firefly-go/firefly/token"
)

// LexReader interface is for accessing a stream of lexemes.
type LexReader interface {

	// More returns true if there are unread lexemes.
	More() bool

	// Peek returns the next lexeme without incrementing to the next.
	Peek() (token.Lexeme, error)

	// Read returns the next lexeme and increments to the next item.
	Read() (token.Lexeme, error)
}

type lexReader struct {
	lr LexReader
}

func (r lexReader) More() bool {
	return r.lr.More()
}

func (r lexReader) Peek() token.Lexeme {
	lx, e := r.lr.Peek()
	if e != nil {
		lexemeReadPanic(e)
	}
	return lx
}

func (r lexReader) Read() token.Lexeme {
	lx, e := r.lr.Read()
	if e != nil {
		lexemeReadPanic(e)
	}
	return lx
}

func lexemeReadPanic(cause error) {
	parsingPanic(cause, "Lexeme read error")
}
