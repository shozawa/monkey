package parser

import (
	"fmt"
	"testing"

	"github.com/shozawa/monkey/ast"
	"github.com/shozawa/monkey/lexer"
)

func TestParseExpressionStatement(t *testing.T) {
	tests := []struct {
		input string
		want  interface{}
	}{
		{"foo", "foo"},
		{"foo;", "foo"},
		{"42", 42},
		{"42;", 42},
	}
	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.Parse()
		if got := len(program.Statements); got != 1 {
			t.Errorf("len(program.Statements) not 1 got=%d\n", got)
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("program.Statements[0] not ast.ExpressionStatement. got=%t\n", program.Statements[0])
		}
		testLiteralExpression(t, stmt.Expression, test.want)
	}
}

func TestParseInfixExpression(t *testing.T) {
	tests := []struct {
		input string
		left  interface{}
		op    string
		right interface{}
	}{
		{"1 + 2;", 1, "+", 2},
		{"2 * 3;", 2, "*", 3},
		{"1 > 3;", 1, ">", 3},
		{"1 < 3;", 1, "<", 3},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.Parse()
		if len(program.Statements) != 1 {
			t.Errorf("len(program.Statements) not 1. got=%d\n", len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("program.Statemetns[0] not ast.ExpressionStatement. got=%t\n", program.Statements[0])
		}
		testInfixExpression(t, stmt.Expression, test.left, test.op, test.right)
	}

}

func TestParseFunctionParameters(t *testing.T) {
	tests := []struct {
		input      string
		parameters []string
	}{
		{"fn() {}", []string{}},
		{"fn(x) {}", []string{"x"}},
		{"fn(x, y) {}", []string{"x", "y"}},
	}
	for _, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.Parse()
		if got := len(program.Statements); got != 1 {
			t.Errorf("len(program.Statements) not 1. got=%d.\n", got)
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("program.Statements[0] not ExpressionStatement. got=%t.\n", program.Statements[0])
		}
		fn, ok := stmt.Expression.(*ast.FunctionLiteral)
		if !ok {
			t.Errorf("stmt.Expression not ast.FunctionLiteral. got=%t.\n", stmt.Expression)
		}
		if got := len(fn.Parameters); got != len(test.parameters) {
			t.Errorf("len(fn.Parameters) not %d. got=%d.\n", len(test.parameters), got)
		}
		for i, ident := range test.parameters {
			testIdentifier(t, fn.Parameters[i], ident)
		}
	}
}

func TestParseFunctionLiteral(t *testing.T) {
	input := "fn(x, y) { x + y; }"
	l := lexer.New(input)
	p := New(l)
	program := p.Parse()
	if got := len(program.Statements); got != 1 {
		t.Errorf("len(program.Statements) not 1. got=%d.\n", got)
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("program.Statements[0] not ast.ExpressionStatement. got=%t.\n", program.Statements[0])
	}
	fn, ok := stmt.Expression.(*ast.FunctionLiteral)
	if !ok {
		t.Errorf("stmt.Expression not ast.FunctionLiteral. got=%t.\n", stmt.Expression)
	}
	if got := len(fn.Parameters); got != 2 {
		t.Errorf("len(fn.Parameters) not 2. got=%d.\n", got)
	}

	// test params
	testIdentifier(t, fn.Parameters[0], "x")
	testIdentifier(t, fn.Parameters[1], "y")

	if got := len(fn.Body.Statements); got != 1 {
		t.Errorf("len(fn.Body.Statements) not 1. got=%d.\n", got)
	}
	stmt, ok = fn.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("fn.Body.Statements[0] not ast.ExpressionStatement. got=%t.\n", fn.Body.Statements[0])
	}
	testInfixExpression(t, stmt.Expression, "x", "+", "y")
}

func TestCallExpression(t *testing.T) {
	input := "add(1, 2 + 3, 4 * 5);"
	l := lexer.New(input)
	p := New(l)
	program := p.Parse()
	if got := len(program.Statements); got != 1 {
		t.Errorf("len(program.Statements) not 1. got=%d.\n", got)
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("program.Statements[0] not ast.ExpressionStatement. got=%t.\n", program.Statements[0])
	}
	call, ok := stmt.Expression.(*ast.CallExpression)
	if !ok {
		t.Errorf("stmt.Expression not ast.CallExpression. got=%t.\n", stmt.Expression)
	}
	testIdentifier(t, call.Function, "add")
	if got := len(call.Arguments); got != 3 {
		t.Errorf("len(call.Arguments) not 3. got=%d.\n", got)
	}
	testLiteralExpression(t, call.Arguments[0], 1)
	testInfixExpression(t, call.Arguments[1], 2, "+", 3)
	testInfixExpression(t, call.Arguments[2], 4, "*", 5)
}

func TestParseLetStatement(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;
	`
	tests := []struct {
		name  string
		value int64
	}{
		{name: "five", value: 5},
		{name: "ten", value: 10},
	}
	program := testParse(t, input)
	for i, test := range tests {
		s := program.Statements[i]
		testLetStatment(t, s, test.name, test.value)
	}
}

func TestParseReturnStatement(t *testing.T) {
	input := "return 10;"
	program := testParse(t, input)
	if got := len(program.Statements); got != 1 {
		t.Errorf("len(program.Statements) not 1. got=%d", got)
	}
	stmt, ok := program.Statements[0].(*ast.ReturnStatement)
	if !ok {
		t.Errorf("program.Statements[0] is not ast.ReturnStatement. got=%T", program.Statements[0])
	}
	if stmt.TokenLiteral() != "return" {
		t.Errorf("stmt.TokenLiteral() is not 'return'. got=%q", stmt.TokenLiteral())
	}
}

func testLetStatment(t *testing.T, s ast.Statement, name string, value int64) bool {
	if literal := s.TokenLiteral(); literal != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", literal)
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s is not *ast.LetStatment. got=%T", s)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}
	integerLiteral, ok := letStmt.Value.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("letStmt.Value not ast.IntegerLiteral. got=%t\n", letStmt.Value)
		return false
	}
	if got := integerLiteral.Value; got != value {
		t.Errorf("integerLiteral.Value not %d. got=%d\n", value, got)
		return false
	}
	return true
}

func TestIfExpression(t *testing.T) {
	input := "if (x) { 1 } else { 2 };"
	l := lexer.New(input)
	p := New(l)
	program := p.Parse()
	if got := len(program.Statements); got != 1 {
		t.Errorf("len(program.Statements) not 1. got=%d\n", got)
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("program.Statements[0] not ExpressionStatement. got=%t\n", program.Statements[0])
	}
	exp, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Errorf("stmt.Expression not ast.IfExpression. got=%t\n", stmt.Expression)
	}
	condition := exp.Condition
	if condition.TokenLiteral() != "x" {
		t.Errorf("condition.TokenLiteral() not 'x'. got=%q\n", condition.TokenLiteral())
	}
}

func TestBoolLiteral(t *testing.T) {
	input := `
	true;
	false;
	`
	l := lexer.New(input)
	p := New(l)
	program := p.Parse()
	if got := len(program.Statements); got != 2 {
		t.Errorf("len(program.Statements) not 2. got=%d\n", got)
	}
}

func TestOperatorPrecedence(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"1 + 1", "(1 + 1)"},
		{"1 * 2", "(1 * 2)"},
		{"1 + 2 + 3", "((1 + 2) + 3)"},
		{"1 + 2 * 3", "(1 + (2 * 3))"},
		{"(1 + 2) * 3", "((1 + 2) * 3)"},
	}
	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.Parse()
		if got := program.String(); got != test.want {
			t.Errorf("[%d] program.String() not %q. input=%q, got=%q\n", i, test.want, test.input, got)
		}
	}
}

func testLiteralExpression(t *testing.T, exp ast.Expression, want interface{}) bool {
	switch v := want.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testIdentifier(t, exp, v)
	}
	t.Errorf("type of exp not handled. got=%t.\n", exp)
	return false
}

func testInfixExpression(
	t *testing.T,
	exp ast.Expression,
	left interface{},
	op string,
	right interface{}) bool {
	infixExp, ok := exp.(*ast.Infix)
	if !ok {
		t.Errorf("exp not ast.Infix. got=%T(%s).\n", exp, exp)
		return false
	}
	if !testLiteralExpression(t, infixExp.Left, left) {
		return false
	}
	if infixExp.Operator != op {
		t.Errorf("infixExp.Operator not %q. got=%q.\n", op, infixExp.Operator)
		return false
	}
	if !testLiteralExpression(t, infixExp.Right, right) {
		return false
	}
	return true
}

func testIntegerLiteral(t *testing.T, exp ast.Expression, want int64) bool {
	lit, ok := exp.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("lit not ast.IntegerLIteral. got=%t.\n", exp)
		return false
	}
	if lit.Value != want {
		t.Errorf("lit.Value not %d. got=%d.\n", want, lit.Value)
		return false
	}
	if got := lit.TokenLiteral(); got != fmt.Sprintf("%d", want) {
		t.Errorf("lit.TokenLiteral() not %d. got=%s.\n", want, got)
	}
	return true
}

func testIdentifier(t *testing.T, exp ast.Expression, want string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not ast.Identifier. got=%t.\n", exp)
		return false
	}
	if ident.Value != want {
		t.Errorf("ident.Value not %q. got=%q.\n", want, ident.Value)
		return false
	}
	if got := ident.TokenLiteral(); got != want {
		t.Errorf("ident.TokenLiteral() not %q. got=%q.\n", want, got)
		return false
	}
	return true
}

func testParse(t *testing.T, input string) ast.Program {
	l := lexer.New(input)
	p := New(l)
	program := p.Parse()
	checkParserErrors(t, p)
	return program
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
