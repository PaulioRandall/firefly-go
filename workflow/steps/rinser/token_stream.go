package rinser

type tokenStream interface {
	access() token.Token
	accept(token.TokenType) bool
	expect(token.TokenType) error
}
