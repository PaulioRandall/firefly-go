package readers

import (
	"github.com/PaulioRandall/firefly-go/workflow/err"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

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
	return rune(0), err.EOF
}

func (r *strReader) Read() (rune, error) {
	ru, _ := r.Peek()

	if ru == rune(0) {
		return rune(0), err.EOF
	}

	r.pos.IncRune(ru)
	return ru, nil
}
