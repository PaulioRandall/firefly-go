package parser2

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

type ReaderOfTokens = inout.Reader[token.Token]
type BufReaderOfTokens = inout.BufReader[token.Token]
type WriterOfNodes = inout.Writer[ast.Node]

var (
	ErrParsing = err.Trackable("Failed to parse tokens")
	ErrReading = err.Trackable("Failed to read from input")
	ErrWriting = err.Trackable("Failed to write to output")
)

func Parse(r ReaderOfTokens, w WriterOfNodes) (e error) {
	br := inout.NewBufReader[token.Token](r)

	defer func() {
		if v := recover(); v != nil {
			e = ErrParsing.Wrap(v.(error), "Recovered from parse fail")
		}
	}()

	parse(br, w)
	return nil
}

// <- {TERM_STATEMENT}
func parse(r BufReaderOfTokens, w WriterOfNodes) {
	acceptEmptyStatements(r)

	for r.More() {
		n := parseTermStatement(r)
		if e := w.Write(n); e != nil {
			panic(ErrWriting.Wrap(e, "Couldn't write AST node to output"))
		}

		acceptEmptyStatements(r)
	}
}

func peekType(r BufReaderOfTokens) token.TokenType {
	tk, e := r.Peek()
	if e != nil {
		panic(ErrReading.Wrap(e, "Couldn't peek at token type"))
	}
	return tk.TokenType
}

func readToken(r BufReaderOfTokens) token.Token {
	tk, e := r.Read()
	if e != nil {
		panic(ErrReading.Wrap(e, "Couldn't read token"))
	}
	return tk
}

func isType(r BufReaderOfTokens, want token.TokenType) bool {
	return r.More() && peekType(r) == want
}

func isAnyOfType(r BufReaderOfTokens, anyOf ...token.TokenType) bool {
	for _, want := range anyOf {
		if isType(r, want) {
			return true
		}
	}
	return false
}

func acceptType(r BufReaderOfTokens, want token.TokenType) bool {
	if isType(r, want) {
		readToken(r)
		return true
	}
	return false
}

func acceptAnyOfType(r BufReaderOfTokens, anyOf ...token.TokenType) bool {
	if isAnyOfType(r, anyOf...) {
		readToken(r)
		return true
	}
	return false
}

func expectType(r BufReaderOfTokens, want token.TokenType) token.Token {
	if acceptType(r, want) {
		return r.Prev()
	}
	panic(ErrParsing.Trackf("Expected %q", want))
}

func expectAnyOfType(r BufReaderOfTokens, anyOf ...token.TokenType) token.Token {
	if acceptAnyOfType(r, anyOf...) {
		return r.Prev()
	}
	panic(ErrParsing.Trackf("Expected any of %q", anyOf))
}
