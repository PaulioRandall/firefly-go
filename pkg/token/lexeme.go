package token

// Lexeme is a container for a value and the token it evaluates to. Sometimes
// a lexeme will be referred to as a token. This is because the token drives
// the logic, the value is required only in some cases, e.g. for number tokens.
type Lexeme struct {
	Token
	Value string
}

// LexemeReader interface is for accessing a stream of lexemes.
type LexemeReader interface {

	// More returns true if there are unread lexemes.
	More() bool

	// Peek returns the next lexeme without incrementing to the next.
	Peek() (Lexeme, error)

	// Read returns the next lexeme and increments to the next item.
	Read() (Lexeme, error)
}

// NewLexemeReader wraps a slice of lexemes for reading as a stream.
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
