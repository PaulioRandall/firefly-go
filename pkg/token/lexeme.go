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

	// Read returns the next lexeme and moves the read head to the next item.
	Read() (Lexeme, error)

	// PutBack puts a lexeme back into the reader so it becomes the next lexeme
	// to be read.
	PutBack(Lexeme) error
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

func (r *lexemeReader) Read() (Lexeme, error) {
	if !r.More() {
		return Lexeme{}, EOF
	}
	lx := r.lxs[r.idx]
	r.idx++
	return lx, nil
}

func (r *lexemeReader) PutBack(lx Lexeme) error {
	head := r.lxs[:r.idx]
	tail := r.lxs[r.idx:]
	tail = append([]Lexeme{lx}, tail...)
	r.lxs = append(head, tail...)
	return nil
}
