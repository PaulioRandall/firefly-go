package main

import (
	"fmt"
)

func main() {

	msg := `
Hello, firefly!

NEXT: Runner: executes a parse tree or slice of parse trees printing out the results

scanner -> grouper -> cleaner -> parser
`

	fmt.Println(msg)
}
