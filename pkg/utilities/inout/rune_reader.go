package inout

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/err"
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
		return rune(0), err.WrapPos(e, r.pos, "Failed to read rune")
	}
	return ru, nil
}

func (r *runeReader) Read() (rune, error) {
	ru, e := r.src.Read()
	if e != nil {
		return rune(0), err.WrapPos(e, r.pos, "Failed to read rune")
	}

	r.pos.Increment(ru)
	return ru, nil
}

func (r runeReader) Where() pos.Pos {
	return r.pos
}
