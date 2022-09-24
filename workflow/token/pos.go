package token

import (
	"fmt"
)

type Pos struct {
	Offset int
	Line   int // index
	Col    int // index
}

func MakePos(offset, line, col int) Pos {
	return Pos{
		Offset: offset,
		Line:   line,
		Col:    col,
	}
}

func (p *Pos) IncRune(ru rune) {
	p.Offset++

	if ru == '\n' {
		p.Line++
		p.Col = 0
	} else {
		p.Col++
	}
}

func (p *Pos) IncString(s string) {
	size := len(s)
	p.Offset += size

	if s == "\n" {
		p.Line++
		p.Col = 0
	} else {
		p.Col += size
	}
}

func (p Pos) String() string {
	return fmt.Sprintf("Offset=%d Line=%d Col=%d", p.Offset, p.Line, p.Col)
}
