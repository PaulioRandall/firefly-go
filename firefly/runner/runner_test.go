package runner

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/PaulioRandall/firefly-go/firefly/ast"
)

type mockWriter struct {
	output []byte
}

func (w *mockWriter) Write(bytes []byte) (int, error) {
	w.output = append(w.output, bytes...)
	return len(bytes), nil
}

func num(n int64) ast.NumberTree {
	return ast.NumberTree{
		Value: n,
	}
}

func infix(n ast.Node, left, right ast.Tree) ast.InfixTree {
	return ast.InfixTree{
		Node:  n,
		Left:  left,
		Right: right,
	}
}

func setupInterpreter(p ast.Block) (Interpreter, *mockWriter, *mockWriter) {
	in := NewInterpreter(p)

	std := &mockWriter{output: []byte{}}
	in.SetStdout(std)

	err := &mockWriter{output: []byte{}}
	in.SetStderr(err)

	return in, std, err
}

func TestInterpreter_0(t *testing.T) {

	// GIVEN a program with an empty statement
	p := ast.Block{
		ast.EmptyTree{},
	}

	// AND an interpreter initialised with the program
	in, stdout, _ := setupInterpreter(p)

	// WHEN the program is executed
	in.Exe()

	// THEN no error is set
	e := in.ExeErr()
	require.Nil(t, e, "%+v", e)

	// AND only a linefeed is written to stdout
	exp := []byte("\n")
	require.Equal(t, exp, stdout.output)
}

func TestInterpreter_1(t *testing.T) {

	// GIVEN a program that prints a number
	p := ast.Block{
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
	p := ast.Block{
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
	p := ast.Block{
		// 1 + 2
		infix(ast.NODE_ADD,
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
	p := ast.Block{
		// 8 + 6 / 3 * 5 - 4 * 3
		// (8 + ((6 / 3) * 5)) - (4 * 3)
		infix(ast.NODE_SUB,
			infix(ast.NODE_ADD,
				num(8),
				infix(ast.NODE_MUL,
					infix(ast.NODE_DIV, num(6), num(3)), // =2
					num(5),
				), // =10
			), // =18
			infix(ast.NODE_MUL, num(4), num(3)), // =12
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
	p := ast.Block{
		// 1 / 0
		infix(ast.NODE_DIV, num(1), num(0)),
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
