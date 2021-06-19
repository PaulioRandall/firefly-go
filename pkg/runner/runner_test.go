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

func infix(t ast.AST, left, right ast.Node) ast.InfixExprNode {
	return ast.InfixExprNode{
		AST:   t,
		Left:  left,
		Right: right,
	}
}

func setupInterpreter(p ast.Program) (Interpreter, *mockWriter, *mockWriter) {
	in := NewInterpreter(p)

	std := &mockWriter{}
	in.SetStdout(std)

	err := &mockWriter{}
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

	// AND only the number is written to stdout
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

	// AND only the number is written to stdout
	exp := []byte("1\n2\n3\n")
	require.Equal(t, exp, stdout.output)
}
