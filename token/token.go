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
	// TokenType for all user-defined identifiers
	IDENT = "IDENT" // add, foo, x, y.
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

	// keywords (need to be looked up from identifiers)
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// TokenType holds all the allowed "type" values for a token.
type TokenType string

// Best Practice:
// Named types are for type safety (explicit conversion), readability, extensibility (by attaching methods/values)
// Also, switch statements can be applied on the type values
// In this case, we can have enumerated values for our TokenType

// Token holds the current literal/character and its type
type Token struct {
	Type    TokenType
	Literal string
}

// Map of all language keywords and its token type,
// to be used in lookup
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdentifier will check if the input identifier is a language keyword,
// and return its token type.
// Otherwise return IDENT (default for all identifiers)
func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
