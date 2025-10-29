
// testing the lexer if they are tokenize properly
package lexer

import (
	"testing"

	"MagicInterpreter/token"
)

// testing token
func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests:= []struct {
		expectedType 	token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	// looping and checking each tests are true
	for i , tt := range tests{
		tok := l.NextToken()

		// if the token type is not expected type, return fatal
		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		
		// if the token Expected Literal is not equal to the Literal, return Fatal error
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test [%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}