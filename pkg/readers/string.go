package readers

import (
	"errors"
)

type strReader struct {
	idx int
	str []rune
}

func FromString(s string) strReader {
	return strReader{
		str: []rune(s),
	}
}

func (r *strReader) More() bool {
	return r.idx < len(r.str)
}

func (r *strReader) Peek() (rune, error) {
	if r.More() {
		return r.str[r.idx], nil
	}
	return rune(0), nil
}

func (r *strReader) Read() (rune, error) {
	ru, _ := r.Peek()

	if ru == rune(0) {
		// TODO: Create err package
		return rune(0), errors.New("EOF: No more runes")
	}

	r.idx++
	return ru, nil
}
