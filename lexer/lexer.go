package lexer

import (
	"fmt"

	"github.com/therandombyte/go-interpreter/token"
)

// ------------------------------------------------------------
// Background:
// Source Code -> Lexer -> Token
// Instead of reading the source code from a file or shell,
// we just give it a string to avoid complexity
// ------------------------------------------------------------

// Lexer takes an input string and reads character by character.
// This character will be converted into a Token
type Lexer struct {
	input        string
	position     int  // current position pointing to ch
	readPosition int  // the next character position
	ch           byte // only supports ascii, so byte will suffice
}

// Creates a Lexer with provided input string and initializes position values.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken will get the current char/literal from Lexer
// and creates a Token. If the literal is a keyword, then look it up.
// otherwise, mark it as illegal
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
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
		if isLetter(l.ch) {
			// If literal is unidentified, check if its a keyword,
			// and lookup its type.
			// otherwise mark it as illegal.
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			fmt.Println("Case Letter ", tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			fmt.Println("Case Digit ", tok.Literal)
			return tok
		} else {
			fmt.Println("Case Illegal ", l.ch)
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

// readIdentifier() will read the next identifier/char from lexer's input.
// Calls readChar() repeatedly
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// newToken() creats a new token for each input literal
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// isLetter() defines what is allowed as a letter in our language.
// Currently its (a-z,A-z,_)
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// readChar() will get the next character from input.
// If it reaches end of input, then output 0 (EOF).
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition] // this works when the next char is one byte
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
