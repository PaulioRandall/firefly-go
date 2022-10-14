package inout

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/pos"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

type readerOfRunes struct {
	src Reader[rune]
	pos pos.Pos
}

func NewReaderOfRunes(r Reader[rune]) *readerOfRunes {
	return &readerOfRunes{
		src: r,
	}
}

func (r readerOfRunes) More() bool {
	return r.src.More()
}

func (r *readerOfRunes) Peek() (rune, error) {
	ru, e := r.src.Peek()
	if e != nil {
		return rune(0), err.WrapPos(e, r.pos, "Failed to read rune")
	}
	return ru, nil
}

func (r *readerOfRunes) Read() (rune, error) {
	ru, e := r.src.Read()
	if e != nil {
		return rune(0), err.WrapPos(e, r.pos, "Failed to read rune")
	}

	r.pos.Increment(ru)
	return ru, nil
}

func (r readerOfRunes) Where() pos.Pos {
	return r.pos
}
