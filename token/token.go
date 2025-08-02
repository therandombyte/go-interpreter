package token

// ------------------------------------------------------------
// Background:
//   -----------        --------      ----------------------
//  | Source Code | -> | Tokens | -> | Abstract Syntax Tree |
//   ------------       --------      ----------------------
//   The first transformation will be done by a tokenizer (lexer)
//   And, Token will be a data structure/ struct that will be input to tokenizer

//   Token needs to define a type: integer, identifiers, keywords, punctuations/operators
// 		keywords     : if, let, return
// 		literals     : 123, "foo"
// 		identifiers  : myVar
// 		delimiters   : (,),{,}
// 		operators    : +,-,*
// ------------------------------------------------------------

// fixed values for TokenType
const (
	// reserved words/keywords
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers (we dont care if its int or variables) + literals
	IDENT = "IDENT" // add, foo, x, y. TokenType for all user-defined identifiers
	INT   = "INT"   // 123

	// operators
	ASSIGN = "="
	PLUS   = "+"

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
// In this case, we can have specific values for our TokenType
type TokenType string

// every token has a "valid" type and holds the current character
type Token struct {
	Type    TokenType
	Literal string
}

// helper function to deduce TokenType from literal
// keywords need to be separated from input literals
// keep all the keywords in a map, later iterate over it

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
