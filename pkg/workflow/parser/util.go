package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func badNextToken(a auditor, parsing string, parsingToken string) error {
	var cause error

	if !a.More() {
		cause = UnexpectedEOF.Trackf("Expected %s but got EOF", parsingToken)
	} else {
		cause = UnexpectedToken.Trackf(
			"Expected %s but got %s",
			parsingToken,
			a.Peek().String(),
		)
	}

	return MissingIdentifier.Wrapf(cause, "Unable to parse %s", parsing)
}

// ******* OLD ********

func isNotEndOfBlock(a auditor) bool {
	return a.More() && a.isNot(token.End)
}

func expectEndOfBlock(a auditor) token.Token {
	return a.expect(token.End)
}

func expectEndOfStmt(a auditor) {
	if a.is(token.Terminator) || a.is(token.Newline) {
		a.Read()
	} else {
		panic(a.unexpectedToken("Terminator or newline", a.Peek()))
	}
}
