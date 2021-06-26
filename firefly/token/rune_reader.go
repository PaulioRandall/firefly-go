package token

// NewRuneReader wraps a slice of runes for reading as a stream.
func NewRuneReader(text []rune) *runeReader {
	return &runeReader{
		text: text,
	}
}

type runeReader struct {
	idx  int
	text []rune
}

func (r *runeReader) More() bool {
	return len(r.text) > r.idx
}

func (r *runeReader) Peek() (rune, error) {
	if !r.More() {
		return rune(0), EOF
	}
	return r.text[r.idx], nil
}

func (r *runeReader) Read() (rune, error) {
	if !r.More() {
		return rune(0), EOF
	}
	ru := r.text[r.idx]
	r.idx++
	return ru, nil
}
