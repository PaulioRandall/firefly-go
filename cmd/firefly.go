package main

import (
	"fmt"
)

func main() {

	msg := `
Hello, firefly!

NEXT: Continue testing firefly pkg
NEXT: Update main to execute scroll
NEXT: Add godoc support to godo
NEXT: Write documentation for pkgs
	- token
	- scanner
	- grouper
	- cleaner
	- ast
	- parser
	- runner
`

	fmt.Print(msg)
}
