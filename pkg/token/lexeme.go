package token

// Lexeme is a value with associated token.
type Lexeme struct {
	Token
	Value string
}

// LexemeReader is the interface for accessing scanned lexemes.
type LexemeReader interface {

	// More returns true if there are unread lexemes.
	More() bool

	// Peek returns the next lexeme without moving the read head.
	Peek() (Lexeme, error)

	// Read returns the next lexeme and moves the read head to the next item.
	Read() (Lexeme, error)
}

// NewLexemeReader wraps a slice of tokens in a Lexeme reader.
func NewLexemeReader(lxs []Lexeme) *lexemeReader {
	return &lexemeReader{
		lxs: lxs,
	}
}

type lexemeReader struct {
	idx int
	lxs []Lexeme
}

func (r *lexemeReader) More() bool {
	return len(r.lxs) > r.idx
}

func (r *lexemeReader) Peek() (Lexeme, error) {
	if !r.More() {
		return Lexeme{}, EOF
	}
	return r.lxs[r.idx], nil
}

func (r *lexemeReader) Read() (Lexeme, error) {
	if !r.More() {
		return Lexeme{}, EOF
	}
	lx := r.lxs[r.idx]
	r.idx++
	return lx, nil
}
