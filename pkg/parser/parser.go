package parser

import (
	"errors"
	"fmt"
	"strconv"

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
func ParseAll(sr token.StmtReader) ([]ast.Node, error) {

	var (
		trees []ast.Node
		tree  ast.Node
		f     = Begin(sr)
		e     error
	)

	for f != nil {
		tree, f, e = f()
		if e != nil {
			return nil, e
		}
		trees = append(trees, tree)
	}

	return trees, nil
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

	first, e := lr.Read()
	if e != nil {
		return nil, e
	}

	switch first.Token {
	case token.TokenNumber:
		return beginsWithNumber(first, lr)

	default:
		return nil, newError("Unknown Token '%s'", first.Token.String())
	}
}

func beginsWithNumber(first token.Lexeme, lr token.LexemeReader) (ast.Node, error) {

	n, e := parseNumber(first)
	if e != nil {
		return nil, e
	}

	if lr.More() {
		return parseExpr(n, lr)
	}
	return n, nil
}

func parseExpr(left ast.Node, lr token.LexemeReader) (ast.Node, error) {

	return nil, nil
}

func parseNumber(num token.Lexeme) (ast.Node, error) {
	n, e := strconv.ParseInt(num.Value, 10, 64)
	if e != nil {
		return nil, e
	}
	return ast.Number{Value: n}, nil
}

func newError(msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return errors.New(msg)
}
