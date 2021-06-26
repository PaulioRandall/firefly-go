package token

// Lexeme is a container for a value and the token it evaluates to. Sometimes
// a lexeme will be referred to as a token. This is because the token drives
// the logic, the value is required only in some cases, e.g. for number tokens.
type Lexeme struct {
	Token
	Value string
}

// NewLexReader wraps a slice of lexemes for reading as a stream.
func NewLexReader(lxs []Lexeme) *lexReader {
	return &lexReader{
		lxs: lxs,
	}
}

type lexReader struct {
	idx int
	lxs []Lexeme
}

func (r *lexReader) More() bool {
	return len(r.lxs) > r.idx
}

func (r *lexReader) Peek() (Lexeme, error) {
	if !r.More() {
		return Lexeme{}, EOF
	}
	return r.lxs[r.idx], nil
}

func (r *lexReader) Read() (Lexeme, error) {
	if !r.More() {
		return Lexeme{}, EOF
	}
	lx := r.lxs[r.idx]
	r.idx++
	return lx, nil
}
