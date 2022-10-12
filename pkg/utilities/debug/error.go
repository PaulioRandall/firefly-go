package debug

import (
	"strings"

	"github.com/PaulioRandall/firefly-go/pkg/models/err"
)

func wrappedError(e error) string {
	sb := strings.Builder{}

	if e == nil {
		sb.WriteString("No error")
	} else {
		addErr(&sb, e)
	}

	return indentLines(sb.String(), 1)
}

func addErr(sb *strings.Builder, e error) {
	if next := err.Unwrap(e); next != nil {
		addErr(sb, next)
		sb.WriteRune('\n')
	}

	sb.WriteString(e.Error())
}
