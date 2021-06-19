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

	result, e := in.computeAstNode(n)

	if e != nil {
		in.logBug(e)
		return
	}

	in.stdPrintln(result.String())
}

func (in *interpreter) computeAstNode(n ast.Node) (ast.NumberNode, error) {

	var zero ast.NumberNode
	var result ast.NumberNode
	var e error

	switch n.Type() {
	case ast.AstNumber:
		result, e = in.computeAstNumber(n)

	case ast.AstAdd, ast.AstSub, ast.AstMul, ast.AstDiv:
		result, e = in.computeAstInfixExpr(n)

	default:
		e = in.newBug("Unknown AST node")
	}

	if e != nil {
		return zero, e
	}
	return result, nil
}

func (in *interpreter) computeAstNumber(n ast.Node) (ast.NumberNode, error) {
	num, ok := n.(ast.NumberNode)
	if !ok {
		return ast.NumberNode{}, in.newBug("ast.NumberNode node expected")
	}
	return num, nil
}

func (in *interpreter) computeAstInfixExpr(n ast.Node) (ast.NumberNode, error) {

	var zero ast.NumberNode

	ien, ok := n.(ast.InfixExprNode)
	if !ok {
		return zero, in.newBug("ast.InfixExprNode node expected")
	}

	var result ast.NumberNode
	var e error

	left, e := in.computeAstNode(ien.Left)
	if e != nil {
		return zero, e
	}

	right, e := in.computeAstNode(ien.Right)
	if e != nil {
		return zero, e
	}

	switch n.Type() {
	case ast.AstAdd:
		result.Value = left.Value + right.Value

	case ast.AstSub:
		result.Value = left.Value - right.Value

	case ast.AstMul:
		result.Value = left.Value * right.Value

	case ast.AstDiv:
		result.Value = left.Value / right.Value

	default:
		return zero, in.newBug("Unknown AST infix expresion node")
	}

	if e != nil {
		return zero, e
	}

	return result, nil
}

func (in *interpreter) stdPrint(s string) {
	_, e := fmt.Fprint(in.stdout, s)
	if e != nil {
		panic(e)
	}
}

func (in *interpreter) stdPrintln(s string) {
	_, e := fmt.Fprint(in.stdout, s+"\n")
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

func (in *interpreter) errPrintln(s string) {
	_, e := fmt.Fprintf(in.stderr, s+"\n")
	if e != nil {
		panic(e)
	}
}

func (in *interpreter) newBug(msg string, args ...interface{}) error {
	msg = "[BUG] " + msg
	return newError(msg, args...)
}

func (in *interpreter) logBug(e error) {
	in.exeErr = e
	in.errPrintln(in.exeErr.Error())
}

func newError(msg string, args ...interface{}) error {
	msg = fmt.Sprintf(msg, args...)
	return errors.New(msg)
}
