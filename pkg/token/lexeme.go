package token

import (
	"errors"
)

// Lexeme is a value with associated token.
type Lexeme struct {
	Token
	Value string
}

// LexemeReader is the interface for accessing scanned lexemes.
type LexemeReader interface {

	// More returns true if there are unread lexemes.
	More() bool

	// Read returns the next lexeme and moves the read head to the next item.
	Read() (Lexeme, error)
}

// NewSliceLexemeReader wraps a slice of tokens in a Lexeme reader.
func NewSliceLexemeReader(lxs []Lexeme) *sliceLexemeReader {
	return &sliceLexemeReader{
		lxs: lxs,
	}
}

type sliceLexemeReader struct {
	idx int
	lxs []Lexeme
}

func (slr *sliceLexemeReader) More() bool {
	return len(slr.lxs) > slr.idx
}

func (slr *sliceLexemeReader) Read() (Lexeme, error) {
	if !slr.More() {
		return Lexeme{}, errors.New("EOF")
	}
	lx := slr.lxs[slr.idx]
	slr.idx++
	return lx, nil
}
