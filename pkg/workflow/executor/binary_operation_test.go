package executor

import (
	"testing"

	"github.com/stretchr/testify/require"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
)

func doBinaryOperationTest(
	t *testing.T,
	left ast.Expr,
	operator string,
	right ast.Expr,
	exp any,
) {
	given := ast.BinaryOperation{
		Left:     left,
		Operator: operator,
		Right:    right,
	}

	act := exeBinaryOperation(newState(), given)

	require.Equal(t, exp, act)
}

func Test_exeBinaryOperation_1(t *testing.T) {
	left := litNumber(1)
	right := litNumber(1)
	doBinaryOperationTest(t, left, "==", right, true)
}

func Test_exeBinaryOperation_2(t *testing.T) {
	left := litNumber(1)
	right := litNumber(2)
	doBinaryOperationTest(t, left, "==", right, false)
}

func Test_exeBinaryOperation_3(t *testing.T) {
	left := litNumber(1)
	right := litNumber(2)
	doBinaryOperationTest(t, left, "!=", right, true)
}

func Test_exeBinaryOperation_4(t *testing.T) {
	left := litNumber(1)
	right := litNumber(1)
	doBinaryOperationTest(t, left, "!=", right, false)
}
