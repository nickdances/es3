package lexer 

import (
	"testing"

	"es3/token"
)

func TestNextToken(t *testing.T) {
	input := `var x = 2
	var y = 3
	var _z = (x + y)
	function (Param1, $param2) {
		var array = [Param1, $param2];
		if (true) {
			Param1
		} else {
			return array[0] <<= "'4'" || 1 > 2
		}
	}
	return _z >>>= 15.9`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.IDENTIFIER, "x"},
		{token.ASSIGN, "="},
		{token.NUMBER, "2"},
		{token.NEWLINE, "\n"},
		{token.VAR, "var"},
		{token.IDENTIFIER, "y"},
		{token.ASSIGN, "="},
		{token.NUMBER, "3"},
		{token.NEWLINE, "\n"},
		{token.VAR, "var"},
		{token.IDENTIFIER, "_z"},
		{token.ASSIGN, "="},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.NEWLINE, "\n"},
		{token.FUNCTION, "function"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "Param1"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "$param2"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.NEWLINE, "\n"},
		{token.VAR, "var"},
		{token.IDENTIFIER, "array"},
		{token.ASSIGN, "="},
		{token.LBRACK, "["},
		{token.IDENTIFIER, "Param1"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "$param2"},
		{token.RBRACK, "]"},
		{token.SEMICOLON, ";"},		
		{token.NEWLINE, "\n"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.TRUE, "true"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.NEWLINE, "\n"},
		{token.IDENTIFIER, "Param1"},
		{token.NEWLINE, "\n"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.NEWLINE, "\n"},
		{token.RETURN, "return"},
		{token.IDENTIFIER, "array"},
		{token.LBRACK, "["},
		{token.NUMBER, "0"},
		{token.RBRACK, "]"},
		{token.LTLTASSIGN, "<<="},
		{token.QUOTEQUOTE, "\""},
		{token.QUOTE, "'"},
		{token.NUMBER, "4"},
		{token.QUOTE, "'"},
		{token.QUOTEQUOTE, "\""},
		{token.OROR, "||"},
		{token.NUMBER, "1"},
		{token.GT, ">"},
		{token.NUMBER, "2"},
		{token.NEWLINE, "\n"},
		{token.RBRACE, "}"},
		{token.NEWLINE, "\n"},
		{token.RBRACE, "}"},
		{token.NEWLINE, "\n"},
		{token.RETURN, "return"},
		{token.IDENTIFIER, "_z"},
		{token.GTGTGTASSIGN, ">>>="},
		{token.NUMBER, "15"},
		{token.DOT, "."},
		{token.NUMBER, "9"},
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