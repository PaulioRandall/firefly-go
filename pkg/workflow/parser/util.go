package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func isNotEndOfBlock(a auditor) bool {
	return a.More() && a.isNot(token.End)
}

func expectEndOfBlock(a auditor) token.Token {
	return a.expect(token.End)
}

func expectEndOfStmt(a auditor) {
	a.expect(token.Terminator)
}

func expectTerminator(a auditor) {
	a.expect(token.Terminator)
}
