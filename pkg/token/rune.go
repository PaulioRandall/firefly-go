package token

// RuneReader interface is for accessing Go runes from a text source.
type RuneReader interface {

	// More returns true if there are unread runes.
	More() bool

	// Peek returns the next rune without incrementing.
	Peek() (rune, error)

	// Read returns the next rune and increments to the next item.
	Read() (rune, error)
}

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
