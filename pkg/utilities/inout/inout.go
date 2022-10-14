package inout

import (
	"io"

	"github.com/PaulioRandall/firefly-go/pkg/models/pos"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var EOF = err.Wrap(io.EOF, "EOF (inout)")

type Reader[T any] interface {
	More() bool
	Peek() (T, error)
	Read() (T, error)
}

type ReaderOfRunes interface {
	Reader[rune]
	Where() pos.Pos
}

type Writer[T any] interface {
	Write(T) error
	WriteMany(...T) error
}

type ListWriter[T any] interface {
	Writer[T]
	List() []T
}
