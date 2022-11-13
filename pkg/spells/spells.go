package spells

type Spell func(mem Memory, params []any) []any

type Memory interface {
	Var(name string) any
	Spell(name string) Spell
}

var All = map[string]Spell{
	"meh": Spell_meh,
	"len": Spell_len,
}

func Spell_meh(mem Memory, params []any) []any {
	return nil
}
