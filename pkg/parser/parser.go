package parser

import (
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
	return func() (ast.Node, StmtParser, error) {

		unparsed, e := r.Read()
		if e != nil {
			return nil, nil, e
		}

		parsed, e := ParseStmt(unparsed)
		if e != nil {
			return nil, nil, e
		}

		var f StmtParser
		if r.More() {
			f = nextParser(r)
		}

		return parsed, f, nil
	}
}

// ParseStmt parses the supplied statement into an AST.
func ParseStmt(stmt token.Statement) (n ast.Node, e error) {

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

	lr := token.NewLexemeReader(stmt)
	r := lexReader{lr: lr}

	if !r.More() {
		return ast.EmptyNode{}, nil
	}

	n = expectExpr(r, 0)
	validateNoMoreTokens(r)

	return n, nil
}

func validateNoMoreTokens(r lexReader) {
	if r.More() {
		lx := r.Read()
		panicParseErr(nil, "Unexpected dangling token '%s'", lx.Token.String())
	}
}
