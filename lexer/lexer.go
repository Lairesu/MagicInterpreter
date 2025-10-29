package lexer

import (
	"MagicInterpreter/token"
)

// lexer represents the state of lexical analyzer
// It's job is to read through the source code and create stream of tokens
type Lexer struct {
	input 			string 	// the entire source code
	position 		int		// current position in input (pointer)
	readPosition 	int		// current reading position in input (position after the current position)
	ch 				byte	// current char
}


// New creates and returns a new lexer instance for the given input string
// It initializes the lexer and reads the first character so its ready to produce token
func New(input string) *Lexer {
	l := &Lexer{input: input}
	// call read char 
	l.readChar()
	return l
}

// ========================
// ReadChar
// ========================
// advances the lexer by one character
// updates the current character , the current position and the curr read Position
// when EOF is reached, sets ch to 0(ASCII-NULL) to signal EOF

func(l *Lexer) readChar() {
	// if the next position is greater than the len it will be EOF so return 0
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else { // ese set the current character to next position
		l.ch = l.input[l.readPosition]
	}
	// set the position to be next position
	l.position = l.readPosition
	// and increase  the read position by 1
	l.readPosition += 1
}

// ========================
// Next token
// ========================

// NextToken returns the next token from the input
func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    switch l.ch {
    case '=':
        tok = newToken(token.ASSIGN, l.ch)
    case ';':
        tok = newToken(token.SEMICOLON, l.ch)
    case '(':
        tok = newToken(token.LPAREN, l.ch)
    case ')':
        tok = newToken(token.RPAREN, l.ch)
    case ',':
        tok = newToken(token.COMMA, l.ch)
    case '+':
        tok = newToken(token.PLUS, l.ch)
    case '{':
        tok = newToken(token.LBRACE, l.ch)
    case '}':
        tok = newToken(token.RBRACE, l.ch)
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
    default:
        // TODO: handle identifiers, numbers, illegal characters
    }

    l.readChar()
    return tok
}

// newToken creates a new token from a single character
func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(ch)}
}