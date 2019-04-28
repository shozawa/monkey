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

func (p *Parser) Parse() (prog ast.Program) {
	for p.curToken.Type != token.EOF {
		switch p.curToken.Type {
		case token.LET:
			prog.Statements = append(prog.Statements, p.parseLetStatement())
		// Expression Statement
		case token.IDENT:
			prog.Statements = append(prog.Statements, p.parserExpressionStatement())
		default:
			// TODO: report parse error
			p.nextToken()
		}
	}
	return
}

func (p *Parser) parserExpressionStatement() *ast.ExpressionStatement {
	// skip
	p.nextToken()
	return &ast.ExpressionStatement{}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	letStmt := ast.LetStatement{Token: p.curToken}
	p.nextToken()
	letStmt.Name = p.parseIdentifier()
	// skip assignment
	p.nextToken()
	// skip Value
	p.nextToken()
	// skip semicolon
	p.nextToken()
	return &letStmt
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	ident := ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	p.nextToken()
	return &ident
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
