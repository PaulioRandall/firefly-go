package token

type TokenType int

const (
	Unknown TokenType = iota

	Newline    // '\n'
	Terminator // ;
	Identifier

	Assign // =
	Define // :=
	Comma  // ,
	Colon  // :
	Spell  // @

	_redundant_begin
	Space
	Comment
	_redundant_end

	_keywords_begin
	If
	For
	In
	Watch
	When
	Is
	Func
	Proc
	End
	_literal_begin
	True
	False
	_keywords_end
	Number
	String
	_literal_end

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
)

var nameMap = map[TokenType]string{
	Newline:      "Newline",
	Terminator:   "Terminator",
	Identifier:   "Identifier",
	Space:        "Space",
	Comment:      "Comment",
	Number:       "Number",
	String:       "String",
	If:           "If",
	For:          "For",
	In:           "In",
	Watch:        "Watch",
	When:         "When",
	Is:           "Is",
	Func:         "Function",
	Proc:         "Procedure",
	End:          "End",
	True:         "True",
	False:        "False",
	Assign:       "Assign",
	Define:       "Define",
	Comma:        "Comma",
	Colon:        "Colon",
	Spell:        "Spell",
	Add:          "Add",
	Sub:          "Subtract",
	Mul:          "Multiply",
	Div:          "Divide",
	Mod:          "Remainder",
	LT:           "Less Than",
	GT:           "More Than",
	LTE:          "Less Than Equal",
	GTE:          "More Than Equal",
	EQU:          "Equal",
	NEQ:          "Not Equal",
	ParenOpen:    "Paren Open",
	ParenClose:   "Paren Close",
	BraceOpen:    "Brace Open",
	BraceClose:   "Brace Close",
	BracketOpen:  "Bracket Open",
	BracketClose: "Bracket Close",
}

func (tt TokenType) String() string {
	return nameMap[tt]
}

func (tt TokenType) Precedence() int {
	switch tt {
	case Mul, Div, Mod:
		return 4
	case Add, Sub:
		return 3
	case LT, GT, LTE, GTE:
		return 2
	case EQU, NEQ:
		return 1
	default:
		return 0
	}
}

func IsLiteral(tt TokenType) bool {
	return tt > _literal_begin && tt < _literal_end
}

func IsRedundant(tt TokenType) bool {
	return tt > _redundant_begin && tt < _redundant_end
}

func IsKeyword(tt TokenType) bool {
	return tt > _keywords_begin && tt < _keywords_end
}

func IsBinaryOperator(tt TokenType) bool {
	return (tt > _arith_begin && tt < _arith_end) ||
		(tt > _cmp_begin && tt < _cmp_end)
}

func filter(f func(TokenType, string) bool) map[TokenType]string {
	res := map[TokenType]string{}

	for tt, symbol := range nameMap {
		if f(tt, symbol) {
			res[tt] = symbol
		}
	}

	return res
}
