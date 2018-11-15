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

func New (lex *lexer.Lexer) *Parser {
	parse := &Parser{lex: lex}

	parse.setNextToken()
	parse.setNextToken()

	return parse
}

func (parse *Parser) setNextToken() {
	parse.curToken = parse.nextToken
	parse.nextToken = parse.lex.NextToken()
}

func (parse *Parser) ParseProgram() *ast.Program {
	return nil
}
