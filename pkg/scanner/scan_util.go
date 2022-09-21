package scanner

import (
	"github.com/PaulioRandall/firefly-go/pkg/token"
)

type sidekick struct {
	start token.Pos
	tt    token.TokenType
	val   []rune
}

func (sk *sidekick) addIf(r Reader, want rune) (bool, error) {
	if !r.More() {
		return false, nil
	}

	have, e := r.Peek()
	if e != nil {
		return false, e
	}

	if have != want {
		return false, nil
	}

	if _, e = r.Read(); e != nil {
		return false, e
	}

	sk.add(have)
	return true, nil
}

func (sk *sidekick) add(ru ...rune) {
	sk.val = append(sk.val, ru...)
}

func (sk sidekick) str() string {
	return string(sk.val)
}
