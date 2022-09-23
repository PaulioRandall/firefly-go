package err

import (
	"errors"
	"fmt"
	"strings"

	"github.com/PaulioRandall/firefly-go/pkg/token"
)

var EOF = errors.New("End of file (EOF)")

type PosErr struct {
	pos   token.Pos
	cause error
	msg   string
}

func Pos(pos token.Pos, cause error, msg string, args ...interface{}) *PosErr {
	return &PosErr{
		pos:   pos,
		cause: cause,
		msg:   fmt.Sprintf(msg, args...),
	}
}

func (e PosErr) Error() string {
	return e.msg
}

func (e PosErr) Unwrap() error {
	return e.cause
}

func (e PosErr) Pos() token.Pos {
	return e.pos
}

func Debug(e error) {
	fmt.Println(DebugString(e))
}

func DebugString(e error) string {
	sb := &strings.Builder{}
	sb.WriteString("[ERROR] ")

	if e == nil {
		sb.WriteString("No error")
		return sb.String()
	}

	addErr(sb, e)
	return sb.String()
}

func addErr(sb *strings.Builder, e error) {
	next := errors.Unwrap(e)

	if next != nil {
		addErr(sb, next)
		sb.WriteRune('\n')
	}

	sb.WriteString(e.Error())
}
