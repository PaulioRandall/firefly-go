package runner

import (
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
	program ast.Block
	counter int
	stdout  io.Writer
	stderr  io.Writer
	exeErr  error
}

func NewInterpreter(b ast.Block) *interpreter {
	return &interpreter{
		program: b,
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

		stmt := in.program[in.counter]
		in.exeStmt(stmt)

		if in.exeErr != nil {
			return
		}
	}
}

func (in *interpreter) exeStmt(n ast.Node) {
	if n.Type() == ast.AstEmpty {
		in.println("")
		return
	}

	result, e := computeNode(n)
	if e != nil {
		in.exeErr = e
		in.printlnErr(in.exeErr.Error())
		return
	}

	in.println(result.String())
}
