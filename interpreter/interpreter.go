package interpreter

import (
	"bytes"
	"io"

	"github.com/shozawa/monkey/evaluator"
	"github.com/shozawa/monkey/lexer"
	"github.com/shozawa/monkey/object"
	"github.com/shozawa/monkey/parser"
)

func Execute(in io.Reader) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(in)
	code := buf.String()
	l := lexer.New(code)
	p := parser.New(l)
	program := p.Parse()
	evaluator.Eval(&program, object.NewEnv())
}
