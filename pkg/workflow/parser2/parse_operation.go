package parser2

/*
import (
	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	ErrMissingOperand  = err.Trackable("Missing operand")
	ErrMissingOperator = err.Trackable("Missing operator")

	ErrBadOperation = err.Trackable("Failed to parse operation")
	ErrBadOperand   = err.Trackable("Failed to parse operand")
)

// TERM := VAR | LITERAL
func acceptOperand(a auditor) (ast.Expr, bool) {
	defer wrapPanic(func(e error) error {
		return ErrBadOperand.Wrap(e, "Bad operand syntax")
	})

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

	panic(unableToParse(a, ErrMissingOperand, "operand"))
}

// OPERATION := EXPR OPERATOR EXPR
func operation(a auditor, left ast.Expr, leftOperatorPriorty int) ast.Expr {
	defer wrapPanic(func(e error) error {
		return ErrBadOperation.Wrap(e, "Bad operation syntax")
	})

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

	panic(unableToParse(a, ErrMissingOperator, "any in [Add | Sub | Mul | Div | Mod | LT | GT | LTE | GTE | EQU | NEQ]"))
}
*/
