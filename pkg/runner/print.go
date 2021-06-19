package runner

import (
	"fmt"
	"io"
)

func (in *interpreter) print(s string)      { output(in.stdout, s) }
func (in *interpreter) println(s string)    { output(in.stdout, s+"\n") }
func (in *interpreter) printErr(s string)   { output(in.stderr, s) }
func (in *interpreter) printlnErr(s string) { output(in.stderr, s+"\n") }

func output(w io.Writer, s string) {
	_, e := fmt.Fprintf(w, s)
	if e != nil {
		panic(e)
	}
}
