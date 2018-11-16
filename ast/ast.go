package ast

import (
	"es3/token"
)
//Node provides a method to return token literals.
type Node interface {
	//TokenLiteral will typically return Token{Literal: string}
	TokenLiteral() string
}
//Statement 
type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}
//TokenLiteral method to satisfy Node interface
func (program *Program) TokenLiteral() string {
	if len(program.Statements) > 0 {
		return program.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type VarStatement struct {
	Token token.Token //token.VAR
	Name *Identifier
	Value Expression
}

func (vs *VarStatement) statementNode() {}
func (vs *VarStatement) TokenLiteral() string { return vs.Token.Literal }

type Identifier struct {
	Token token.Token // token.IDENTIFIER
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
