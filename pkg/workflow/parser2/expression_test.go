package parser2

import (
	"testing"

	"github.com/stretchr/testify/require"

	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/debug"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"

	"github.com/PaulioRandall/firefly-go/pkg/models/token/tokentest"
)

func doExpressionTest(t *testing.T, given []token.Token, exp ast.Node) {
	defer func() {
		if e := recover(); e != nil {
			require.Fail(t, debug.String(e))
		}
	}()

	r := inout.NewListReader(given)
	br := inout.NewBufReader[token.Token](r)
	act := parseExpression(br)

	require.Equal(t, exp, act, debug.String(act))
}

func doSimpleBinaryOperationTest(
	t *testing.T,
	operator token.TokenType,
	operatorSymbol string,
) {

	defer func() {
		if e := recover(); e != nil {
			debug.Println(e)
			panic(e)
		}
	}()

	gen := tokentest.NewTokenGenerator()
	given := []token.Token{
		gen(token.Number, "1"), // Actual literals don't matter
		gen(operator, operatorSymbol),
		gen(token.Bool, "true"), // Actual literals don't matter
	}

	exp := ast.BinaryOperation{
		Left:     ast.Literal{Value: float64(1)},
		Operator: operatorSymbol,
		Right:    ast.Literal{Value: true},
	}

	r := inout.NewListReader(given)
	br := inout.NewBufReader[token.Token](r)
	act := parseExpression(br)

	require.Equal(t,
		exp, act,
		"Expected ast.BinaryOperation: 1 %s true", operatorSymbol,
	)
}

func Test_parseExpression_1(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// true
	given := []token.Token{
		gen(token.Bool, "true"),
	}

	exp := ast.Literal{
		Value: true,
	}

	doExpressionTest(t, given, exp)
}

func Test_parseExpression_2(t *testing.T) {
	gen := tokentest.NewTokenGenerator()

	// x
	given := []token.Token{
		gen(token.Ident, "x"),
	}

	exp := ast.Variable{
		Name: "x",
	}

	doExpressionTest(t, given, exp)
}

func Test_parseExpression_3(t *testing.T) {
	doSimpleBinaryOperationTest(t, token.Add, "+")
}

func Test_parseExpression_4(t *testing.T) {
	doSimpleBinaryOperationTest(t, token.Sub, "-")
}

func Test_parseExpression_5(t *testing.T) {
	doSimpleBinaryOperationTest(t, token.Mul, "*")
}

func Test_parseExpression_6(t *testing.T) {
	doSimpleBinaryOperationTest(t, token.Div, "/")
}

func Test_parseExpression_7(t *testing.T) {
	doSimpleBinaryOperationTest(t, token.Mod, "%")
}

func Test_parseExpression_8(t *testing.T) {
	doSimpleBinaryOperationTest(t, token.Lt, "<")
}

func Test_parseExpression_9(t *testing.T) {
	doSimpleBinaryOperationTest(t, token.Mt, ">")
}

func Test_parseExpression_10(t *testing.T) {
	doSimpleBinaryOperationTest(t, token.Lte, "<=")
}

func Test_parseExpression_11(t *testing.T) {
	doSimpleBinaryOperationTest(t, token.Mte, ">=")
}

func Test_parseExpression_12(t *testing.T) {
	doSimpleBinaryOperationTest(t, token.Equ, "==")
}

func Test_parseExpression_13(t *testing.T) {
	doSimpleBinaryOperationTest(t, token.Neq, "!=")
}

func Test_parseExpression_14(t *testing.T) {
	doSimpleBinaryOperationTest(t, token.And, "&&")
}

func Test_parseExpression_15(t *testing.T) {
	doSimpleBinaryOperationTest(t, token.Or, "||")
}
