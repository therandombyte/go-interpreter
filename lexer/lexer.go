package lexer

import "go/token"

// Background:
// Source Code -> Lexer -> Token
// Instead of reading the source code from a file or shell,
// we just give it a string to avoid complexity

type Lexer struct {
	input string
	position int
	readPosition int
	ch byte
}

func New(input string) *Lexer {
	return &Lexer{input: input}
}

func NextToken() token.Token {
	var tok token.Token

	return tok
}
