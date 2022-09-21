package scanner

import (
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

type sidekick struct {
	start token.Pos
	tt    token.TokenType
	val   []rune
}

func (sk *sidekick) add(ru ...rune) {
	sk.val = append(sk.val, ru...)
}

func (sk sidekick) str() string {
	return string(sk.val)
}
