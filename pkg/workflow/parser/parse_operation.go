package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// TERM := VAR | LITERAL
func acceptOperand(a auditor) (ast.Expr, bool) {
	if n, ok := acceptParenExpr(a); ok {
		return n, true
	}

	if n, ok := acceptTerm(a); ok {
		return n, true
	}

	return nil, false
}

func expectOperand(a auditor) ast.Expr {
	if n, ok := acceptOperand(a); ok {
		return n
	}

	panic(unableToParse(a, MissingExpr, "operand"))
}

// OPERATION := EXPR OPERATOR EXPR
func operation(a auditor, left ast.Expr, leftOperatorPriorty int) ast.Expr {
	if !a.notMatch(token.IsBinaryOperator) {
		return left
	}

	if leftOperatorPriorty >= a.Peek().Precedence() {
		return left
	}

	op := expectOperator(a)

	right := expectOperand(a)
	right = operation(a, right, op.Precedence())

	left = ast.BinaryOperation{
		Left:     left,
		Operator: op,
		Right:    right,
	}

	return operation(a, left, leftOperatorPriorty)
}

// OPERATOR := Add | Sub | Mul | Div | Mod | LT | GT | LTE | GTE | EQU | NEQ
func expectOperator(a auditor) token.Token {
	if a.isAny(
		token.Add, token.Sub, token.Mul, token.Div, token.Mod,
		token.LT, token.GT,
		token.LTE, token.GTE,
		token.EQU, token.NEQ,
	) {
		return a.Read()
	}

	panic(unableToParse(a, MissingExpr, "any in [Add | Sub | Mul | Div | Mod | LT | GT | LTE | GTE | EQU | NEQ]"))
}
