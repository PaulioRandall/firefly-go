package token

type TokenType int

const (
	Unknown TokenType = iota

	Newline    // '\n'
	Terminator // ;
	Identifier

	Assign // =
	Comma  // ,
	Colon  // :
	Spell  // @

	Space
	Comment

	Bool
	Number
	String

	Def
	If
	For
	In
	Watch
	When
	Is
	Func
	Proc
	End

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

	And // &&
	Or  // ||

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
	Newline:    "Newline",
	Terminator: "Terminator",
	Identifier: "Identifier",
	Space:      "Space",
	Comment:    "Comment",

	Bool:   "Bool",
	Number: "Number",
	String: "String",

	Def:   "Define",
	If:    "If",
	For:   "For",
	In:    "In",
	Watch: "Watch",
	When:  "When",
	Is:    "Is",
	Func:  "Function",
	Proc:  "Procedure",
	End:   "End",

	Assign: "Assign",
	Comma:  "Comma",
	Colon:  "Colon",
	Spell:  "Spell",

	Add: "Add",
	Sub: "Subtract",
	Mul: "Multiply",
	Div: "Divide",
	Mod: "Remainder",

	LT:  "Less Than",
	GT:  "More Than",
	LTE: "Less Than Equal",
	GTE: "More Than Equal",
	EQU: "Equal",
	NEQ: "Not Equal",

	And: "And",
	Or:  "Or",

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
		return 6
	case Add, Sub:
		return 5
	case LT, GT, LTE, GTE:
		return 4
	case EQU, NEQ:
		return 3
	case And:
		return 2
	case Or:
		return 1
	default:
		return 0
	}
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
