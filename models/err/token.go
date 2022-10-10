package err

import (
	"fmt"

	"github.com/PaulioRandall/firefly-go/models/token"
)

type tokErr struct {
	tk    token.Token
	after bool
	cause error
	msg   string
}

func AtToken(tk token.Token, cause error, msg string, args ...interface{}) *tokErr {
	return &tokErr{
		tk:    tk,
		cause: cause,
		msg:   fmt.Sprintf(msg, args...),
	}
}

func AfterToken(tk token.Token, cause error, msg string, args ...interface{}) *tokErr {
	return &tokErr{
		after: true,
		tk:    tk,
		cause: cause,
		msg:   fmt.Sprintf(msg, args...),
	}
}

func (e tokErr) Error() string {
	return e.msg
}

func (e tokErr) Unwrap() error {
	return e.cause
}

func (e tokErr) Token() token.Token {
	return e.tk
}

func (e tokErr) After() bool {
	return e.after
}
