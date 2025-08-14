package parser

import (
	"github.com/therandombyte/go-interpreter/ast"
	"github.com/therandombyte/go-interpreter/lexer"
	"github.com/therandombyte/go-interpreter/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// Parser repeatedly advances the tokens and checks the current token to decide 
// what to do next: either call another parsing function or throw an error. 
// Each function then does its job and possibly constructs an AST node so that 
// the “main loop” in parseProgram() can advance the tokens and decide what to do again
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	return nil
}
