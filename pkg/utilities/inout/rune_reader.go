package inout

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
)

type runeReader struct {
	src Reader[rune]
	pos pos.Pos
}

func NewRuneReader(r Reader[rune]) *runeReader {
	return &runeReader{
		src: r,
	}
}

func (r runeReader) More() bool {
	return r.src.More()
}

func (r *runeReader) Peek() (rune, error) {
	ru, e := r.src.Peek()
	if e != nil {
		return rune(0), e // TODO: Wrap error with pos
	}
	return ru, nil
}

func (r *runeReader) Read() (rune, error) {
	ru, e := r.src.Read()
	if e != nil {
		return rune(0), e // TODO: Wrap error with pos
	}

	r.pos.IncRune(ru)
	return ru, nil
}

func (r runeReader) Where() pos.Pos {
	return r.pos
}
