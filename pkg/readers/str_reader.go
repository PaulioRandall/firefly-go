package readers

import (
	"errors"

	"github.com/PaulioRandall/firefly-go/pkg/token"
)

type strReader struct {
	pos token.Pos
	str []rune
}

func NewStringRuneReader(s string) *strReader {
	return &strReader{
		str: []rune(s),
	}
}

func (r *strReader) Pos() token.Pos {
	return r.pos
}

func (r *strReader) More() bool {
	return r.pos.Idx < len(r.str)
}

func (r *strReader) Peek() (rune, error) {
	if r.More() {
		return r.str[r.pos.Idx], nil
	}

	// TODO: Create err package
	return rune(0), errors.New("EOF: No more runes")
}

func (r *strReader) Read() (rune, error) {
	ru, _ := r.Peek()

	if ru == rune(0) {
		// TODO: Create err package
		return rune(0), errors.New("EOF: No more runes")
	}

	r.pos.Inc(ru)
	return ru, nil
}
