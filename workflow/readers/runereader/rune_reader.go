package runereader

import (
	"github.com/PaulioRandall/firefly-go/workflow/token"
)

type RuneReader interface {
	Pos() token.Pos
	More() bool
	Peek() (rune, error)
	Read() (rune, error)
}
