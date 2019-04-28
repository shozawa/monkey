package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"let":  LET,
	"if":   IF,
	"else": ELSE,
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "ASSIGN"
	PLUS   = "PLUS"

	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
)

func LookupIdent(ident string) Token {
	if t, ok := keywords[ident]; ok {
		return Token{Type: t, Literal: ident}
	}
	return Token{Type: IDENT, Literal: ident}
}
