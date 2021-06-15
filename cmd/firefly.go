package main

import (
	"fmt"
)

func main() {

	msg := `
Hello, firefly!

NEXT: Stater: splits up a slice of tokens into newline separated statements
THEN: Cleaner: removes redundant tokens (e.g. whitespace) and empty statements


LATER: Parser: converts statements (slices of tokens) into parse trees
LATER: Runner: executes a parse tree or slice of parse trees printing out the results
`

	fmt.Println(msg)
}
