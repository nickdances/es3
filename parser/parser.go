package parser

import (
	"es3/ast"
	"es3/lexer"
	"es3/token"
)

type Parser struct {
	lex *lexer.Lexer

	curToken token.Token
	nextToken token.Token
}

func new (lex *lexer.Lexer) *Parser {

}
