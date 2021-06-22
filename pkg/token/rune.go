package token

import (
	"errors"
)

// RuneReader is the interface for accessing Go runes from a text source.
type RuneReader interface {

	// More returns true if there are unread runes.
	More() bool

	// Read returns the next rune and moves the read head to the next item.
	Read() (rune, error)

	// PutBack puts a rune back into the reader so it becomes the next one to be
	// read.
	PutBack(rune) error
}

// NewRuneReader wraps a slice of runes for reading.
func NewRuneReader(text []rune) *runeReader {
	return &runeReader{
		text: text,
	}
}

type runeReader struct {
	idx  int
	text []rune
}

func (r *runeReader) More() bool {
	return len(r.text) > r.idx
}

func (r *runeReader) Read() (rune, error) {
	if !r.More() {
		return rune(0), errors.New("EOF")
	}
	ru := r.text[r.idx]
	r.idx++
	return ru, nil
}

func (r *runeReader) PutBack(ru rune) error {
	head := r.text[:r.idx]
	tail := r.text[r.idx:]
	tail = append([]rune{ru}, tail...)
	r.text = append(head, tail...)
	return nil
}
