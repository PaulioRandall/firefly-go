package inout

import (
	"io"

	"github.com/PaulioRandall/firefly-go/pkg/models/pos"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
)

var (
	EOF             = err.Wrap(io.EOF, "EOF (inout)")
	ErrRead         = err.Trackable("Failed to read or peek next value")
	ErrReadDelegate = err.Trackable("Failed to read from delegate")
)

type Reader[T any] interface {
	More() bool
	Peek() (T, error)
	Read() (T, error)
}

type BufReader[T any] interface {
	Reader[T]
	Prev() T
	Putback(T)
}

type ReaderOfRunes interface {
	Reader[rune]
	Where() pos.Pos
}

type PanicReader[T any] interface {
	More() bool
	Peek() T
	Read() T
}

type PosReader[T pos.Positioned] interface {
	PanicReader[T]
	Prev() T
	Putback(T)
}

type Writer[T any] interface {
	Write(T) error
	WriteMany(...T) error
}

type ListWriter[T any] interface {
	Writer[T]
	List() []T
}
