package runereader

import (
	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type strRuneReader struct {
	pos   token.Pos
	runes []rune
}

func FromString(s string) *strRuneReader {
	return &strRuneReader{
		runes: []rune(s),
	}
}

func (rr strRuneReader) Pos() token.Pos {
	return rr.pos
}

func (rr strRuneReader) More() bool {
	return rr.pos.Offset < len(rr.runes)
}

func (rr strRuneReader) Peek() (rune, error) {
	if rr.More() {
		return rr.runes[rr.pos.Offset], nil
	}
	return rune(0), err.EOF
}

func (rr *strRuneReader) Read() (rune, error) {
	ru, _ := rr.Peek()

	if ru == rune(0) {
		return rune(0), err.EOF
	}

	rr.pos.IncRune(ru)
	return ru, nil
}
