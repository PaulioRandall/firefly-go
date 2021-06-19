package runner

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/pkg/ast"
)

type mockWriter struct {
	output []byte
}

func (w *mockWriter) Write(bytes []byte) (int, error) {
	w.output = append(w.output, bytes...)
	return len(bytes), nil
}

func num(n int64) ast.NumberNode {
	return ast.NumberNode{
		Value: n,
	}
}

func infix(astType ast.AST, left, right ast.Node) ast.InfixExprNode {
	return ast.InfixExprNode{
		AST:   astType,
		Left:  left,
		Right: right,
	}
}

func setupInterpreter(p ast.Program) (Interpreter, *mockWriter, *mockWriter) {
	in := NewInterpreter(p)

	std := &mockWriter{output: []byte{}}
	in.SetStdout(std)

	err := &mockWriter{output: []byte{}}
	in.SetStderr(err)

	return in, std, err
}

func TestInterpreter_1(t *testing.T) {

	// GIVEN a program that prints a number
	p := ast.Program{
		num(9),
	}

	// AND an interpreter initialised with the program
	in, stdout, _ := setupInterpreter(p)

	// WHEN the program is executed
	in.Exe()

	// THEN no error is set
	e := in.ExeErr()
	require.Nil(t, e, "%+v", e)

	// AND only the number and a linefeed are written to stdout
	exp := []byte("9\n")
	require.Equal(t, exp, stdout.output)
}

func TestInterpreter_2(t *testing.T) {

	// GIVEN a program that prints numbers on multiple lines
	p := ast.Program{
		num(1),
		num(2),
		num(3),
	}

	// AND an interpreter initialised with the program
	in, stdout, _ := setupInterpreter(p)

	// WHEN the program is executed
	in.Exe()

	// THEN no error is set
	e := in.ExeErr()
	require.Nil(t, e, "%+v", e)

	// AND only the number and a linefeed are written to stdout
	exp := []byte("1\n2\n3\n")
	require.Equal(t, exp, stdout.output)
}

func TestInterpreter_3(t *testing.T) {

	// GIVEN a program with a single expression
	p := ast.Program{
		// 1 + 2
		infix(ast.AstAdd,
			num(1),
			num(2),
		),
	}

	// AND an interpreter initialised with the program
	in, stdout, _ := setupInterpreter(p)

	// WHEN the program is executed
	in.Exe()

	// THEN no error is set
	e := in.ExeErr()
	require.Nil(t, e, "%+v", e)

	// AND only the expression result and a linefeed are written to stdout
	exp := []byte("3\n")
	require.Equal(t, exp, stdout.output)
}

func TestInterpreter_4(t *testing.T) {

	// GIVEN a program with a complex expression
	p := ast.Program{
		// 8 + 6 / 3 * 5 - 4 * 3
		// (8 + ((6 / 3) * 5)) - (4 * 3)
		infix(ast.AstSub,
			infix(ast.AstAdd,
				num(8),
				infix(ast.AstMul,
					infix(ast.AstDiv, num(6), num(3)), // =2
					num(5),
				), // =10
			), // =18
			infix(ast.AstMul, num(4), num(3)), // =12
		), // =6
	}

	// AND an interpreter initialised with the program
	in, stdout, _ := setupInterpreter(p)

	// WHEN the program is executed
	in.Exe()

	// THEN no error is set
	e := in.ExeErr()
	require.Nil(t, e, "%+v", e)

	// AND only the expression result and a linefeed are written to stdout
	exp := []byte("6\n")
	require.Equal(t, exp, stdout.output)
}

func TestInterpreter_5(t *testing.T) {

	// GIVEN an expression which attempts to divide by zero
	p := ast.Program{
		// 1 + 2
		infix(ast.AstDiv,
			num(1),
			num(0),
		),
	}

	// AND an interpreter initialised with the program
	in, stdout, errout := setupInterpreter(p)

	// WHEN the program is executed
	in.Exe()

	// THEN an error is set
	e := in.ExeErr()
	require.NotNil(t, e, "Expected error")

	// AND something is written out to errout
	require.True(t, len(errout.output) > 0, "Expected error output")

	// AND nothing is written out to stdout
	require.Equal(t, []byte{}, stdout.output)
}
