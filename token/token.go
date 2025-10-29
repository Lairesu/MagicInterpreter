// Token = token are the smallest meaningful units in the source code
// such as keywords, operators, identifiers, literals and delimiters



package token

// TokenType is just a custom type based on string
// TokenType represents the type of category of a token
// It is defined as a string for readability and easy comparison
type TokenType string


// Predefined token types used by the lexer
// These constants classify the different  kinds of that can appear
// in the source code
const (
	// special toke
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	
	// Identifiers + literals
	IDENT = "IDENT"
	INT = "INT"

	// operators
	ASSIGN = "="
	PLUS = "+"

	// Delimiters
	COMMA = ","
	SEMICOLON  = ","
	
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)

// Token represents a single lexical token extracted from the source code
// Type : The category/Type of the token(operator, keywords, literals etc)
// Literal : the exact string value from the source code
type Token struct {
	Type TokenType
	Literal string
}

