package ast

import (
	"bytes"
	"fmt"

	"github.com/shozawa/monkey/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, stmt := range p.Statements {
		out.WriteString(stmt.String())
	}
	return out.String()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) statementNode() {}
func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}
func (l *LetStatement) String() string {
	return fmt.Sprintf("let %v = %v;", l.Name.String(), l.Value.String())
}

type ExpressionStatement struct {
	Expression Expression
}

func (e *ExpressionStatement) statementNode() {}
func (e *ExpressionStatement) TokenLiteral() string {
	return e.Expression.TokenLiteral()
}
func (e *ExpressionStatement) String() string {
	return e.Expression.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) String() string {
	return i.TokenLiteral()
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (i *IntegerLiteral) expressionNode() {}
func (i *IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}
func (i *IntegerLiteral) String() string {
	return i.TokenLiteral()
}

type BoolLiteral struct {
	Token token.Token
	Value string
}

func (b *BoolLiteral) expressionNode() {}
func (b *BoolLiteral) TokenLiteral() string {
	return b.Token.Literal
}
func (b *BoolLiteral) String() string {
	return b.TokenLiteral()
}

type Infix struct {
	Token    token.Token
	Operator string
	Left     Expression
	Right    Expression
}

func (i *Infix) expressionNode() {}
func (i *Infix) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Infix) String() string {
	return fmt.Sprintf("(%s %s %s)", i.Left.String(), i.TokenLiteral(), i.Right.String())
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (b *BlockStatement) statementNode() {}
func (b *BlockStatement) TokenLiteral() string {
	return b.Token.Literal
}
func (b *BlockStatement) String() string {
	return "TODO"
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (i *IfExpression) expressionNode() {}
func (i *IfExpression) TokenLiteral() string {
	return i.Token.Literal
}
func (b *IfExpression) String() string {
	return "TODO"
}
