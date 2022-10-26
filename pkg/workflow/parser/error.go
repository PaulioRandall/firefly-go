package parser

import (
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	ErrForLoopControls = err.Trackable("Bad for loop initialiser, condition, or advancement")

	UnexpectedEOF   = err.Trackable("Unexpected end of file")
	UnexpectedToken = err.Trackable("Unexpected token")

	MissingVar  = err.Trackable("Missing variable")
	MissingStmt = err.Trackable("Missing statement")
	MissingEnd  = err.Trackable("Missing end")
)

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
