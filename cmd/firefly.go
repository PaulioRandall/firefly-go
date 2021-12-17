// Program firefly is a toy programming language built on the following ideas:
//
//		Fantasy themed
//		Specification first
//		No install
//		Easily modifiable and extendable (to the extent that is reasonable possible)
package main

import (
	"fmt"
	"os"

	"github.com/PaulioRandall/firefly-go/firefly/firefly"
)

func main() {
	file := getFilenameArg()
	run(file)
}

func getFilenameArg() string {
	args := os.Args
	if len(args) < 2 {
		panic("expected filename argument")
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
`)
}
