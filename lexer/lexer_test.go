package lexer

import (
	"testing"

	"github.com/shozawa/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `
	5 > 3;
	5 < 3;
	1 + 2 - 3 * 4 / 5;
	let five = 5;
	let ten = 10;
	1 == 1;
	1 != 1;
	!false;
	`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// 5 > 3;
		{token.INT, "5"},
		{token.GT, ">"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		// 5 < 3;
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		//  1 + 2 - 3 * 4 / 5;
		{token.INT, "1"},
		{token.PLUS, "+"},
		{token.INT, "2"},
		{token.MINUS, "-"},
		{token.INT, "3"},
		{token.ASTERISK, "*"},
		{token.INT, "4"},
		{token.SLASH, "/"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
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
		// 1 == 1;
		{token.INT, "1"},
		{token.EQ, "=="},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		// 1 =! 1;
		{token.INT, "1"},
		{token.NOT_EQ, "!="},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		// !false;
		{token.BANG, "!"},
		{token.FALSE, "false"},
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

func TestConsumeStringLiteral(t *testing.T) {
	tests := []struct {
		input string
		want  []token.Token
	}{
		{
			`let str = "hello";`,
			[]token.Token{
				token.Token{Type: token.LET, Literal: "let"},
				token.Token{Type: token.IDENT, Literal: "str"},
				token.Token{Type: token.ASSIGN, Literal: "="},
				token.Token{Type: token.STRING, Literal: "hello"},
				token.Token{Type: token.SEMICOLON, Literal: ";"},
			},
		},
	}
	for _, test := range tests {
		l := New(test.input)
		for _, want := range test.want {
			tok := l.NextToken()
			if tok != want {
				t.Errorf("tok is not %v. got=%v", want, tok)
			}
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

func TestIfStatement(t *testing.T) {
	input := "if (true) { 1 } else { 2 };"
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		// if (true) { 1 } else { 2 };
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.TRUE, "true"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.INT, "1"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.INT, "2"},
		{token.RBRACE, "}"},
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
