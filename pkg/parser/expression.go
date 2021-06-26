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
		parsingPanic(nil, "Unexpected closing parenthesis")

	default:
		parsingPanic(nil, "Unexpected Token '%s'", lx.Token.String())
	}

	return nil // Unreachable but required
}

func parseInfix(r lexReader, left ast.Node, leftPriority int) ast.Node {

	if !r.More() {
		return left
	}

	op, leftWins := leftHasPriority(r, leftPriority)
	if leftWins {
		return left
	}

	right := expectExpr(r, op.Precedence())
	n := buildExpr(op, left, right)

	return parseInfix(r, n, leftPriority)
}

func leftHasPriority(r lexReader, leftPriority int) (token.Lexeme, bool) {

	op := r.Read()
	if !op.Token.IsOperator() {
		parsingPanic(nil, "Expected operator, got '%s'", op.Token.String())
	}

	if !op.IsCloser() && leftPriority < op.Precedence() {
		return op, false
	}

	r.PutBack(op)
	return token.Lexeme{}, true
}

func buildExpr(op token.Lexeme, left, right ast.Node) ast.Node {

	astType := mapInfixTokenToAST(op.Token)
	if astType == ast.AstUndefined {
		parsingPanic(nil, "Unknown operation '%s'", op.Token.String())
	}

	return ast.InfixNode{
		AST:   astType,
		Left:  left,
		Right: right,
	}
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
