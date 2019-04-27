package parser

import (
	"testing"

	"github.com/shozawa/monkey/ast"
	"github.com/shozawa/monkey/lexer"
)

func TestParseLetStatement(t *testing.T) {
	input := "let five = 5;"
	l := lexer.New(input)
	p := New(l)
	program := p.Parse()
	s := program.Statements[0]
	testLetStatment(t, s, "five")
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