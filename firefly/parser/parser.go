// Package parser converts a token statements into Abstract Syntax Trees (AST).
//
// To use, call the Begin function with a StmtReader to get the first
// ParseStatement function. Invoking it will return a token statement and the
// next ParseStatement function.
package parser

import (
	"github.com/PaulioRandall/firefly-go/firefly/ast"
	"github.com/PaulioRandall/firefly-go/firefly/token"
)

// ParseStatement is a recursion based function that parses its token statement
// into an AST. On error or while obtaining the last AST, the function will be
// nil.
type ParseStatement func() (ast.Tree, ParseStatement, error)

// StmtReader interface is for reading statements from a stream.
type StmtReader interface {

	// More returns true if there are unread statements.
	More() bool

	// Read returns the next statement and increments to the next item.
	Read() (token.Statement, error)
}

// Begin returns a new ParseStatement function from which to begin parsing
// token statements. Nil is returned if the supplied reader has already reached
// the end of its stream.
func Begin(r StmtReader) ParseStatement {
	if r.More() {
		return nextParser(r)
	}

	return func() (ast.Tree, ParseStatement, error) {
		return ast.EmptyTree{}, nil, nil
	}
}

// ParseAll is a convenience function and example for parsing all [remaining]
// token statements from a reader into a AST block.
func ParseAll(r StmtReader) (ast.Block, error) {

	var (
		parsed        ast.Block
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

func nextParser(r StmtReader) ParseStatement {
	return func() (ast.Tree, ParseStatement, error) {

		unparsed, e := r.Read()
		if e != nil {
			return nil, nil, e
		}

		parsed, e := Parse(unparsed)
		if e != nil {
			return nil, nil, e
		}

		var nextParseFunc ParseStatement
		if r.More() {
			nextParseFunc = nextParser(r)
		}

		return parsed, nextParseFunc, nil
	}
}

// Parse parses a statement into an AST.
func Parse(stmt token.Statement) (tr ast.Tree, e error) {

	// Due to the intrinsic complexity of parsing, panics are generated for error
	// handling and recovered here. The error is then returned.
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
		parsingPanic(nil, "unexpected dangling token '%s'", tk)
	}
}
