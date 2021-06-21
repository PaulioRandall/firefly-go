package parser

import (
	"fmt"

	"github.com/go-errors/errors"

	"github.com/PaulioRandall/firefly-go/pkg/ast"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

// StmtParser is a recursion based function that parses its statement and then
// returns a parser for the next statement. On error or while obtaining the last
// AST tree, the function will be nil.
type StmtParser func() (ast.Node, StmtParser, error)

// Begin returns a new StmtParser function.
func Begin(sr token.StmtReader) StmtParser {
	if sr.More() {
		return nextParser(sr)
	}
	return nil
}

// ParseAll parses all statement in the statement reader.
func ParseAll(sr token.StmtReader) (ast.Program, error) {

	var (
		p = ast.Program{}
		n ast.Node
		f = Begin(sr)
		e error
	)

	for f != nil {
		n, f, e = f()
		if e != nil {
			return nil, e
		}
		p = append(p, n)
	}

	return p, nil
}

func nextParser(sr token.StmtReader) StmtParser {
	return func() (ast.Node, StmtParser, error) {

		unparsed, e := sr.Read()
		if e != nil {
			return nil, nil, e
		}

		parsed, e := ParseStmt(unparsed)
		if e != nil {
			return nil, nil, e
		}

		if sr.More() {
			return parsed, nextParser(sr), nil
		}
		return parsed, nil, nil
	}
}

// ParseStmt parses the supplied statement into an AST.
func ParseStmt(stmt token.Statement) (ast.Node, error) {

	lr := token.NewSliceLexemeReader(stmt)
	if !lr.More() {
		return ast.EmptyNode{}, nil
	}

	n, e := expectExpr(lr, 0)
	if e != nil {
		return nil, e
	}

	e = validateNoMoreTokens(lr)
	if e != nil {
		return nil, e
	}

	return n, nil
}

func validateNoMoreTokens(lr token.LexemeReader) error {
	if !lr.More() {
		return nil
	}

	lx, e := lr.Read()
	if e != nil {
		return e
	}

	return newError("Unexpected dangling token '%s'", lx.Token.String())
}

func newError(msg string, args ...interface{}) error {
	e := fmt.Errorf(msg, args...)
	return errors.Wrap(e, 1)
}
