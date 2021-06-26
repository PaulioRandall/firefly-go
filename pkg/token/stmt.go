package token

// Program in the form of a slice of statements.
type Program []Statement

// Statement in the form of a slice of lexemes.
type Statement []Lexeme

// StmtReader is the interface for reading statements from some source.
type StmtReader interface {

	// More returns true if there are unread statements.
	More() bool

	// Read returns the next statement and moves the read head to the next item.
	Read() (Statement, error)
}

// NewStmtReader wraps a programs (slice of statements) for reading.
func NewStmtReader(p Program) *stmtReader {
	return &stmtReader{
		stmts: p,
	}
}

type stmtReader struct {
	idx   int
	stmts Program
}

func (r *stmtReader) More() bool {
	return len(r.stmts) > r.idx
}

func (r *stmtReader) Read() (Statement, error) {
	if !r.More() {
		return nil, EOF
	}
	stmt := r.stmts[r.idx]
	r.idx++
	return stmt, nil
}
