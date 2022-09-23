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

func (p *Pos) Inc(ru rune) {
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
	Start Pos
	End   Pos // exclusive
}

func MakeRange(start, end Pos) Range {
	return Range{
		Start: start,
		End:   end,
	}
}

func MakeInlineRange(offset, line, col, length int) Range {
	return Range{
		Start: Pos{
			Offset: offset,
			Line:   line,
			Col:    col,
		},
		End: Pos{
			Offset: offset + length,
			Line:   line,
			Col:    col + length,
		},
	}
}

func (r Range) String() string {
	return fmt.Sprintf("Start(%v), End(%v)", r.Start, r.End)
}
