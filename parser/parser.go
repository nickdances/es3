package parser

import (
	"fmt"
	
	"es3/ast"
	"es3/lexer"
	"es3/token"
)

type Parser struct {
	lex *lexer.Lexer
	errors []string

	curToken token.Token
	nextToken token.Token
}

func New (lex *lexer.Lexer) *Parser {
	parse := &Parser{
		lex: lex,
		errors: []string{},	
	}

	parse.setNextToken()
	parse.setNextToken()

	return parse
}

func (parse *Parser) Errors() []string {
	return parse.errors
}

func (parse *Parser) setError(t token.TokenType) {
	msg := fmt.Sprintf("want token %s, have token %s", t, parse.nextToken.Type)
	parse.errors = append(parse.errors, msg)
}

func (parse *Parser) setNextToken() {
	parse.curToken = parse.nextToken
	parse.nextToken = parse.lex.NextToken()
}

func (parse *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !parse.curTokenIs(token.EOF) {
		stmt := parse.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		parse.setNextToken()
	}

	return program
}

func (parse *Parser) parseStatement() ast.Statement {
	switch parse.curToken.Type{
	case token.VAR:
		return parse.parseVarStatement()
	default:
		return nil
	}
}

func (parse *Parser) parseVarStatement() *ast.VarStatement {
	stmt := &ast.VarStatement{Token: parse.curToken}

	if !parse.nextType(token.IDENTIFIER) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: parse.curToken, Value: parse.curToken.Literal}

	if !parse.nextType(token.ASSIGN) {
		return nil
	}

	for !parse.curTokenIs(token.SEMICOLON) && !parse.curTokenIs(token.LINEBREAK)  {
		parse.setNextToken()
	}

	return stmt

}

func (parse *Parser) curTokenIs(t token.TokenType) bool {
	return parse.curToken.Type == t
}

func (parse *Parser) nextTokenIs(t token.TokenType) bool {
	return parse.nextToken.Type == t
}

func (parse *Parser) nextType(t token.TokenType) bool {
	if parse.nextTokenIs(t) {
		parse.setNextToken()
		return true
	} else {
		parse.setError(t)
		return false
	}
}
