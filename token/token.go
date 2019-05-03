package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"let":    LET,
	"return": RETURN,
	"if":     IF,
	"fn":     FUNCTION,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT  = "IDENT"
	INT    = "INT"
	RETURN = "RETURN"

	ASSIGN   = "ASSIGN"
	PLUS     = "PLUS"
	ASTERISK = "ASTERISK"

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"

	LT = "<"
	GT = ">"
)

func LookupIdent(ident string) Token {
	if t, ok := keywords[ident]; ok {
		return Token{Type: t, Literal: ident}
	}
	return Token{Type: IDENT, Literal: ident}
}
