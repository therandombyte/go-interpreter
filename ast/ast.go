package ast

// ------------------------------------------------------------
// Background:
// After the lexer produces tokens, the Parser will take those
// and convert to a data structure called Abstract Syntax Tree (AST)
// Implementation-wise, it will be a Tree made of parsed tokens
// called as Nodes.

// Language parsing needs two nodes: expressions and statements
// ex: let x = 10
//     let <identifier> = <expression>
// ------------------------------------------------------------

// Node will be the building block of AST.
// Every node will return its associated token literal.
// These nodes will be connected to form a tree.
// Some nodes will implement Statement, others Expresion.
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program will be the root Node of AST.
// Every program is a series of statements,
// which in turn is a series of Nodes that implement Statement.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
