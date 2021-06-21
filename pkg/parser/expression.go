package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/ast"
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

func expectExpr(lr token.LexemeReader, leftPriority int) (ast.Node, error) {

	var expr ast.Node
	var e error

	lx, e := lr.Read()
	if e != nil {
		return nil, e
	}

	switch lx.Token {
	case token.TokenNumber:
		expr, e = parseNumber(lx)

	case token.TokenParenOpen:
		expr, e = parseParenExpr(lr, lx)

	case token.TokenParenClose:
		e = newError("Unexpected closing parenthesis")

	default:
		e = newError("Unexpected Token '%s'", lx.Token.String())
	}

	if e != nil {
		return nil, e
	}

	return parseInfixExpr(lr, expr, leftPriority)
}

func parseInfixExpr(lr token.LexemeReader, left ast.Node, leftPriority int) (ast.Node, error) {

	if !lr.More() {
		return left, nil
	}

	op, e := lr.Read()
	if e != nil {
		return nil, e
	}

	if op.IsCloser() || leftPriority >= op.Precedence() {
		e = lr.PutBack(op)
		if e != nil {
			return nil, e
		}
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

	return parseInfixExpr(lr, n, leftPriority)
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
