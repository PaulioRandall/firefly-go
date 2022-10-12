package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/auditor"
)

func notEndOfBlock(a *auditor.Auditor) bool {
	return a.More() && !a.IsNext(token.End)
}
