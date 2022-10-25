package inout

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
)

type readerOfRunes struct {
	Reader[rune]
	pos pos.Pos
}

func NewReaderOfRunes(r Reader[rune]) *readerOfRunes {
	return &readerOfRunes{
		Reader: r,
	}
}

func (r *readerOfRunes) Peek() (rune, error) {
	ru, e := r.Reader.Peek()
	if e != nil {
		return rune(0), ErrReadDelegate.WrapPos(e, r.pos, "Failed to peek rune")
	}
	return ru, nil
}

func (r *readerOfRunes) Read() (rune, error) {
	ru, e := r.Reader.Read()
	if e != nil {
		return rune(0), ErrReadDelegate.WrapPos(e, r.pos, "Failed to read rune")
	}

	r.pos.Increment(ru)
	return ru, nil
}

func (r readerOfRunes) Where() pos.Pos {
	return r.pos
}
