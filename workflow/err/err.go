package err

import (
	"errors"
	"fmt"

	"github.com/PaulioRandall/firefly-go/workflow/token"
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
