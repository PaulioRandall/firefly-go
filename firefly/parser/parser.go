package parser

import (
	"github.com/PaulioRandall/firefly-go/firefly/ast"
	"github.com/PaulioRandall/firefly-go/firefly/token"
)

// StmtParser is a recursion based function that parses its statement and then
// returns a parser for the next statement. On error or while obtaining the last
// AST tree, the function will be nil.
type StmtParser func() (ast.Tree, StmtParser, error)

// StmtReader interface is for reading statements from a stream.
type StmtReader interface {

	// More returns true if there are unread statements.
	More() bool

	// Read returns the next statement and increments to the next item.
	Read() (token.Statement, error)
}

// Begin returns a new StmtParser function.
func Begin(r StmtReader) StmtParser {
	if r.More() {
		return nextParser(r)
	}
	return nil
}

// ParseAll parses all statement in the statement reader.
func ParseAll(r StmtReader) (ast.Block, error) {

	var (
		parsed        = ast.Block{}
		tree          ast.Tree
		nextParseFunc = Begin(r)
		e             error
	)

	for nextParseFunc != nil {
		tree, nextParseFunc, e = nextParseFunc()
		if e != nil {
			return nil, e
		}
		parsed = append(parsed, tree)
	}

	return parsed, nil
}

func nextParser(r StmtReader) StmtParser {
	return func() (ast.Tree, StmtParser, error) {

		unparsed, e := r.Read()
		if e != nil {
			return nil, nil, e
		}

		parsed, e := ParseStmt(unparsed)
		if e != nil {
			return nil, nil, e
		}

		var nextParseFunc StmtParser
		if r.More() {
			nextParseFunc = nextParser(r)
		}

		return parsed, nextParseFunc, nil
	}
}

// ParseStmt parses the supplied statement into an AST.
func ParseStmt(stmt token.Statement) (tr ast.Tree, e error) {

	defer func() {
		err := recover()
		if err == nil {
			return
		}

		var ok bool
		if e, ok = err.(error); !ok {
			panic("[BUG] All parse panics must recover as an error!")
		}
	}()

	lr := token.NewLexReader(stmt)

	if lr.More() {
		tr = parseStmt(lr)
	} else {
		tr = ast.EmptyTree{}
	}

	return tr, e
}

func parseStmt(lr LexReader) ast.Tree {
	r := lexReader{lr: lr}
	tr := expectExpr(r, 0)
	validateNoMoreTokens(r)
	return tr
}

func validateNoMoreTokens(r lexReader) {
	if r.More() {
		tk := r.Peek().Token.String()
		parsingPanic(nil, "Unexpected dangling token '%s'", tk)
	}
}
