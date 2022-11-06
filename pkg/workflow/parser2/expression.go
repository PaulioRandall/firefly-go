package parser2

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

// EXPRESSIONS := [EXPRESSION {Comma EXPRESSION}]
func parseExpressions(r BufReaderOfTokens) []ast.Expr {
	var result []ast.Expr

	if !isExpression(r) {
		return result
	}

	result = append(result, parseExpression(r))
	for acceptType(r, token.Comma) {
		result = append(result, parseExpression(r))
	}

	return result
}

// EXPRESSION == LITERAL
func isExpression(r BufReaderOfTokens) bool {
	return isTerm(r)
}

// EXPRESSION := OPERATION
func parseExpression(r BufReaderOfTokens) ast.Expr {
	switch {
	case isTerm(r):
		return parseOperation(r, nil, 0)
	default:
		panic(ErrParsing.Track("Expected expression"))
	}
}

// TERM == VARIABLE | LITERAL
func isTerm(r BufReaderOfTokens) bool {
	return isVariable(r) || isLiteral(r)
}

// TERM := VARIABLE | LITERAL
func parseTerm(r BufReaderOfTokens) ast.Term {
	switch {
	case isVariable(r):
		return parseVariable(r)
	case isLiteral(r):
		return parseLiteral(r)
	default:
		panic(ErrParsing.Track("Expected term"))
	}
}

// OPERATION := EXPRESSION {OPERATOR EXPRESSION}
func parseOperation(
	r BufReaderOfTokens,
	left ast.Expr,
	priorPrecedence int,
) ast.Expr {

	if left == nil {
		left = parseTerm(r)
	}

	if !isBinaryOperator(r) {
		return left
	}

	if priorPrecedence >= peekType(r).Precedence() {
		return left
	}

	operator := readToken(r)
	right := parseOperation(r, nil, operator.Precedence())

	left = ast.BinaryOperation{
		Left:     left,
		Operator: operator.Value,
		Right:    right,
	}

	return parseOperation(r, left, priorPrecedence)
}

// BINARY_OPERATOR == ARITHMETIC_OPERATOR | COMPARISON_OPERATOR
func isBinaryOperator(r BufReaderOfTokens) bool {
	return isArithmeticOperator(r) ||
		isComparisonOperator(r) ||
		isBooleanOperator(r)
}

// ARITHMETIC_OPERATOR == Add | Sub | Mul | Div | Mod
func isArithmeticOperator(r BufReaderOfTokens) bool {
	return isAnyOfType(r,
		token.Add,
		token.Sub,
		token.Mul,
		token.Div,
		token.Mod,
	)
}

// COMPARISON_OPERATOR == LT | GT | LTE | GTE | EQU | NEQ
func isComparisonOperator(r BufReaderOfTokens) bool {
	return isAnyOfType(r,
		token.Lt,
		token.Mt,
		token.Lte,
		token.Mte,
		token.Equ,
		token.Neq,
	)
}

// BOOLEAN_OPERATOR == And | Or
func isBooleanOperator(r BufReaderOfTokens) bool {
	return isAnyOfType(r,
		token.And,
		token.Or,
	)
}
