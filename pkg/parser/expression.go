package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/ast"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func expectExpr(r lexReader, leftPriority int) ast.Node {
	expr := expectTerm(r)
	return parseInfix(r, expr, leftPriority)
}

func expectTerm(r lexReader) ast.Node {

	lx := r.Read()

	switch lx.Token {
	case token.TokenNumber:
		return parseNumber(lx)

	case token.TokenParenOpen:
		return parseParenExpr(r, lx)

	case token.TokenParenClose:
		panicParseErr(nil, "Unexpected closing parenthesis")
	}

	panicParseErr(nil, "Unexpected Token '%s'", lx.Token.String())
	return nil
}

func parseInfix(r lexReader, left ast.Node, leftPriority int) ast.Node {

	if !r.More() {
		return left
	}

	op, leftAssoc := leftHasPriority(r, leftPriority)
	if leftAssoc {
		return left
	}

	right := expectExpr(r, op.Precedence())
	n := buildExpr(op, left, right)

	return parseInfix(r, n, leftPriority)
}

func leftHasPriority(r lexReader, leftPriority int) (token.Lexeme, bool) {

	op := r.Read()

	if !op.Token.IsOperator() {
		panicParseErr(nil, "Expected operator, got '%s'", op.Token.String())
	}

	if !op.IsCloser() && leftPriority < op.Precedence() {
		return op, false
	}

	r.PutBack(op)
	return token.Lexeme{}, true
}

func buildExpr(op token.Lexeme, left, right ast.Node) ast.Node {

	n := ast.InfixNode{
		AST:   mapInfixTokenToAST(op.Token),
		Left:  left,
		Right: right,
	}

	if n.AST == ast.AstUndefined {
		panicParseErr(nil, "Unknown operation '%s'", op.Token.String())
	}

	return n
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
