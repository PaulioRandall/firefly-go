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
		return beginsWithNumber(lr, first)

	default:
		return nil, newError("Unknown Token '%s'", first.Token.String())
	}
}

func beginsWithNumber(lr token.LexemeReader, first token.Lexeme) (ast.Node, error) {
	node, e := parseNumber(first)
	if e != nil {
		return nil, e
	}
	return parseComplexExpr(lr, node)
}

func parseComplexExpr(lr token.LexemeReader, left ast.Node) (ast.Node, error) {

	if !lr.More() {
		return left, nil
	}

	node, e := parseExpr(lr, left, 0)
	if e != nil {
		return nil, e
	}

	for lr.More() {
		node, e = parseExpr(lr, node, 0)
		if e != nil {
			return nil, e
		}
	}

	return node, nil
}

func parseExpr(lr token.LexemeReader, left ast.Node, leftPriority int) (ast.Node, error) {

	op, e := lr.Read()
	if e != nil {
		return nil, e
	}

	if leftPriority >= op.Precedence() {
		lr.PutBack(op)
		return left, nil
	}

	right, e := parseExprRight(lr, op.Precedence())
	if e != nil {
		return nil, e
	}

	if !lr.More() {
		return buildExpr(op, left, right)
	}

	right, e = parseExpr(lr, right, op.Precedence())
	if e != nil {
		return nil, e
	}

	return buildExpr(op, left, right)
}

func parseExprRight(lr token.LexemeReader, leftPriority int) (ast.Node, error) {
	n, e := expectNumber(lr)
	if e != nil {
		return nil, e
	}

	if !lr.More() {
		return n, nil
	}

	return parseExpr(lr, n, leftPriority)
}

func buildExpr(op token.Lexeme, left, right ast.Node) (ast.Node, error) {
	opNode := ast.InfixOperation{
		Left:  left,
		Right: right,
	}

	switch op.Token {
	case token.TokenAdd:
		return ast.Add{InfixOperation: opNode}, nil

	case token.TokenSub:
		return ast.Sub{InfixOperation: opNode}, nil

	case token.TokenMul:
		return ast.Mul{InfixOperation: opNode}, nil

	case token.TokenDiv:
		return ast.Div{InfixOperation: opNode}, nil

	default:
		return nil, newError("Unknown operation '%s'", op.Token.String())
	}
}

func expectNumber(lr token.LexemeReader) (ast.Node, error) {
	lx, e := lr.Read()
	if e != nil {
		return nil, e
	}
	if lx.Token != token.TokenNumber {
		return nil, newError("Expected number, got '%s'", lx.Token.String())
	}
	return parseNumber(lx)
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
