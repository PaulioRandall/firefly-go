package ast

import (
	"testing"
)

func TestEnforceTypes(t *testing.T) {
	var n Node

	n = EmptyNode{}
	n = NumberNode{}
	n = InfixNode{}

	_ = n
}
