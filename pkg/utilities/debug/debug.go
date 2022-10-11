package debug

import (
	"errors"
	"fmt"
	"strings"
)

func Print(v any) {
	fmt.Print(String(v))
}

func Println(v any) {
	fmt.Println(String(v))
}

func String(v any) string {
	switch t := v.(type) {
	case error:
		Error(t)
	}

	return fmt.Sprintf("%+v", v)
}

func Error(e error) string {
	sb := strings.Builder{}

	if e == nil {
		sb.WriteString("No error")
	} else {
		addErr(&sb, e)
	}

	return indentLines(sb.String())
}

func addErr(sb *strings.Builder, e error) {
	if next := errors.Unwrap(e); next != nil {
		addErr(sb, next)
		sb.WriteRune('\n')
	}

	sb.WriteString(e.Error())
}

func indentLines(s string) string {
	s = strings.ReplaceAll(s, "\n", "\n\t")
	return strings.TrimSpace(s)
}
