package parser

import (
	"errors"
)

var (
	UnexpectedToken = errors.New("Unexpected token")
	MissingVar      = errors.New("Missing variable")
	MissingExpr     = errors.New("Missing expression")
	MissingStmt     = errors.New("Missing statement")
	MissingEnd      = errors.New("Missing end")
)