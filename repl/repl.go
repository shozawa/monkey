package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/shozawa/monkey/evaluator"
	"github.com/shozawa/monkey/lexer"
	"github.com/shozawa/monkey/object"
	"github.com/shozawa/monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnv()
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.Parse()
		obj := evaluator.Eval(&program, env)
		if obj != nil {
			fmt.Printf("%q\n", obj.Inspect())
		} else {
			// FIXME: print correct value
			fmt.Print("nil\n")
		}
	}
}
