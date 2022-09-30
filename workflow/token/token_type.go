package token

type TokenType int

const (
	Unknown TokenType = iota

	Newline    // '\n'
	Terminator // ;
	Var

	_redundant_begin
	Space
	Comment
	_redundant_end

	_literal_begin
	Number
	String
	_literal_end

	_keywords_begin
	If
	For
	In
	Watch
	When
	Is
	E
	F
	End
	True
	False
	_keywords_end

	_operators_begin
	Assign // =
	Define // :=
	Comma  // ,
	Colon  // :
	Spell  // @

	_arith_begin
	Add // +
	Sub // -
	Mul // *
	Div // /
	Mod // %
	_arith_end

	_cmp_begin
	LT  // <
	GT  // >
	LTE // <=
	GTE // >=
	EQU // ==
	NEQ // !=
	_cmp_end

	_paren_begin
	ParenOpen    // (
	ParenClose   // )
	BraceOpen    // {
	BraceClose   // }
	BracketOpen  // [
	BracketClose // ]
	_paren_end
	_operators_end
)

var symbolMap = map[TokenType]string{
	Newline:      "\n",
	Terminator:   ";",
	Var:          "",
	Space:        " ",
	Comment:      "//",
	Number:       "0",
	String:       `""`,
	If:           "if",
	For:          "for",
	In:           "in",
	Watch:        "watch",
	When:         "when",
	Is:           "is",
	E:            "E",
	F:            "F",
	End:          "end",
	True:         "true",
	False:        "false",
	Assign:       "=",
	Define:       ":=",
	Comma:        ",",
	Colon:        ":",
	Spell:        "@",
	Add:          "+",
	Sub:          "-",
	Mul:          "*",
	Div:          "/",
	Mod:          "%",
	LT:           "<",
	GT:           ">",
	LTE:          "<=",
	GTE:          ">=",
	EQU:          "==",
	NEQ:          "!=",
	ParenOpen:    "(",
	ParenClose:   ")",
	BraceOpen:    "{",
	BraceClose:   "}",
	BracketOpen:  "[",
	BracketClose: "]",
}

func (tt TokenType) IsRedundant() bool {
	return tt > _redundant_begin && tt < _redundant_end
}

func (tt TokenType) IsLiteral() bool {
	return tt > _literal_begin && tt < _literal_end
}

func (tt TokenType) IsKeyword() bool {
	return tt > _keywords_begin && tt < _keywords_end
}

func (tt TokenType) IsOperator() bool {
	return tt > _operators_begin && tt < _operators_end
}

func (tt TokenType) String() string {
	return symbolMap[tt]
}

func Operators() map[TokenType]string {
	return filter(func(tt TokenType, _ string) bool {
		return tt.IsOperator()
	})
}

func IdentifyWordType(s string) TokenType {
	tt := find(func(tt TokenType, symbol string) bool {
		return tt.IsKeyword() && s == symbol
	})

	if tt == Unknown {
		return Var
	}
	return tt
}

func IdentifyOperatorType(s string) TokenType {
	return find(func(tt TokenType, symbol string) bool {
		return tt.IsOperator() && s == symbol
	})
}

func find(f func(TokenType, string) bool) TokenType {
	for tt, symbol := range symbolMap {
		if f(tt, symbol) {
			return tt
		}
	}
	return Unknown
}

func filter(f func(TokenType, string) bool) map[TokenType]string {
	res := map[TokenType]string{}

	for tt, symbol := range symbolMap {
		if f(tt, symbol) {
			res[tt] = symbol
		}
	}

	return res
}
