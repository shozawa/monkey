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

func TestParseLetStatement(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;
	`
	tests := []struct {
		name string
	}{
		{name: "five"},
		{name: "ten"},
	}
	l := lexer.New(input)
	p := New(l)
	program := p.Parse()
	for i, test := range tests {
		s := program.Statements[i]
		testLetStatment(t, s, test.name)
	}
}

func testLetStatment(t *testing.T, s ast.Statement, name string) bool {
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
	return true
}
