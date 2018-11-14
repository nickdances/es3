package lexer 

import (
	"testing"

	"es3/token"
)

func TestNextToken(t *testing.T) {
	input := `===[]=={}().;,<>=!?|&+-!=!==`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGNASSIGNASSIGN, "==="},
		{token.LBRACK, "["},
		{token.RBRACK, "]"},
		{token.ASSIGNASSIGN, "=="},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.DOT, "."},
		{token.SEMICOLON, ";"},
		{token.COMMA, ","},
		{token.LT, "<"},
		{token.GTASSIGN, ">="},
		{token.BANG, "!"},
		{token.QUESTION, "?"},
		{token.OR, "|"},
		{token.AND, "&"},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.BANGASSIGN, "!="},
		{token.BANGASSIGNASSIGN, "!=="},
	}
	
	l := New(input)

	for i, testToken := range tests {
		tok := l.NextToken()

		if tok.Type != testToken.expectedType {
			t.Fatalf("test %d expected %q got %q", i, testToken.expectedType, tok.Type)
		}
		if tok.Literal != testToken.expectedLiteral {
			t.Fatalf("test %d expected %q got %q", i, testToken.expectedLiteral, tok.Literal)
		}
	}
}

// const (
// 	ILLEGAL = "ILLEGAL"
// 	EOF = "EOF"
// 	//Punctuators
// 	LBRACE = "{"
// 	RBRACE = "}"
// 	LPAREN = "("
// 	RPAREN = ")"
// 	LBRACK = "["
// 	RBRACK = "]"
// 	DOT = "."
// 	SEMICOLON = ";"
// 	COMMA = ","
// 	LT = "<"
// 	GT = ">"
// 	LTASSIGN = "<="
// 	GTASSIGN = ">="
// 	ASSIGN = "="
// 	ASSIGNASSIGN = "=="
// 	ASSIGNASSIGNASSIGN = "==="
// 	BANGASSIGN = "!="
// 	BANGASSIGNASSIGN = "!=="
// 	PLUS = "+"
// 	MINUS = "-"
// 	STAR = "*"
// 	PERCENT = "%"
// 	PLUSPLUS = "++"
// 	MINUSMINUS = "--"
// 	LTLT = "<<"
// 	GTGT = ">>"
// 	GTGTGT = ">>>"
// 	AND = "&"
// 	OR = "|"
// 	CARAT = "^"
// 	BANG = "!"
// 	TILDE = "~"
// 	ANDAND = "&&"
// 	OROR = "||"
// 	QUESTION = "?"
// 	COLON = ":"
// 	ASSIGN = "="
// 	PLUSASSIGN = "+="
// 	MINUSASSIGN = "-="
// 	STARASSIGN = "*="
// 	PERCENTASSIGN = "%="
// 	LTLTASSIGN = "<<="
// 	GTGTASSIGN = ">>="
// 	GTGTGTASSIGN = ">>>="
// 	ANDASSIGN = "&="
// 	ORASSIGN = "|="
// 	CARATASSIGN = "^="
// 	//DivPunctuators
// 	DIV = "/"
// 	DIVASSIGN = "/="
// 	//Keywords
// 	BREAK = "BREAK"
// 	CASE = "CASE"
// 	CATCH = "CATCH"
// 	CONTINUE = "CONTINUE"
// 	DEFAULT = "DEFAULT"
// 	DELETE = "DELETE"
// 	DO = "DO"
// 	ELSE = "ELSE"
// 	FINALLY = "FINALLY"
// 	FOR = "FOR"
// 	FUNCTION = "FUNCTION"
// 	IF = "IF"
// 	IN = "IN"
// 	INSTANCEOF = "INSTANCEOF"
// 	NEW = "NEW"
// 	RETURN = "RETURN"
// 	SWITCH = "SWITCH"
// 	THIS = "THIS"
// 	THROW = "THROW"
// 	TRY = "TRY"
// 	TYPEOF = "TYPEOF"
// 	VAR = "VAR"
// 	VOID = "VOID"
// 	WHILE = "WHILE"
// 	WITH = "WITH"
// 	//Future reserved words
// 	ABSTRACT = "ABSTRACT"
// 	BOOLEAN = "BOOLEAN"
// 	BYTE = "BYTE"
// 	CHAR = "CHAR"
// 	CLASS = "CLASS"
// 	CONST = "CONST"
// 	DEBUGGER = "DEBUGGER"
// 	DOUBLE = "DOUBLE"
// 	ENUM = "ENUM"
// 	EXPORT = "EXPORT"
// 	EXTENDS = "EXTENDS"
// 	FINAL = "FINAL"
// 	FLOAT = "FLOAT"
// 	GOTO = "GOTO"
// 	IMPLEMENTS = "IMPLEMENTS"
// 	IMPORT = "IMPORT"
// 	INT = "INT"
// 	INTERFACE = "INTERFACE"
// 	LONG = "LONG"
// 	NATIVE = "NATIVE"
// 	PACKAGE = "PACKAGE"
// 	PRIVATE = "PRIVATE"
// 	PROTECTED = "PROTECTED"
// 	PUBLIC = "PUBLIC"
// 	SHORT = "SHORT"
// 	STATIC = "STATIC"
// 	SUPER = "SUPER"
// 	SYNCHRONIZED = "SYNCHRONIZED"
// 	THROWS = "THROWS"
// 	TRANSIENT = "TRANSIENT"
// 	VOLATILE = "VOLATILE"
// 	//Literals 
// 	TRUE = "TRUE"
// 	FALSE = "FALSE"
// 	NULL = "NULL"
// 	UNDEFINED = "UNDEFINED"
// 	NUMBER = "NUMBER"
// 	STRING = "STRING"
// )