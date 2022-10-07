// Package parser parses a series of tokens into series of abstract syntax trees
package parser

import (
	"errors"

	"github.com/PaulioRandall/firefly-go/workflow/ast"
	"github.com/PaulioRandall/firefly-go/workflow/token"
	"github.com/PaulioRandall/firefly-go/workflow/tokenreader"
)

func Parse(tr tokenreader.TokenReader) ([]ast.Node, error) {
	var (
		nodes []ast.Node
		n     ast.Node
		e     error
	)

	for tr.More() {
		tk := tr.Peek()

		switch {
		case token.IsLiteral(tk.TokenType):
			n, e = expectLiteral(tr)
		default:
			return nil, errors.New("Unexpected token") // TODO: Make proper error
		}

		if e != nil {
			return nil, e // TODO: Wrap error
		}

		nodes = append(nodes, n)

		if e = expectTerminator(tr); e != nil {
			return nil, errors.New("Expected terminator") // TODO: Make proper error
		}
	}

	return nodes, nil
}

func acceptLiteral(tr tokenreader.TokenReader) (ast.Literal, bool) {
	zero := ast.Literal{}

	if !tr.More() {
		return zero, false
	}

	if token.IsLiteral(tr.Peek().TokenType) {
		n := ast.Literal{Token: tr.Read()}
		return n, true
	}

	return zero, false
}

func expectLiteral(tr tokenreader.TokenReader) (ast.Literal, error) {
	zero := ast.Literal{}

	if !tr.More() {
		return zero, errors.New("Unexpected end of file") // TODO: Make proper error
	}

	tk := tr.Read()
	if !token.IsLiteral(tk.TokenType) {
		return zero, errors.New("Expected literal") // TODO: Make proper error
	}

	n := ast.Literal{Token: tk}
	return n, nil
}

func acceptTerminator(tr tokenreader.TokenReader) bool {
	if tr.More() && tr.Peek().TokenType == token.Terminator {
		tr.Read()
		return true
	}
	return false
}

func expectTerminator(tr tokenreader.TokenReader) error {
	if !tr.More() {
		return errors.New("Expected terminator") // TODO: Make proper error
	}

	if tk := tr.Read(); tk.TokenType != token.Terminator {
		return errors.New("Expected terminator") // TODO: Make proper error
	}

	return nil
}
