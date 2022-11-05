package parser2

import (
	"strconv"

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

// := {TERM_STATEMENT}
func parse(r BufReaderOfTokens, w WriterOfNodes) {
	for r.More() {
		acceptEmptyStatements(r)

		n := parseTermStatement(r)
		if e := w.Write(n); e != nil {
			panic(ErrWriting.Wrap(e, "Couldn't write AST node to output"))
		}
	}
}

// := {TERM}
func acceptEmptyStatements(r BufReaderOfTokens) {
	for acceptType(r, token.Terminator) || acceptType(r, token.Newline) {
	}
}

// TERM_STATEMENT := STATEMENT TERM
func parseTermStatement(r BufReaderOfTokens) ast.Node {
	n := parseStatement(r)
	expectTerm(r)
	return n
}

// STATEMENT := ASSIGNMENT
func parseStatement(r BufReaderOfTokens) ast.Node {
	if isAssignment(r) {
		return parseAssignment(r)
	}

	panic(ErrParsing.Track("Expected statement"))
}

// := Identifier Comma
// := Identifier Assign
func isAssignment(r BufReaderOfTokens) bool {
	if peekType(r) != token.Identifier {
		return false
	}

	ident := readToken(r)
	defer r.Putback(ident)

	tt := peekType(r)
	return tt == token.Comma || tt == token.Assign
}

// ASSIGNMENT := VARIABLES Assign EXPRESSIONS
func parseAssignment(r BufReaderOfTokens) ast.Node {
	n := ast.Assign{}

	n.Dst = parseVariables(r)
	expectType(r, token.Assign)
	n.Src = parseExpressions(r)

	return n
}

// VARIABLES := [VARIABLE {Comma VARIABLE}]
func parseVariables(r BufReaderOfTokens) []ast.Variable {
	var result []ast.Variable

	if !isVariable(r) {
		return result
	}

	result = append(result, parseVariable(r))
	for acceptType(r, token.Comma) {
		result = append(result, parseVariable(r))
	}

	return result
}

// := identifier
func isVariable(r BufReaderOfTokens) bool {
	return peekType(r) == token.Identifier
}

// VARIABLE := Identifier
func parseVariable(r BufReaderOfTokens) ast.Variable {
	return ast.Variable{
		Name: expectType(r, token.Identifier).Value,
	}
}

// EXPRESSIONS := [EXPRESSION {Comma EXPRESSION}]
func parseExpressions(r BufReaderOfTokens) []ast.Expr {
	var result []ast.Expr

	if !isExpression(r) {
		return result
	}

	result = append(result, parseExpression(r))
	for acceptType(r, token.Comma) {
		result = append(result, parseExpression(r))
	}

	return result
}

// := LITERAL
func isExpression(r BufReaderOfTokens) bool {
	return isLiteral(r)
}

// EXPRESSION := LITERAL
func parseExpression(r BufReaderOfTokens) ast.Expr {
	return parseLiteral(r)
}

// := Number | String | True | False
func isLiteral(r BufReaderOfTokens) bool {
	tt := peekType(r)
	return tt == token.Number ||
		tt == token.String ||
		tt == token.True ||
		tt == token.False
}

// LITERAL := NUMBER | String | True | False
func parseLiteral(r BufReaderOfTokens) ast.Literal {
	switch peekType(r) {
	case token.Number:
		return parseNumber(r)
	case token.String:
		return parseString(r)
	default:
		panic(ErrParsing.Track("Expected literal"))
	}
}

// NUMBER := Number
func parseNumber(r BufReaderOfTokens) ast.Literal {
	v := expectType(r, token.Number).Value
	num, e := strconv.ParseFloat(v, 64)

	if e != nil {
		panic(ErrParsing.Track("Unable to parse number"))
	}

	return ast.Literal{
		Value: num,
	}
}

// STRING := String
func parseString(r BufReaderOfTokens) ast.Literal {
	str := expectType(r, token.String).Value
	str = str[1 : len(str)-1] // Slice off delimiters
	return ast.Literal{
		Value: str,
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

func acceptType(r BufReaderOfTokens, want token.TokenType) bool {
	if peekType(r) == want {
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

// TERM := Terminator | Newline
func expectTerm(r BufReaderOfTokens) {
	if acceptType(r, token.Terminator) || acceptType(r, token.Newline) {
		return
	}

	panic(ErrParsing.Track("Expected terminator or newline"))
}
