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
		{"-5", -5},
		{"1 + 1;", 2},
		{"10 / 5", 2},
		{"1 + 2 + 3;", 6},
		{"1 + 2 * 3", 7},
		{"(1 + 2) * 3", 9},
		{"2 * ((1 + 2) * 3)", 18},
	}
	for _, test := range tests {
		evaluated := testEval(test.input)
		testIntegerObject(t, evaluated, test.want)
	}
}

func TestEvalStringExpression(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{`"hello, world"`, "hello, world"},
	}
	for _, test := range tests {
		evaluated := testEval(test.input)
		str, ok := evaluated.(*object.String)
		if !ok {
			t.Errorf("evaluated is not String. got=%T", evaluated)
		}
		if str.Value != test.want {
			t.Errorf("str.Value is not %q. got=%q", test.want, str.Value)
		}
	}
}

func TestEvalBoolExpression(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"true;", true},
		{"true", true},
		{"false;", false},
		{"1 == 1", true},
		{"1 == 2", false},
		{"1 != 2", true},
		{"1 != 1", false},
		{"1 < 3;", true},
		{"3 < 1;", false},
		{"3 > 1;", true},
		{"1 > 3;", false},
	}
	for _, test := range tests {
		evaluated := testEval(test.input)
		testBoolObject(t, evaluated, test.want)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
	}
	for _, test := range tests {
		evaluated := testEval(test.input)
		testBoolObject(t, evaluated, test.want)
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

func TestIfElseExpression(t *testing.T) {
	tests := []struct {
		input string
		want  interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
	}
	for _, test := range tests {
		evaluated := testEval(test.input)
		integer, ok := test.want.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
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

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input string
		want  int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{`
		if (true) {
			if (true) {
				return 10;
			}
			return 9;
		}
		`, 10},
	}
	for _, test := range tests {
		evaluated := testEval(test.input)
		testIntegerObject(t, evaluated, test.want)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"5 + true", "type mismatch: INTEGER + BOOLEAN"},
		{"5 + true; 9;", "type mismatch: INTEGER + BOOLEAN"},
		{`
		if (true) {
			if (true) {
				5 + true;
			}
			9;
		}
		`, "type mismatch: INTEGER + BOOLEAN"},
	}
	for _, test := range tests {
		evaluated := testEval(test.input)
		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)
			continue
		}
		if errObj.Message != test.want {
			t.Errorf("wrong error message. expected=%q, got=%q", test.want, errObj.Message)
		}
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
		{`
		let fib = fn(x) {
			if (x < 2) {
				return x;
			} else {
				return fib(x - 1) + fib(x - 2);
			}
		};
		fib(10);
		`, 55},
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
		t.Errorf("object is not IntegerObject. got=%T (%+v).\n", obj, obj)
		return false
	}
	if integer.Value != want {
		t.Errorf("integer.Value not %d. got=%d.\n", want, integer.Value)
		return false
	}
	return true
}

func testBoolObject(
	t *testing.T,
	obj object.Object,
	want bool,
) bool {
	boolObj, ok := obj.(*object.Bool)
	if !ok {
		t.Errorf("object is not BoolObject. got=%T (%+v).\n", obj, obj)
		return false
	}
	if boolObj.Value != want {
		t.Errorf("boolObj.Value not %t. got=%t.\n", want, boolObj.Value)
		return false
	}
	return true
}

func testNullObject(
	t *testing.T,
	obj object.Object,
) bool {
	if obj != NULL {
		t.Errorf("object is not Null. got=%T (%+v)", obj, obj)
		return false
	}
	return true
}
