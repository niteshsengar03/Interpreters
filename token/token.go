package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string 
}

//defining a set of constants in Go that represent the possible token types
const (
	// WE DONT KNOW ABOUT THE TOKEN
	ILLEGAL = "ILLEGAL"
	//END OF FILE, TELLS PARSER TO STOP
	EOF = "EOF"

	
	IDENT = "IDENT"
	INT = "INT"
	
	// OPERATOR
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"

	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	lBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET = "LET"
)