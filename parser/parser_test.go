package parser

import (
	"testing"

	"github.com/shozawa/monkey/ast"
	"github.com/shozawa/monkey/lexer"
)

func TestParseExpressionStatement(t *testing.T) {
	input := "foo;"
	l := lexer.New(input)
	p := New(l)
	program := p.Parse()
	if got := len(program.Statements); got != 1 {
		t.Errorf("len(program.Statements) not 1 got=%d\n", got)
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("program.Statements[0] not ast.ExpressionStatement. got=%t\n", program.Statements[0])
	}
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Errorf("stmt.Expression not ast.Identifier. got=%t\n", stmt.Expression)
	}
	if got := ident.TokenLiteral(); got != "foo" {
		t.Errorf("ident.TokenLiteral not 'foo'. got=%q\n", got)
	}
}

func TestMultilineExpressionStatement(t *testing.T) {
	input := `
	foo;
	bar;
	`
	l := lexer.New(input)
	p := New(l)
	program := p.Parse()
	if got := len(program.Statements); got != 2 {
		t.Errorf("len(program.Statements) not 2. got=%d\n", got)
	}
}

func TestParseIntLiteral(t *testing.T) {
	input := "42;"
	l := lexer.New(input)
	p := New(l)
	program := p.Parse()
	if got := len(program.Statements); got != 1 {
		t.Errorf("len(program.Statements) not 1. got=%d\n", got)
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("program.Statements[0] not ast.ExpressionStatement. got=%t\n", program.Statements[0])
	}
	integerLiteral, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("stmt.Expression not ast.IntegerLiteral. got=%t\n", stmt.Expression)
	}
	if integerLiteral.Value != 42 {
		t.Errorf("integerLiteral.Value not 42. got=%d\n", integerLiteral.Value)
	}
}

func TestParsePlus(t *testing.T) {
	input := "1 + 2;"
	l := lexer.New(input)
	p := New(l)
	program := p.Parse()
	if len(program.Statements) != 1 {
		t.Errorf("len(program.Statements) not 1. got=%d\n", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Errorf("program.Statemetns[0] not ast.ExpressionStatement. got=%t\n", program.Statements[0])
	}
	plus, ok := stmt.Expression.(*ast.Infix)
	if !ok {
		t.Errorf("stmt.Expression not ast.Infix. got=%t\n", stmt.Expression)
	}
	left, ok := plus.Left.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("plus.Left not ast.IntegerLiteral. got=%t\n", plus.Left)
	}
	if left.Value != 1 {
		t.Errorf("left.Value not 1. got=%d\n", left.Value)
	}
	right, ok := plus.Right.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("plus.Right not ast.IntegerLiteral. got=%t\n", plus.Right)
	}
	if right.Value != 2 {
		t.Errorf("left.Value not 1. got=%d\n", right.Value)
	}
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
	l := lexer.New(input)
	p := New(l)
	program := p.Parse()
	for i, test := range tests {
		s := program.Statements[i]
		testLetStatment(t, s, test.name, test.value)
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
