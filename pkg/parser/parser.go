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
		parsed        = ast.Program{}
		node          ast.Node
		nextParseFunc = Begin(r)
		e             error
	)

	for nextParseFunc != nil {
		node, nextParseFunc, e = nextParseFunc()
		if e != nil {
			return nil, e
		}
		parsed = append(parsed, node)
	}

	return parsed, nil
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

		var nextParseFunc StmtParser
		if r.More() {
			nextParseFunc = nextParser(r)
		}

		return parsed, nextParseFunc, nil
	}
}

// ParseStmt parses the supplied statement into an AST.
func ParseStmt(stmt token.Statement) (n ast.Node, e error) {

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

	lr := token.NewLexemeReader(stmt)

	if lr.More() {
		n = parseStmt(lr)
	} else {
		n = ast.EmptyNode{}
	}

	return n, e
}

func parseStmt(lr token.LexemeReader) ast.Node {
	r := lexReader{lr: lr}
	n := expectExpr(r, 0)
	validateNoMoreTokens(r)
	return n
}

func validateNoMoreTokens(r lexReader) {
	if r.More() {
		lx := r.Read()
		parsingPanic(nil, "Unexpected dangling token '%s'", lx.Token.String())
	}
}
