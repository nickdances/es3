package lexer 

import (
	"testing"

	"github.com/nicholasgriffen/es3/token"
)

func TestNextToken(t *testing.T) {
	input := `var x = 2
	var 	y = 3
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
		{token.LINEBREAK, "LINEBREAK"},
		{token.VAR, "var"},
		{token.IDENTIFIER, "y"},
		{token.ASSIGN, "="},
		{token.NUMBER, "3"},
		{token.LINEBREAK, "LINEBREAK"},
		{token.VAR, "var"},
		{token.IDENTIFIER, "_z"},
		{token.ASSIGN, "="},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.LINEBREAK, "LINEBREAK"},
		{token.FUNCTION, "function"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "Param1"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "$param2"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.LINEBREAK, "LINEBREAK"},
		{token.VAR, "var"},
		{token.IDENTIFIER, "array"},
		{token.ASSIGN, "="},
		{token.LBRACK, "["},
		{token.IDENTIFIER, "Param1"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "$param2"},
		{token.RBRACK, "]"},
		{token.SEMICOLON, ";"},		
		{token.LINEBREAK, "LINEBREAK"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.TRUE, "true"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.LINEBREAK, "LINEBREAK"},
		{token.IDENTIFIER, "Param1"},
		{token.LINEBREAK, "LINEBREAK"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.LINEBREAK, "LINEBREAK"},
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
		{token.LINEBREAK, "LINEBREAK"},
		{token.RBRACE, "}"},
		{token.LINEBREAK, "LINEBREAK"},
		{token.RBRACE, "}"},
		{token.LINEBREAK, "LINEBREAK"},
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
