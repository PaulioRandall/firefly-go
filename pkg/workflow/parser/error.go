package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	UnexpectedEOF   = err.Trackable("Unexpected end of file")
	UnexpectedToken = err.Trackable("Unexpected token")
)

func wrapPanic(thunk func(error) error) {
	if e := recover(); e != nil {
		panic(thunk(e.(error)))
	}
}

func badNextToken(a auditor, parsing string) error {
	if !a.More() {
		return UnexpectedEOF.Trackf("Expected %s but got EOF", parsing)
	}

	return UnexpectedToken.Trackf(
		"Expected %s but got %s",
		parsing,
		a.Peek().String(),
	)
}

func unableToParse(a auditor, te err.TrackableError, parsing string) error {
	e := badNextToken(a, parsing)
	return te.Wrapf(e, "Unable to parse %s", parsing)
}
