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
func Begin(r token.StmtReader) StmtParser {
	if r.More() {
		return nextParser(r)
	}
	return nil
}

// ParseAll parses all statement in the statement reader.
func ParseAll(r token.StmtReader) (ast.Program, error) {

	var (
		p = ast.Program{}
		n ast.Node
		f = Begin(r)
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

func nextParser(r token.StmtReader) StmtParser {
	return func() (n ast.Node, f StmtParser, e error) {

		defer func() {
			r := recover()
			if r == nil {
				return
			}

			var ok bool
			if e, ok = r.(error); !ok {
				panic("[BUG] All parse panics must recover as an error!")
			}
		}()

		unparsed, e := r.Read()
		if e != nil {
			return
		}

		n, e = ParseStmt(unparsed)
		if e != nil {
			return
		}

		if r.More() {
			f = nextParser(r)
		}

		return
	}
}

// ParseStmt parses the supplied statement into an AST.
func ParseStmt(stmt token.Statement) (ast.Node, error) {

	r := token.NewLexemeReader(stmt)
	if !r.More() {
		return ast.EmptyNode{}, nil
	}

	n, e := expectExpr(r, 0)
	if e != nil {
		return nil, e
	}

	e = validateNoMoreTokens(r)
	if e != nil {
		return nil, e
	}

	return n, nil
}

func validateNoMoreTokens(r token.LexemeReader) error {
	if !r.More() {
		return nil
	}

	lx, e := r.Read()
	if e != nil {
		return e
	}

	return newError("Unexpected dangling token '%s'", lx.Token.String())
}

func newError(msg string, args ...interface{}) error {
	e := fmt.Errorf(msg, args...)
	return errors.Wrap(e, 1)
}
