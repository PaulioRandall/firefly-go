package rinser

import (
	"errors"
	"fmt"

	"github.com/PaulioRandall/firefly-go/workflow/token"
)

var (
	EOF             = errors.New("Unexpected end of file")
	UnexpectedToken = errors.New("Unexpected token")
)

type rinseErr struct {
	tk    token.Token
	after bool
	cause error
	msg   string
}

func errAt(tk token.Token, cause error, msg string, args ...interface{}) *rinseErr {
	return &rinseErr{
		tk:    tk,
		cause: cause,
		msg:   fmt.Sprintf(msg, args...),
	}
}

func errAfter(tk token.Token, cause error, msg string, args ...interface{}) *rinseErr {
	return &rinseErr{
		tk:    tk,
		after: true,
		cause: cause,
		msg:   fmt.Sprintf(msg, args...),
	}
}

func (e rinseErr) Error() string {
	return e.msg
}

func (e rinseErr) Unwrap() error {
	return e.cause
}

func (e rinseErr) Token() token.Token {
	return e.tk
}

func (e rinseErr) After() bool {
	return e.after
}
