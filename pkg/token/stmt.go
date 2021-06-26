package token

// NewStmtReader wraps a slice of statements (program) for reading as a stream.
func NewStmtReader(b Block) *stmtReader {
	return &stmtReader{
		stmts: b,
	}
}

type stmtReader struct {
	idx   int
	stmts Block
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
