package main

import (
	"fmt"
)

func main() {

	msg := `
Hello, firefly!

THEN: Cleaner: removes redundant tokens (e.g. whitespace)


LATER: Parser: converts statements (slices of tokens) into parse trees
LATER: Runner: executes a parse tree or slice of parse trees printing out the results
`

	fmt.Println(msg)
}
