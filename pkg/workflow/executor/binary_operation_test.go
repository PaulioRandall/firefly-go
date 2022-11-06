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
	left := mockNumber(1)
	right := mockNumber(1)
	doBinaryOperationTest(t, left, "==", right, true)
}

func Test_exeBinaryOperation_2(t *testing.T) {
	left := mockNumber(1)
	right := mockNumber(2)
	doBinaryOperationTest(t, left, "==", right, false)
}

func Test_exeBinaryOperation_3(t *testing.T) {
	left := mockNumber(1)
	right := mockNumber(2)
	doBinaryOperationTest(t, left, "!=", right, true)
}

func Test_exeBinaryOperation_4(t *testing.T) {
	left := mockNumber(1)
	right := mockNumber(1)
	doBinaryOperationTest(t, left, "!=", right, false)
}

func Test_exeBinaryOperation_5(t *testing.T) {
	left := mockNumber(1)
	right := mockNumber(2)
	doBinaryOperationTest(t, left, "<", right, true)
}

func Test_exeBinaryOperation_6(t *testing.T) {
	left := mockNumber(2)
	right := mockNumber(1)
	doBinaryOperationTest(t, left, "<", right, false)
}

func Test_exeBinaryOperation_7(t *testing.T) {
	left := mockNumber(2)
	right := mockNumber(1)
	doBinaryOperationTest(t, left, ">", right, true)
}

func Test_exeBinaryOperation_8(t *testing.T) {
	left := mockNumber(1)
	right := mockNumber(2)
	doBinaryOperationTest(t, left, ">", right, false)
}

func Test_exeBinaryOperation_9(t *testing.T) {
	left := mockNumber(1)
	right := mockNumber(2)
	doBinaryOperationTest(t, left, "<=", right, true)
}

func Test_exeBinaryOperation_10(t *testing.T) {
	left := mockNumber(1)
	right := mockNumber(1)
	doBinaryOperationTest(t, left, "<=", right, true)
}

func Test_exeBinaryOperation_11(t *testing.T) {
	left := mockNumber(2)
	right := mockNumber(1)
	doBinaryOperationTest(t, left, "<=", right, false)
}

func Test_exeBinaryOperation_12(t *testing.T) {
	left := mockNumber(2)
	right := mockNumber(1)
	doBinaryOperationTest(t, left, ">=", right, true)
}

func Test_exeBinaryOperation_13(t *testing.T) {
	left := mockNumber(1)
	right := mockNumber(1)
	doBinaryOperationTest(t, left, ">=", right, true)
}

func Test_exeBinaryOperation_14(t *testing.T) {
	left := mockNumber(1)
	right := mockNumber(2)
	doBinaryOperationTest(t, left, ">=", right, false)
}

func Test_exeBinaryOperation_15(t *testing.T) {
	left := mockBool(true)
	right := mockBool(true)
	doBinaryOperationTest(t, left, "&&", right, true)
}

func Test_exeBinaryOperation_16(t *testing.T) {
	left := mockBool(true)
	right := mockBool(false)
	doBinaryOperationTest(t, left, "&&", right, false)
}

func Test_exeBinaryOperation_17(t *testing.T) {
	left := mockBool(false)
	right := mockBool(true)
	doBinaryOperationTest(t, left, "&&", right, false)
}

func Test_exeBinaryOperation_18(t *testing.T) {
	left := mockBool(false)
	right := mockBool(false)
	doBinaryOperationTest(t, left, "&&", right, false)
}

func Test_exeBinaryOperation_19(t *testing.T) {
	left := mockBool(true)
	right := mockBool(true)
	doBinaryOperationTest(t, left, "||", right, true)
}

func Test_exeBinaryOperation_20(t *testing.T) {
	left := mockBool(true)
	right := mockBool(false)
	doBinaryOperationTest(t, left, "||", right, true)
}

func Test_exeBinaryOperation_21(t *testing.T) {
	left := mockBool(false)
	right := mockBool(true)
	doBinaryOperationTest(t, left, "||", right, true)
}

func Test_exeBinaryOperation_22(t *testing.T) {
	left := mockBool(false)
	right := mockBool(false)
	doBinaryOperationTest(t, left, "||", right, false)
}

func Test_exeBinaryOperation_23(t *testing.T) {
	left := mockNumber(1)
	right := mockNumber(1)
	doBinaryOperationTest(t, left, "+", right, float64(2))
}

func Test_exeBinaryOperation_24(t *testing.T) {
	left := mockNumber(1)
	right := mockNumber(1)
	doBinaryOperationTest(t, left, "-", right, float64(0))
}

func Test_exeBinaryOperation_25(t *testing.T) {
	left := mockNumber(2)
	right := mockNumber(2)
	doBinaryOperationTest(t, left, "*", right, float64(4))
}

func Test_exeBinaryOperation_26(t *testing.T) {
	left := mockNumber(12)
	right := mockNumber(3)
	doBinaryOperationTest(t, left, "/", right, float64(4))
}

func Test_exeBinaryOperation_27(t *testing.T) {
	left := mockNumber(7)
	right := mockNumber(4)
	doBinaryOperationTest(t, left, "%", right, float64(3))
}

func Test_exeBinaryOperation_28(t *testing.T) {

	// (2 * 2) + (3 * 3)

	left := ast.BinaryOperation{
		Left:     mockNumber(2),
		Operator: "*",
		Right:    mockNumber(2),
	}

	right := ast.BinaryOperation{
		Left:     mockNumber(3),
		Operator: "*",
		Right:    mockNumber(3),
	}

	exp := (2 * 2) + (3 * 3)
	doBinaryOperationTest(t, left, "+", right, float64(exp))
}

func Test_exeBinaryOperation_29(t *testing.T) {

	// (1 == 2) || ((2 > 3) && (4 <= 5))

	left := ast.BinaryOperation{
		Left:     mockNumber(1),
		Operator: "==",
		Right:    mockNumber(2),
	}

	b := ast.BinaryOperation{
		Left:     mockNumber(2),
		Operator: ">",
		Right:    mockNumber(3),
	}

	c := ast.BinaryOperation{
		Left:     mockNumber(4),
		Operator: "<=",
		Right:    mockNumber(5),
	}

	right := ast.BinaryOperation{
		Left:     b,
		Operator: "&&",
		Right:    c,
	}

	exp := (1 == 2) || ((2 > 3) && (4 <= 5))
	doBinaryOperationTest(t, left, "||", right, exp)
}

func Test_exeBinaryOperation_30(t *testing.T) {

	// (1 + (2 / 3)) - ((4 * 5) % 6)

	div := ast.BinaryOperation{
		Left:     mockNumber(2),
		Operator: "/",
		Right:    mockNumber(3),
	}

	add := ast.BinaryOperation{
		Left:     mockNumber(1),
		Operator: "+",
		Right:    div,
	}

	mul := ast.BinaryOperation{
		Left:     mockNumber(4),
		Operator: "*",
		Right:    mockNumber(5),
	}

	mod := ast.BinaryOperation{
		Left:     mul,
		Operator: "%",
		Right:    mockNumber(6),
	}

	exp := (1 + (float64(2) / 3)) - ((4 * 5) % 6)

	doBinaryOperationTest(t, add, "-", mod, float64(exp))
}
