package parser

import (
	"github.com/shozawa/monkey/ast"
	"github.com/shozawa/monkey/lexer"
	"github.com/shozawa/monkey/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) Parse() ast.Program {
	ident := ast.Identifier{Token: token.Token{Type: token.IDENT, Literal: "five"}, Value: "five"}
	mock := ast.LetStatement{Token: token.Token{Type: token.LET, Literal: "let"}, Name: &ident}
	return ast.Program{Statements: []ast.Statement{&mock}}
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
