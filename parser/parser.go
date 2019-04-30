package parser

import (
	"fmt"
	"strconv"

	"github.com/shozawa/monkey/ast"
	"github.com/shozawa/monkey/lexer"
	"github.com/shozawa/monkey/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -x or !x
	CALL        // myFunction()
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parserIntegerLiteral)

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
	// TODO: refactor
	case token.IDENT, token.INT, token.IF, token.LPAREN, token.TRUE, token.FALSE:
		return p.parserExpressionStatement()
	default:
		// TODO: report parse error
		p.nextToken()
		panic("parseStatement error")
		// return nil
	}
}

func (p *Parser) parserExpressionStatement() *ast.ExpressionStatement {
	stmt := ast.ExpressionStatement{}
	stmt.Expression = p.parseExpression()

	// peek?
	if p.curToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	return &stmt
}

func (p *Parser) parseExpression() ast.Expression {
	if p.curToken.Type == token.LPAREN {
		p.nextToken()
	}
	var left ast.Expression
	switch p.curToken.Type {
	case token.IDENT:
		left = p.parseIdentifier()
	case token.INT:
		left = p.parserIntegerLiteral()
	case token.IF:
		left = p.parseIfExpression()
	case token.TRUE, token.FALSE:
		left = p.parseBoolLiteral()
	default:
		left = nil
	}
	if p.curToken.Type == token.LPAREN {
		p.nextToken()
	}
	// FIXME
	if p.curToken.Type == token.PLUS || p.curToken.Type == token.ASTERISK {
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

func (p *Parser) parseIfExpression() *ast.IfExpression {
	ifExpression := &ast.IfExpression{}
	ifExpression.Token = p.curToken
	p.nextToken()
	ifExpression.Condition = p.parseExpression()
	p.nextToken() // parse ')'
	ifExpression.Consequence = p.parseBlockStatement()
	if p.curToken.Type == token.ELSE {
		p.nextToken()
		ifExpression.Alternative = p.parseBlockStatement()
	}
	return ifExpression
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{}
	block.Token = p.curToken
	p.nextToken() // consume '{'
	for p.curToken.Type != token.RBRACE {
		if stmt := p.parseStatement(); stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
	}
	p.nextToken()
	return block
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parserIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

func (p *Parser) parseBoolLiteral() *ast.BoolLiteral {
	literal := &ast.BoolLiteral{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
	p.nextToken()
	return literal
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) peekTokenIs(tokenType token.TokenType) bool {
	return p.peekToken.Type == tokenType
}
