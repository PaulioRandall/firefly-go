package ast

// Node is the type of an AST.
type Node int

const (
	NODE_UNDEFINED Node = iota
	NODE_EMPTY
	NODE_NUM
	NODE_ADD
	NODE_SUB
	NODE_MUL
	NODE_DIV
)

var nodeNames = map[Node]string{
	NODE_EMPTY: "EMPTY",
	NODE_NUM:   "NUMBER",
	NODE_ADD:   "ADD",
	NODE_SUB:   "SUBTRACT",
	NODE_MUL:   "MULTIPLY",
	NODE_DIV:   "DIVIDE",
}

// String returns the string representation of the Node.
func (n Node) String() string {
	return nodeNames[n]
}
