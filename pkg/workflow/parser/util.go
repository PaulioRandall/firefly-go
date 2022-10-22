package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/token"
)

func isNotEndOfBlock(a auditor) bool {
	return a.More() && a.isNot(token.End)
}
