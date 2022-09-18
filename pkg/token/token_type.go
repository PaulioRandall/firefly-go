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
	Var
)

var strMap = map[TokenType]string{
	If:    "if",
	For:   "for",
	Watch: "watch",
	When:  "when",
	E:     "expression",
	F:     "function",
	End:   "end",
	Var:   "variable",
}

var keywordMap = map[string]TokenType{
	"if":    If,
	"for":   For,
	"watch": Watch,
	"when":  When,
	"E":     E,
	"F":     F,
	"end":   End,
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

func IdentifyWordType(s string) TokenType {
	k, ok := keywordMap[s]
	if ok {
		return k
	}
	return Var
}
