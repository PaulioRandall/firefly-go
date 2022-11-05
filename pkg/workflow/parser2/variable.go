package parser2

import (
	ast "github.com/PaulioRandall/firefly-go/pkg/models/ast2"
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

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

// == identifier
func isVariable(r BufReaderOfTokens) bool {
	return peekType(r) == token.Identifier
}

// VARIABLE := Identifier
func parseVariable(r BufReaderOfTokens) ast.Variable {
	return ast.Variable{
		Name: expectType(r, token.Identifier).Value,
	}
}
