package token

type Pos struct {
	Offset int
	Line   int
	Col    int // Index on line
}

type Range struct {
	Start Pos
	End   Pos // Exclusive
}

func MakePos(offset, line, col int) Pos {
	return Pos{
		Offset: offset,
		Line:   line,
		Col:    col,
	}
}

func MakeRange(start, end Pos) Range {
	return Range{
		Start: start,
		End:   end,
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
