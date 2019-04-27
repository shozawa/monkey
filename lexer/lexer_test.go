package lexer

import (
	"testing"

	"github.com/shozawa/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;
	`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// let five = 5;
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		// let ten = 10;
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, test := range tests {
		tok := l.NextToken()
		if tok.Type != test.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expecting=%q, got=%q", i, test.expectedType, tok.Type)
		}
		if tok.Literal != test.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expecting=%q, got=%q", i, test.expectedLiteral, tok.Literal)
		}
	}
}

func TestIsLetter(t *testing.T) {
	tests := []struct {
		input byte
		want  bool
	}{
		{'a', true},
		{'b', true},
		{'y', true},
		{'z', true},
		{'A', true},
		{'B', true},
		{'Y', true},
		{'Z', true},
		{'_', true},
		{'1', false},
		{'!', false},
	}
	for _, test := range tests {
		if got := isLetter(test.input); got != test.want {
			t.Fatalf("input=%v, want=%v, got=%v", test.input, test.want, got)
		}
	}
}
