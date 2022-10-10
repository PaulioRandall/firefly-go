// Package parser parses a series of tokens into series of abstract syntax trees
package parser

import (
	"errors"

	"github.com/PaulioRandall/firefly-go/pkg/models/ast"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

type ASTWriter = inout.Writer[ast.Node]

func Parse(r TokenReader, w ASTWriter) (e error) {
	a := newAuditor(r)

	defer func() {
		if v := recover(); v != nil {
			e = v.(error) // TODO: Wrap error
		}
	}()

	for r.More() {
		n := parseNext(a)
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
		n = parseStartingWithVariable(a, a.prev)

	case a.isNext(token.If):
		n = parseIf(a)

	default:
		panic(UnexpectedToken)
	}

	if n == nil {
		panic(errors.New("Sanity check! Nil Node should never appear"))
	}

	a.expect(token.Terminator)
	return n
}

func parseStartingWithVariable(a *auditor, first token.Token) ast.Node {
	if a.isNext(token.Comma) || a.isNext(token.Assign) {
		a.putback(first)
		return parseAssignment(a)
	}

	panic(UnexpectedToken)
}

func parseVariable(a *auditor) ast.Variable {
	return ast.Variable{
		Token: a.expect(token.Var),
	}
}

func parseVariables(a *auditor) []ast.Variable {
	var nodes []ast.Variable

	v := parseVariable(a)
	nodes = append(nodes, v)

	for a.accept(token.Comma) {
		v := parseVariable(a)
		nodes = append(nodes, v)
	}

	return nodes
}

func parseExpression(a *auditor) ast.Expr {
	return ast.Literal{
		Token: a.expectIf(token.IsLiteral, "literal"),
	}
}

func parseExpressions(a *auditor) []ast.Expr {
	var nodes []ast.Expr

	v := parseExpression(a)
	nodes = append(nodes, v)

	for a.accept(token.Comma) {
		v := parseExpression(a)
		nodes = append(nodes, v)
	}

	return nodes
}

func parseAssignment(a *auditor) ast.Assign {
	n := ast.Assign{}

	n.Left = parseVariables(a)
	n.Token = a.expect(token.Assign)
	n.Right = parseExpressions(a)

	// TODO: Move specific parameter checks to the validator
	if len(n.Left) > len(n.Right) {
		panic(MissingExpr)
	} else if len(n.Left) < len(n.Right) {
		panic(MissingVar)
	}

	return n
}

func parseIf(a *auditor) ast.Stmt {
	return nil
}
