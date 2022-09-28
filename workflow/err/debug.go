package err

import (
	"errors"
	"fmt"
	"strings"
)

func Println(e error) {
	fmt.Println(Debug(e))
}

func Debug(e error) string {
	sb := &strings.Builder{}
	sb.WriteString("[DEBUG ERROR]\n")

	if e == nil {
		sb.WriteString("No error")
	} else {
		addErr(sb, e)
	}

	return indentEnsuingLines(sb.String())
}

func addErr(sb *strings.Builder, e error) {
	if next := errors.Unwrap(e); next != nil {
		addErr(sb, next)
		sb.WriteRune('\n')
	}

	sb.WriteString(e.Error())
}

func indentEnsuingLines(s string) string {
	s = strings.ReplaceAll(s, "\n", "\n\t")
	return strings.TrimSpace(s)
}
