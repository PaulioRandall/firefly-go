package auditor

import (
	"github.com/PaulioRandall/firefly-go/pkg/models/token"

	"github.com/PaulioRandall/firefly-go/pkg/utilities/container"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/err"
	"github.com/PaulioRandall/firefly-go/pkg/utilities/inout"
)

var zero token.Token

type ReaderOfTokens = inout.Reader[token.Token]

type Auditor struct {
	reader ReaderOfTokens
	buffer container.Queue[token.Token]
	prev   token.Token
}

func NewAuditor(r ReaderOfTokens) *Auditor {
	return &Auditor{
		reader: r,
		buffer: &container.LinkedQueue[token.Token]{},
	}
}

func (a Auditor) Prev() token.Token {
	return a.prev
}

func (a *Auditor) More() bool {
	return a.buffer.More() || a.reader.More()
}

func (a *Auditor) Peek() token.Token {
	a.loadBuffer()

	if tk, ok := a.buffer.First(); ok {
		return tk
	}

	panic(err.NewPos(a.prev.To, "Failed to peek token from buffer"))
}

func (a *Auditor) Read() token.Token {
	a.loadBuffer()

	if tk, ok := a.buffer.Take(); ok {
		a.prev = tk
		return tk
	}

	panic(err.NewPos(a.prev.To, "Failed to read token from buffer"))
}

func (a *Auditor) Putback(tk token.Token) {
	a.buffer.Return(tk)
}

func (a *Auditor) loadBuffer() {
	if a.buffer.More() {
		return
	}

	tk, e := a.reader.Read()
	if e != nil {
		panic(err.WrapPos(e, a.prev.To, "Failed to read token from reader"))
	}

	a.buffer.Add(tk)
}
