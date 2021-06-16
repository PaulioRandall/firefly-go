package slicer

import (
	"errors"

	"github.com/PaulioRandall/firefly-go/pkg/token"
)

// LexemeReader is the interface for accessing scanned lexemes.
type LexemeReader interface {

	// More returns true if there are unread lexemes.
	More() bool

	// Read returns the next lexeme and moves the read head to the next item.
	Read() (token.Lexeme, error)
}

func NewSliceLexemeReader(lxs []token.Lexeme) *sliceLexemeReader {
	return &sliceLexemeReader{
		lxs: lxs,
	}
}

type sliceLexemeReader struct {
	idx int
	lxs []token.Lexeme
}

func (slr *sliceLexemeReader) More() bool {
	return len(slr.lxs) > slr.idx
}

func (slr *sliceLexemeReader) Read() (token.Lexeme, error) {
	if !slr.More() {
		return token.Lexeme{}, errors.New("EOF")
	}
	lx := slr.lxs[slr.idx]
	slr.idx++
	return lx, nil
}
