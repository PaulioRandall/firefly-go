package ast

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/pos"
)

// Node represents an Abstract Syntax Tree (AST) or executable statement
//
// Nodes maybe nested and dependent upon others in an acyclic manner such that
// at runtime any dependees are executed first and their results used in the
// node's execution.
type Node interface {
	// Where returns the start and end positions of the node within a source file
	Where() (from, to pos.Pos)
	node()
}
