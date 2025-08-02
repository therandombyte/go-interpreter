package lexer

import "github.com/therandombyte/go-interpreter/token"

// ------------------------------------------------------------
// Background:
// Source Code -> Lexer -> Token
// Instead of reading the source code from a file or shell,
// we just give it a string to avoid complexity
// ------------------------------------------------------------

// Lexer will take an input string and read character by character
// in this case, the character will be a token, that has a valid type
type Lexer struct {
	input        string
	position     int  // current position pointing to ch
	readPosition int  // the next character position
	ch           byte // only supports ascii, so byte will suffice
}

// for any struct, if its fields are not straight-forward to initialize,
// then invoke an init function in the constructor
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar() will get the next character in input
// if we reach the end of input, then EOF
// otherwise, set ch to the next character,
// point "position" also to the next character,
// and then increment "readPosition"
// the character returned by readChar() will be used by NextToken() to create a Token
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition] // this works when the next char is one byte
	}
	l.position = l.readPosition
	l.readPosition += 1
}

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
	case '0':
		tok.Literal = " "
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			// setting literal, but type needs to be deduced
			tok.Literal = l.readIdentifier()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}
