package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// OPERATION := EXPR OPERATOR EXPR
// OPERATOR  := Add | Sub | Mul | Div | Mod | LT | GT | LTE | GTE | EQU | NEQ

func acceptOperand(a auditor) ast.Expr {
	switch {
	case !a.More():
		return nil
	case a.is(token.Identifier):
		return parseVariable(a)
	case a.match(token.IsLiteral):
		return expectLiteral(a)
	default:
		return nil
	}
}

// TERM := VAR | LITERAL
func expectOperand(a auditor) ast.Expr {
	if !a.More() {
		panic(a.unexpectedEOF("operand"))
	}

	if expr := acceptOperand(a); expr != nil {
		return expr
	}

	panic(a.unexpectedToken("operand", a.Peek()))
}

func operation(a auditor, left ast.Expr, leftOperatorPriorty int) ast.Expr {
	if !a.notMatch(token.IsBinaryOperator) {
		return left
	}

	if leftOperatorPriorty >= a.Peek().Precedence() {
		return left
	}

	op := a.Read()

	var right ast.Expr
	if a.is(token.ParenOpen) {
		right = parseParenExpr(a)
	} else {
		right = expectOperand(a)
	}

	right = operation(a, right, op.Precedence())

	left = ast.BinaryOperation{
		Left:     left,
		Operator: op,
		Right:    right,
	}

	return operation(a, left, leftOperatorPriorty)
}
