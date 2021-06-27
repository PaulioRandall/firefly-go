// Package runner performs the instructions for ASTs (Abstract Syntax Trees).
package runner

import (
	"io"
	"os"

	"github.com/PaulioRandall/firefly-go/firefly/ast"
)

// Interpreter interface documents the functionality provided by interpreters.
type Interpreter interface {

	// SetStdout does what it says. Implementations use a default io.Writer when
	// one is not explicitly set.
	SetStdout(io.Writer)

	// SetStderr does what it says. Implementations use a default io.Writer when
	// one is not explicitly set.
	SetStderr(io.Writer)

	// ExeErr returns the currently set error. To make the interpreter as
	// flexible as possible with regards to concurrency, errors are set in a
	// field and returned on request rather than returned by the Exe function.
	ExeErr() error

	// Exe traverses and executes the AST block. If an error occurs then an error
	// field it updated and the execution finishs. The error can be retrieved
	// using the ExeErr function. When Exe exits the ExeErr must be checked to
	// determine if the end of the AST block was reached or an error encountered.
	Exe()
}

type interpreter struct {
	program ast.Block
	counter int
	stdout  io.Writer
	stderr  io.Writer
	exeErr  error
}

// NewInterpreter create a new interpreter that uses os.Stdout and os.Stderr
// as default output writers. These can be changed with the SetStdout and
// SetStderr receiver functions.
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

		tr := in.program[in.counter]
		in.exeTree(tr)

		if in.exeErr != nil {
			return
		}
	}
}

func (in *interpreter) exeTree(tr ast.Tree) {
	if tr.Type() == ast.NODE_EMPTY {
		in.println("")
		return
	}

	result, e := computeTree(tr)
	if e != nil {
		in.exeErr = e
		in.printlnErr(in.exeErr.Error())
		return
	}

	in.println(result.String())
}
