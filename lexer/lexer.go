package lexer

import (
	"MagicInterpreter/token"
)

// lexer represents the state of lexical analyzer
// It's job is to read through the source code and create stream of tokens
type Lexer struct {
	input        string // the entire source code
	position     int    // current position in input (pointer)
	readPosition int    // current reading position in input (position after the current position)
	ch           byte   // current char
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

func (l *Lexer) readChar() {
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

// NextToken scans the input and returns the next lexical token
// It is the core function of the lexer, producing a stream of tokens from the raw source code
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
			l.readChar()
			return tok
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}

	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
			l.readChar()
			return tok
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// newToken creates a new token from a single character
// helper function that creates new token struct: tokenType - category of the token, ch - the actual token
// returns token struct populated with Type and Literals
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// ==========================
// Read Identifier, isLetter , skipWhiteSpace, readNumber, isDigit, peekChar
// ==========================
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// helper function to determine the start and end of identifier
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// helper function to skip whitespace from the source code
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// helper function to read the sequence of digits starting at lexer current position
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// isDigit reports whether the given byte is digit , helps to determine if the current position is number
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// returns the next character int he input without advancing the lexer position
// looks at l.readPosition and returns the byte at that position. if beyond EOF returns 0 (ASCII NULL),
// useful for check two character operation like `!=` and `==`
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
