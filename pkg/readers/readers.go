package readers

import (
	"errors"

	"github.com/PaulioRandall/firefly-go/pkg/token"
)

var EOF = errors.New("EOF, no more runes")

type strReader struct {
	pos   token.Pos
	runes []rune
}

func NewRuneStringReader(s string) *strReader {
	return &strReader{
		runes: []rune(s),
	}
}

func (r *strReader) Pos() token.Pos {
	return r.pos
}

func (r *strReader) More() bool {
	return r.pos.Offset < len(r.runes)
}

func (r *strReader) Peek() (rune, error) {
	if r.More() {
		return r.runes[r.pos.Offset], nil
	}

	// TODO: Create err package
	return rune(0), EOF
}

func (r *strReader) Read() (rune, error) {
	ru, _ := r.Peek()

	if ru == rune(0) {
		// TODO: Create err package
		return rune(0), EOF
	}

	r.pos.Inc(ru)
	return ru, nil
}
