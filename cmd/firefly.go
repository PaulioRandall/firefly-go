package main

import (
	"fmt"
	"os"

	"github.com/PaulioRandall/firefly-go/pkg/firefly"
)

func main() {
	file := getFilenameArg()
	run(file)
}

func getFilenameArg() string {
	args := os.Args
	if len(args) < 2 {
		panic("Expected filename argument")
	}
	return args[1]
}

func run(file string) {
	e := firefly.RunFile(file)
	if e != nil {
		panic(e)
	}
}

func printTaskList() {
	fmt.Print(`
Hello, firefly!

NEXT: Add godoc support to godo
NEXT: Write documentation for pkgs
	- token
	- scanner
	- grouper
	- cleaner
	- ast
	- parser
	- runner
`)
}
