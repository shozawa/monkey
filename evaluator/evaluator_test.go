package evaluator

import (
	"testing"

	"github.com/shozawa/monkey/lexer"
	"github.com/shozawa/monkey/object"
	"github.com/shozawa/monkey/parser"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input string
		want  int64
	}{
		{"5;", 5},
		{"5;", 5},
		{"42", 42},
	}
	for _, test := range tests {
		evaluated := testEval(test.input)
		testIntegerObject(t, evaluated, test.want)
	}
}

func TestEvalBoolLiteral(t *testing.T) {
	input := "true;"
	o := testEval(input)
	b, ok := o.(*object.Bool)
	if !ok {
		t.Errorf("o not Bool. got=%t\n", o)
	}
	if b.Value != true {
		t.Errorf("b.Value not true. got=%v", b.Value)
	}
}

func TestEvalIfExpression(t *testing.T) {
	input := "if (false) { 1 } else { 2 };"
	o := testEval(input)
	integer, ok := o.(*object.Integer)
	if !ok {
		t.Errorf("o not Integer. got=%t\n", o)
	}
	if integer.Value != 2 {
		t.Errorf("b.Value not true. got=%v", integer.Value)
	}
}

func TestEvalLetStatement(t *testing.T) {
	input := `
	let five = 5;
	five;
	`
	o := testEval(input)
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
	o := testEval(input)
	integer, ok := o.(*object.Integer)
	if !ok {
		t.Errorf("o not Integer. got=%t\n", o)
	}
	if integer.Value != 7 {
		t.Errorf("integer.Value not 7. got=%d", integer.Value)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "fn(x) { x + 2; };"
	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Errorf("object is not Function. got=%T (%+v).\n", evaluated, evaluated)
	}
	if len(fn.Parameters) != 1 {
		t.Errorf("function has wrong parameters. Parameters=%+v", fn.Parameters)
	}
	// TODO
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input string
		want  int64
	}{
		{"let identity = fn(x) { x; }; identity(5);", 5},
		{"let STEP = 10; let plus = fn(n) { n + STEP }; plus(2);", 12},
	}
	for _, test := range tests {
		obj := testEval(test.input)
		testIntegerObject(t, obj, test.want)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.Parse()
	return Eval(&program, object.NewEnv())
}

func testIntegerObject(
	t *testing.T,
	obj object.Object,
	want int64,
) bool {
	integer, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not IntegerObject. got=%T (%+v).\t", obj, obj)
		return false
	}
	if integer.Value != want {
		t.Errorf("integer.Value not %d. got=%d.\n", want, integer.Value)
		return false
	}
	return true
}
