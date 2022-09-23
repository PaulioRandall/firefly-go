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
	return fmt.Sprintf("Offset: %d, Line: %d, Col: %d", p.Offset, p.Line, p.Col)
}

type Range struct {
	From Pos
	To   Pos // exclusive
}

func MakeRange(from, to Pos) Range {
	return Range{
		From: from,
		To:   to,
	}
}

func MakeInlineRange(offset, line, col, length int) Range {
	return Range{
		From: Pos{
			Offset: offset,
			Line:   line,
			Col:    col,
		},
		To: Pos{
			Offset: offset + length,
			Line:   line,
			Col:    col + length,
		},
	}
}

func (r *Range) IncString(s string) {
	r.From = r.To
	r.To.IncString(s)
}

func (r Range) String() string {
	return fmt.Sprintf("From(%v), To(%v)", r.From, r.To)
}
