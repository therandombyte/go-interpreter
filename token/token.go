package token

// Background:
//   Source Code -> Tokens -> Abstract Syntax Tree
//   The first transformation will be done by a tokenizer (lexer)
//   And, Token will be a data structure/ struct that will be input to tokenizer

//   Token needs to define a type: integer, identifiers, keywords, punctuations/operators
// 		keywords     : if, let, return
// 		literals     : 123, "foo"
// 		identifiers  : myVar
// 		delimiters   : (,),{,}
// 		operators    : +,-,*


const (
	// reserved words/keywords
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers (we dont care its int or variables) + literals
	IDENT   = "IDENT" // add, foo, x, y
	INT     = "INT"   // 123

	// operators
	ASSIGN  = "="
	PLUS    = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// Named type:
// It's for type safety (explicit conversion), readability, extensibility (by attaching methods/values)
// Also, switch statements can be applied on the type values
// In this case, we can have specific values for out TokenType
type TokenType string



type Token struct {
	Type TokenType
	Literal string
}


