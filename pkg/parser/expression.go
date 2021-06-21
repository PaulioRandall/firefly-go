package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/ast"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func expectExpr(lr token.LexemeReader, leftPriority int) (ast.Node, error) {

	expr, e := expectTerm(lr)

	if e != nil {
		return nil, e
	}

	return parseInfix(lr, expr, leftPriority)
}

func expectTerm(lr token.LexemeReader) (ast.Node, error) {

	lx, e := lr.Read()
	if e != nil {
		return nil, e
	}

	switch lx.Token {
	case token.TokenNumber:
		return parseNumber(lx)

	case token.TokenParenOpen:
		return parseParenExpr(lr, lx)

	case token.TokenParenClose:
		return nil, newError("Unexpected closing parenthesis")

	default:
		return nil, newError("Unexpected Token '%s'", lx.Token.String())
	}
}

func parseInfix(lr token.LexemeReader, left ast.Node, leftPriority int) (ast.Node, error) {

	if !lr.More() {
		return left, nil
	}

	op, leftAssoc, e := leftHasPriority(lr, leftPriority)
	if e != nil {
		return nil, e
	}

	if leftAssoc {
		return left, nil
	}

	right, e := expectExpr(lr, op.Precedence())
	if e != nil {
		return nil, e
	}

	n, e := buildExpr(op, left, right)
	if e != nil {
		return nil, e
	}

	return parseInfix(lr, n, leftPriority)
}

func leftHasPriority(lr token.LexemeReader, leftPriority int) (token.Lexeme, bool, error) {
	var zero token.Lexeme

	op, e := lr.Read()
	if e != nil {
		return zero, false, e
	}

	if !op.Token.IsOperator() {
		return zero, false, newError("Expected operator, got '%s'", op.Token.String())
	}

	if op.IsCloser() || leftPriority >= op.Precedence() {
		e = lr.PutBack(op)
		return zero, true, e
	}

	return op, false, nil
}

func buildExpr(op token.Lexeme, left, right ast.Node) (ast.Node, error) {

	n := ast.InfixNode{
		AST:   mapInfixTokenToAST(op.Token),
		Left:  left,
		Right: right,
	}

	if n.AST == ast.AstUndefined {
		return nil, newError("Unknown operation '%s'", op.Token.String())
	}

	return n, nil
}

func mapInfixTokenToAST(tk token.Token) ast.AST {
	switch tk {
	case token.TokenAdd:
		return ast.AstAdd

	case token.TokenSub:
		return ast.AstSub

	case token.TokenMul:
		return ast.AstMul

	case token.TokenDiv:
		return ast.AstDiv

	default:
		return ast.AstUndefined
	}
}
