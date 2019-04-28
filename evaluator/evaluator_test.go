package evaluator

import (
	"testing"

	"github.com/shozawa/monkey/lexer"
	"github.com/shozawa/monkey/object"
	"github.com/shozawa/monkey/parser"
)

func TestEval(t *testing.T) {
	input := "5;"
	l := lexer.New(input)
	p := parser.New(l)
	program := p.Parse()
	o := Eval(&program)
	integer, ok := o.(*object.Integer)
	if !ok {
		t.Errorf("o not Integer. got=%t\n", o)
	}
	if integer.Value != 5 {
		t.Errorf("integer.Value not 5. got=%d", integer.Value)
	}
}
