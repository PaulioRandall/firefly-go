package runner

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/PaulioRandall/firefly-go/pkg/ast"
)

type Interpreter interface {
	SetStdout(io.Writer)
	SetStderr(io.Writer)
	Counter() int
	ExeErr() error
	Exe()
}

type interpreter struct {
	program ast.Program
	counter int
	stdout  io.Writer
	stderr  io.Writer
	exeErr  error
}

func NewInterpreter(p ast.Program) *interpreter {
	return &interpreter{
		program: p,
		stdout:  os.Stdout,
		stderr:  os.Stderr,
	}
}

func (in *interpreter) SetStdout(w io.Writer) {
	if w == nil {
		panic("Nil stdout")
	}
	in.stdout = w
}

func (in *interpreter) SetStderr(w io.Writer) {
	if w == nil {
		panic("Nil stderr")
	}
	in.stderr = w
}

func (in *interpreter) Counter() int {
	return in.counter
}

func (in *interpreter) ExeErr() error {
	return in.exeErr
}

func (in *interpreter) Exe() {
	if in.exeErr == nil {
		in.continueExe()
	}
}

func (in *interpreter) continueExe() {
	for ; in.counter < len(in.program); in.counter++ {
		n := in.program[in.counter]
		in.exeAstNode(n)
		if in.exeErr != nil {
			return
		}
	}
}

func (in *interpreter) exeAstNode(n ast.Node) {
	switch n.Type() {
	case ast.AstNumber:
		in.exeAstNumber(n)

	default:
		in.bug("Unknown AST node")
	}
}

func (in *interpreter) exeAstNumber(n ast.Node) {
	num, ok := n.(ast.Number)
	if !ok {
		in.bug("ast.Number node expected")
		return
	}
	in.stdPrint(num.String())
}

func (in *interpreter) stdPrint(s string) {
	_, e := fmt.Fprint(in.stdout, s)
	if e != nil {
		panic(e)
	}
}

func (in *interpreter) errPrint(s string) {
	_, e := fmt.Fprintf(in.stderr, s)
	if e != nil {
		panic(e)
	}
}

func (in *interpreter) bug(msg string, args ...interface{}) {
	msg = "[BUG] " + msg
	in.exeErr = newError(msg, args...)
	in.errPrint(in.exeErr.Error())
}

func newError(msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return errors.New(msg)
}
