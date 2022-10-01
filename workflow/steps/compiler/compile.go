// Package compiler parses a list of tokens into abstract syntax trees
package compiler

import (
	"errors"

	"github.com/PaulioRandall/firefly-go/workflow/ast"
	"github.com/PaulioRandall/firefly-go/workflow/readers/tokenreader"
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

func Compile(tr tokenreader.TokenReader) ([]ast.Node, error) {
	var (
		nodes []ast.Node
		n     ast.Node
		e     error
	)

	for tr.More() {
		tk := tr.Peek()

		switch {
		case tk.TokenType.IsLiteral():
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

	if tr.Peek().TokenType.IsLiteral() {
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
	if !tk.TokenType.IsLiteral() {
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
