package parser

import (
	"errors"
)

var (
	MissingVar      = errors.New("Missing variable")
	MissingExpr     = errors.New("Missing expression")
	UnexpectedToken = errors.New("Unexpected token")
)
