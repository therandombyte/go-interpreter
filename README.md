# Interpreter.... in 2025

This project is a re-write of Thorsen Ball's "Writing an Interpreter in Go". Much thanks to author for spreading the love.

In an ideal world, programming in plain text would have been sweet, but it has been tried and failed in the past.

So, the way forward is to break the code down into pieces (Lexing) and then forming a Tree structure (Parsing) out of it, to connect them all.  

This would be the result of our interpreter:

```
//   -----------        --------      ----------------------
//  | Source Code | -> | Tokens | -> | Abstract Syntax Tree |
//   ------------       --------      ----------------------
//	           (LEXER)        (PARSER)
```

The first conversion is called "Lexing" done by a lexer that we will write. The Lexer will break down the input source code into Tokens that will then be fed into the Parser.
Parser will in-turn create an Abstract Tree structure from the tokens known as AST.

## LEXING:

A simple code will be "let x = 5" which can also be represented as "let <identifier> = <expression>"
Varible names is normally called as "identifiers", and pre-defined strings like "let" and "fn" will be categorized as "kewords". Identifiers can be grouped together, however keywords cannot, as each of them mean a different logic.

1) Break input program code as Tokens. Tokens hold each character of the program code and its type.
2) During this phase, it does not matter if the code makes any sense

## PARSING
