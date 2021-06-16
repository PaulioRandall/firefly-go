package main

import (
	"fmt"
)

func main() {

	msg := `
Hello, firefly!

NEXT: Parser: converts statements (slices of tokens) into parse trees


LATER: Runner: executes a parse tree or slice of parse trees printing out the results

scanner -> grouper -> cleaner
`

	fmt.Println(msg)
}
