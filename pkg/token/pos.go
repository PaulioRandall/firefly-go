package token

type Pos int
type Span struct {
	Start     Pos
	End       Pos // Exclusive
	Line      Pos
	Len       Pos
	LineStart Pos
	LineEnd   Pos // Exclusive
}

func MakeSpan(start, end, line, lineStart Pos) Span {
	length := end - start

	return Span{
		Start:     start,
		End:       end,
		Len:       length,
		Line:      line,
		LineStart: lineStart,
		LineEnd:   lineStart + length,
	}
}
