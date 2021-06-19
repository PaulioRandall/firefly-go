package runner

import (
	"io"

	"github.com/PaulioRandall/firefly-go/pkg/ast"
)

type Interpreter interface {
	SetStdout(io.Writer)
	SetStderr(io.Writer)
	ExeErr() error
	Start()
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

func (in *interpreter) ExeErr() error {
	return in.exeErr
}

func (in *interpreter) Start() {
	// TODO
}
