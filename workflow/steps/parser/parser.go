// Package parser parses a series of tokens into series of abstract syntax trees
package parser

import (
	"errors"

	"github.com/PaulioRandall/firefly-go/workflow/ast"
	"github.com/PaulioRandall/firefly-go/workflow/inout"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type ASTWriter = inout.Writer[ast.Node]

var UnexpectedToken = errors.New("Unexpected token")
var MissingVar = errors.New("Missing variable")
var MissingExpr = errors.New("Missing expression")

func Parse(r TokenReader, w ASTWriter) (e error) {
	a := auditor{
		TokenReader: r,
	}

	defer func() {
		if v := recover(); v != nil {
			e = v.(error) // TODO: Wrap error
		}
	}()

	for r.More() {
		n := parseNext(&a)
		e := w.Write(n)
		if e != nil {
			panic(e)
		}
	}

	return nil
}

func parseNext(a *auditor) (n ast.Node) {
	switch {
	case a.accept(token.Var):
		n = parseStartingWithVariable(a)

	default:
		panic(UnexpectedToken)
	}

	if n == nil {
		panic(errors.New("Sanity check! Nil Node should never appear"))
	}

	a.expect(token.Terminator)
	return n
}

func parseStartingWithVariable(a *auditor) ast.Node {
	if a.isNext(token.Comma) || a.isNext(token.Assign) {
		return parseAssignment(a, true)
	}

	panic(UnexpectedToken)
}

func parseVariable(a *auditor, alreadyRead bool) ast.Variable {
	var tk token.Token

	if alreadyRead {
		tk = a.getPrev()
	} else {
		tk = a.expect(token.Var)
	}

	return ast.Variable{
		Token: tk,
	}
}

func parseVariables(a *auditor, firstAlreadyRead bool) []ast.Variable {
	var nodes []ast.Variable

	v := parseVariable(a, firstAlreadyRead)
	nodes = append(nodes, v)

	for a.accept(token.Comma) {
		v := parseVariable(a, false)
		nodes = append(nodes, v)
	}

	return nodes
}

func parseExpression(a *auditor, alreadyRead bool) ast.Expr {
	var tk token.Token

	if alreadyRead {
		tk = a.getPrev()
	} else {
		tk = a.expect(token.Number)
	}

	return ast.Literal{
		Token: tk,
	}
}

func parseExpressions(a *auditor, firstAlreadyRead bool) []ast.Expr {
	var nodes []ast.Expr

	v := parseExpression(a, firstAlreadyRead)
	nodes = append(nodes, v)

	for a.accept(token.Comma) {
		v := parseExpression(a, false)
		nodes = append(nodes, v)
	}

	return nodes
}

func parseAssignment(a *auditor, firstAlreadyRead bool) ast.Assign {
	n := ast.Assign{}

	n.Left = parseVariables(a, firstAlreadyRead)
	n.Token = a.expect(token.Assign)
	n.Right = parseExpressions(a, false)

	if len(n.Left) > len(n.Right) {
		panic(MissingExpr)
	} else if len(n.Left) < len(n.Right) {
		panic(MissingVar)
	}

	return n
}
