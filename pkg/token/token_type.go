package token

type TokenType int

const (
	Unknown TokenType = iota
	If
	For
	Watch
	When
	E
	F
	End
)

var strMap = map[TokenType]string{
	If:    "if",
	For:   "for",
	Watch: "watch",
	When:  "when",
	E:     "expression",
	F:     "function",
	End:   "end",
}

func (tt TokenType) String() string {
	return strMap[tt]
}

func (tt TokenType) IsKeyword() bool {
	switch tt {
	case If, For, Watch, When, E, F, End:
		return true
	default:
		return false
	}
}
