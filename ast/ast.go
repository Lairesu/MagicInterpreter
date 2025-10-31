package ast

import (
	"MagicInterpreter/token"
)

// ========================
// Node, Statement and Expression Interface
// ========================

// all ast node  implement the node interface
// node represents a node in the abstract syntax tre
type Node interface {
	TokenLiteral() string
}

// represent the statement node, does not produce value, all statement must apply this node
type Statement interface {
	Node
	StatementNode()
}

// represents the expression node, produce value and can be evaluated, must apply this node
type Expression interface {
	Node
	ExpressionNode()
}

// ======================
// program node
// ======================

// root of all ast , implements statement interface, contains every parsed statement
type Program struct {
	Statements []Statement
}

// helper function , return literal string of the token, return first statement's token literal if not ""
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// ======================
// LET STATEMENT
// ======================

// represents:  let <identifier> = <expression>
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// satisfy the Statement interface
func (ls *LetStatement) StatementNode() {}

// returns the literal value of the `let` token
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// ======================
//	IDENTIFIER
// ======================

type Identifier struct {
	Token token.Token
	Value string
}

// satisfy the Expression interface
func (i *Identifier) ExpressionNode() {}

// returns literal value of the identifier token
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
