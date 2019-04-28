package parser

import (
	"strconv"

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
		if stmt := p.parseStatement(); stmt != nil {
			prog.Statements = append(prog.Statements, stmt)
		}
	}
	return
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	// Expression Statement
	case token.IDENT, token.INT:
		return p.parserExpressionStatement()
	default:
		// TODO: report parse error
		p.nextToken()
		return nil
	}
}

func (p *Parser) parserExpressionStatement() *ast.ExpressionStatement {
	stmt := ast.ExpressionStatement{}
	stmt.Expression = p.parseExpression()

	return &stmt
}

func (p *Parser) parseExpression() ast.Expression {
	var left ast.Expression
	switch p.curToken.Type {
	case token.IDENT:
		left = p.parseIdentifier()
	case token.INT:
		left = p.parserIntegerLiteral()
	default:
		left = nil
	}
	// FIXME
	if p.curToken.Type == token.PLUS {
		return p.parseInfix(left)
	}
	return left
}

func (p *Parser) parseInfix(left ast.Expression) *ast.Infix {
	infix := &ast.Infix{}
	infix.Token = p.curToken
	infix.Left = left
	p.nextToken()
	infix.Right = p.parseExpression()
	return infix
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	letStmt := ast.LetStatement{Token: p.curToken}
	p.nextToken()
	letStmt.Name = p.parseIdentifier()
	// skip assignment
	p.nextToken()
	if p.curToken.Type == token.INT {
		letStmt.Value = p.parserIntegerLiteral()
	} else {
		// skip Value
		p.nextToken()
	}
	// skip semicolon
	p.nextToken()
	return &letStmt
}

func (p *Parser) parseIdentifier() *ast.Identifier {
	ident := ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	p.nextToken()
	return &ident
}

func (p *Parser) parserIntegerLiteral() *ast.IntegerLiteral {
	// TODO: error handling
	i, _ := strconv.Atoi(p.curToken.Literal)
	integerLiteral := ast.IntegerLiteral{Token: p.curToken, Value: int64(i)}
	p.nextToken()
	return &integerLiteral
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
