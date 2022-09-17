package scanner

import (
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

type Scan func(tk token.Token, f Scan, e error)

func ScanAll() []token.Token {
	return []token.Token{}
}
