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
	o := Eval(&program, nil)
	integer, ok := o.(*object.Integer)
	if !ok {
		t.Errorf("o not Integer. got=%t\n", o)
	}
	if integer.Value != 5 {
		t.Errorf("integer.Value not 5. got=%d", integer.Value)
	}
}

func TestEvalBoolLiteral(t *testing.T) {
	input := "true;"
	l := lexer.New(input)
	p := parser.New(l)
	program := p.Parse()
	o := Eval(&program, nil)
	b, ok := o.(*object.Bool)
	if !ok {
		t.Errorf("o not Bool. got=%t\n", o)
	}
	if b.Value != true {
		t.Errorf("b.Value not true. got=%v", b.Value)
	}
}

func TestEvalLetStatement(t *testing.T) {
	input := `
	let five = 5;
	five;
	`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.Parse()
	o := Eval(&program, object.NewEnv())
	integer, ok := o.(*object.Integer)
	if !ok {
		t.Errorf("o not Integer. got=%t\n", o)
	}
	if integer.Value != 5 {
		t.Errorf("integer.Value not 5. got=%d", integer.Value)
	}
}
func TestEvalPlus(t *testing.T) {
	input := `
	let five = 5;
	let two = 2;
	five + two;
	`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.Parse()
	o := Eval(&program, object.NewEnv())
	integer, ok := o.(*object.Integer)
	if !ok {
		t.Errorf("o not Integer. got=%t\n", o)
	}
	if integer.Value != 7 {
		t.Errorf("integer.Value not 7. got=%d", integer.Value)
	}
}
