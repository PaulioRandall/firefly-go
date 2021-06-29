package parser

import (
	"github.com/PaulioRandall/firefly-go/firefly/ast"
	"github.com/PaulioRandall/firefly-go/firefly/token"
)

func expectExpr(r lexReader, leftPriority int) ast.Tree {
	expr := expectTerm(r)
	return parseInfix(r, expr, leftPriority)
}

func expectTerm(r lexReader) ast.Tree {

	tk := r.Peek().Token

	switch tk {
	case token.TK_NUMBER:
		return parseNumber(r.Read())

	case token.TK_PAREN_OPEN:
		return parseParenExpr(r, r.Read())

	case token.TK_PAREN_CLOSE:
		parsingPanic(nil, "unexpected closing parenthesis")

	default:
		parsingPanic(nil, "unexpected Token '%s'", tk.String())
	}

	return nil // Unreachable but required
}

func parseInfix(r lexReader, left ast.Tree, leftPriority int) ast.Tree {

	if !r.More() {
		return left
	}

	opToken, leftWins := leftHasPriority(r, leftPriority)
	if leftWins {
		return left
	}

	right := expectExpr(r, opToken.Precedence())
	n := buildExpr(opToken, left, right)

	return parseInfix(r, n, leftPriority)
}

func leftHasPriority(r lexReader, leftPriority int) (token.Token, bool) {

	opToken := r.Peek().Token

	if !opToken.IsOperator() {
		parsingPanic(nil, "expected operator, got '%s'", opToken.String())
	}

	if !opToken.IsCloser() && leftPriority < opToken.Precedence() {
		return r.Read().Token, false
	}

	return token.TK_UNDEFINED, true
}

func buildExpr(opToken token.Token, left, right ast.Tree) ast.Tree {

	n := mapInfixTokenToAST(opToken)
	if n == ast.NODE_UNDEFINED {
		parsingPanic(nil, "unknown operation '%s'", opToken.String())
	}

	return ast.InfixTree{
		Node:  n,
		Left:  left,
		Right: right,
	}
}

func mapInfixTokenToAST(tk token.Token) ast.Node {
	switch tk {
	case token.TK_ADD:
		return ast.NODE_ADD

	case token.TK_SUB:
		return ast.NODE_SUB

	case token.TK_MUL:
		return ast.NODE_MUL

	case token.TK_DIV:
		return ast.NODE_DIV

	default:
		return ast.NODE_UNDEFINED
	}
}
