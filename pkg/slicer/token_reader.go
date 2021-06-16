package slicer

import (
	"errors"

	"github.com/PaulioRandall/firefly-go/pkg/token"
)

// TokenReader is the interface for accessing scanned tokens.
type TokenReader interface {

	// More returns true if there are unread tokens.
	More() bool

	// Read returns the next token and moves the read head to the next item.
	Read() (token.Lexeme, error)
}

func SliceTokenReader(tks []token.Lexeme) sliceTokenReader {
	return sliceTokenReader{
		tks: tks,
	}
}

type sliceTokenReader struct {
	idx int
	tks []token.Lexeme
}

func (str *sliceTokenReader) More() bool {
	return len(str.tks) > str.idx
}

func (str *sliceTokenReader) Read() (token.Lexeme, error) {
	if !str.More() {
		return token.Lexeme{}, errors.New("EOF")
	}
	tk := str.tks[str.idx]
	str.idx++
	return tk, nil
}
