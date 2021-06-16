package scanner

import (
	"errors"
)

// ScrollReader is the interface for accessing Go runes from a text source.
type ScrollReader interface {

	// More returns true if there are unread runes.
	More() bool

	// Read returns the next rune in the scroll and moves the read head to the
	// next item.
	Read() (rune, error)

	// PutBack puts a rune back into the scoll reader so it becomes the next
	// rune to be read.
	PutBack(rune) error
}

// NewStringScrollReader wraps a slice of runes in a scroll reader.
func NewStringScrollReader(text []rune) *stringScrollReader {
	return &stringScrollReader{
		text: text,
	}
}

type stringScrollReader struct {
	idx  int
	text []rune
}

func (ssr *stringScrollReader) More() bool {
	return len(ssr.text) > ssr.idx
}

func (ssr *stringScrollReader) Read() (rune, error) {
	if !ssr.More() {
		return rune(0), errors.New("EOF")
	}
	ru := ssr.text[ssr.idx]
	ssr.idx++
	return ru, nil
}

func (ssr *stringScrollReader) PutBack(ru rune) error {
	remaining := ssr.text[ssr.idx:]
	alreadyRead := append(ssr.text[:ssr.idx], ru)
	ssr.text = append(alreadyRead, remaining...)
	return nil
}
