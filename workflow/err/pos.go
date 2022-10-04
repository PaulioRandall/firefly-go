package err

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/workflow/pos"
)

type PosErr struct {
	pos   pos.Pos
	cause error
	msg   string
}

func AtPos(pos pos.Pos, cause error, msg string, args ...interface{}) *PosErr {
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

func (e PosErr) Pos() pos.Pos {
	return e.pos
}
