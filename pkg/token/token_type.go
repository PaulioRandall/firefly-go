package token

type TokenType int

const (
	Unknown TokenType = iota

	// Keywords
	If // 1
	For
	Watch
	When
	E
	F
	End
	True
	False // 9

	// Variables
	Var
)

var syntaxMap = map[TokenType]string{
	If:    "if",
	For:   "for",
	Watch: "watch",
	When:  "when",
	E:     "E",
	F:     "F",
	End:   "end",
	True:  "true",
	False: "false",
	Var:   "variable",
}

var keywords = map[string]TokenType{}

func init() {
	for k, v := range syntaxMap {
		if k.IsKeyword() {
			keywords[v] = k
		}
	}
}

func (tt TokenType) IsKeyword() bool {
	return tt >= 1 && tt <= 9
}

func (tt TokenType) String() string {
	return syntaxMap[tt]
}

func IdentifyWordType(s string) TokenType {
	k, ok := keywords[s]
	if ok {
		return k
	}
	return Var
}
