// Package grouper splits a token stream into a lists of tokens that each
// represent a token statement, a statement or fragement of a statement
// comprised of the tokens that are parsed to form it. To use, call the Begin
// function with a LexemeReader to get the first GroupTokens function. Invoking
// it will return a token statement and the next GroupTokens function.
package grouper

import (
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

// GroupTokens is a recursion based function that returns the next token
// statement, list of tokens that form a statement or statement fragment. On
// error or while obtaining the last token statement, the function will be nil.
type GroupTokens func() (token.Statement, GroupTokens, error)

// Begin returns a new GroupTokens function from which to begin parsing token
// statements. Nil is returned if the supplied reader has already reached the
// end of its stream.
func Begin(r token.LexemeReader) GroupTokens {
	if r.More() {
		return group(r)
	}
	return nil
}

// GroupAll is a convenience function and example for grouping all [remaining]
// tokens from a reader into a token block of token statements.
func GroupAll(r token.LexemeReader) (token.Block, error) {

	var (
		block = token.Block{}
		stmt  token.Statement
		f     = Begin(r)
		e     error
	)

	for f != nil {
		stmt, f, e = f()
		if e != nil {
			return nil, e
		}
		block = append(block, stmt)
	}

	return block, nil
}

func group(r token.LexemeReader) GroupTokens {
	return func() (token.Statement, GroupTokens, error) {

		stmt, e := sliceStmt(r)
		if e != nil {
			return stmt, nil, e
		}

		if r.More() {
			return stmt, group(r), nil
		}

		return stmt, nil, nil
	}
}

func sliceStmt(r token.LexemeReader) (token.Statement, error) {

	var stmt token.Statement

	for r.More() {

		lx, e := r.Read()
		if e != nil {
			return nil, e
		}

		if lx.Token == token.TK_NEWLINE {
			break
		}

		stmt = append(stmt, lx)
	}

	return stmt, nil
}
