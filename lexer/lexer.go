package lexer 

import (
	"es3/token"
)

type Lexer struct {
	input string
	position int 
	nextPosition int
	character byte
}
//New takes a string and returns a lexer.
func New(input string) Lexer {
	lex := Lexer{input: input}
	lex.setNextCharacter()
	return lex
}
 
func (lex *Lexer) setNextCharacter() {
	if lex.nextPosition >= len(lex.input) {
		lex.character = 0
	} else {
		lex.character = lex.input[lex.nextPosition]
	}
	lex.position = lex.nextPosition
	lex.nextPosition++
}

func (lex *Lexer) getNextCharacter() byte {
	if lex.nextPosition >= len(lex.input) {
		return 0
	}
	return lex.input[lex.nextPosition]
}

func (lex *Lexer) NextToken() token.Token {
	var tok token.Token
	switch lex.character {
	case '[':
		tok = newToken(token.LBRACK, lex.character)
	case ']': 
		tok = newToken(token.RBRACK, lex.character)
	case '{':
		tok = newToken(token.LBRACE, lex.character)
	case '}':
		tok = newToken(token.RBRACE, lex.character)
	case '(':
		tok = newToken(token.LPAREN, lex.character)
	case ')':
		tok = newToken(token.RPAREN, lex.character)
	case '.':
		tok = newToken(token.DOT, lex.character)
	case ';':
		tok = newToken(token.SEMICOLON, lex.character)
	case ':':
		tok = newToken(token.COLON, lex.character)
	case ',':
		tok = newToken(token.COMMA, lex.character)
	case '?':
		tok = newToken(token.QUESTION, lex.character)
	case '~':
		tok = newToken(token.TILDE, lex.character)
	case '%':
		if lex.getNextCharacter() == '=' {
			formerCharacter := lex.character 
			lex.setNextCharacter()
			tok = token.Token{Type: token.PERCENTASSIGN, Literal: string(formerCharacter)+ string(lex.character)}
		} else {
			tok = newToken(token.PERCENT, lex.character)
		}
	case '^':
		if lex.getNextCharacter() == '=' {
			formerCharacter := lex.character 
			lex.setNextCharacter()
			tok = token.Token{Type: token.CARATASSIGN, Literal: string(formerCharacter)+ string(lex.character)}
		} else {
			tok = newToken(token.CARAT, lex.character)
		}
	case '<':
		if lex.getNextCharacter() == '<' {
			formerCharacter := lex.character 
			lex.setNextCharacter()
			if lex.getNextCharacter() == '=' {
				latterCharacter := lex.character
				lex.setNextCharacter()
				tok = token.Token{Type: token.LTLTASSIGN, Literial: string(formerCharacter) + string(latterCharacter) + string(lex.character)}
			} else {
				tok = token.Token{Type: token.LTLT, Literal: string(formerCharacter)+ string(lex.character)}
			}
		} else if lex.getNextCharacter() == '=' {
			formerCharacter := lex.character 
			lex.setNextCharacter()
			tok = token.Token{Type: token.LTASSIGN, Literal: string(formerCharacter) + string(lex.character)}	
		} else {
			tok = newToken(token.LT, lex.character)
		}
	case '>':
		if lex.getNextCharacter() == '>' {
			formerCharacter := lex.character 
			lex.setNextCharacter()
			if lex.getNextCharacter() == '>' {
				latterCharacter := lex.character
				lex.setNextCharacter()
				if lex.getNextCharacter() == '=' {
					previousCharacter := lex.character
					lex.setNextCharacter()				
					tok = token.Token{Type: token.GTGTGTASSIGN, Literal: string(formerCharacter)+ string(latterCharacter) + string(previousCharacter) + string(lex.character) }
				} else {
					tok = token.Token{Type: token.GTGTGT, Literal: string(formerCharacter)+ string(latterCharacter) + string(lex.character)}
				}
			} else if lex.getNextCharacter() == '=' {
				latterCharacter := lex.character
				lex.setNextCharacter()
				tok = token.Token{Type: token.GTGTASSIGN, Literal: string(formerCharacter)+ string(latterCharacter) + string(lex.character)}
			} else {
				tok = token.Token{Type: token.GTGT, Literal: string(formerCharacter)+ string(lex.character)}
			}	
		} else if lex.getNextCharacter() == '=' {
			formerCharacter := lex.character 
			lex.setNextCharacter()
			tok = token.Token{Type: token.GTASSIGN, Literal: string(formerCharacter) + string(lex.character)}	
		} else {
			tok = newToken(token.GT, lex.character)
		}
	case '=':
		if lex.getNextCharacter() == '=' {
			formerCharacter := lex.character
			lex.setNextCharacter()
			if lex.getNextCharacter() == '=' {
				latterCharacter := lex.character 
				lex.setNextCharacter()
				tok = token.Token{Type: token.ASSIGNASSIGNASSIGN, Literal: string(formerCharacter)+ string(latterCharacter) + string(lex.character)}
			} else {
				tok = token.Token{Type: token.ASSIGNASSIGN, Literal: string(formerCharacter) + string(lex.character)}
			}
		} else {
			tok = newToken(token.ASSIGN, lex.character)
		}
	case '!':
		if lex.getNextCharacter() == '=' {
			formerCharacter := lex.character
			lex.setNextCharacter()
			if lex.getNextCharacter() == '=' {
				latterCharacter := lex.character 
				lex.setNextCharacter()
				tok = token.Token{Type: token.BANGASSIGNASSIGN, Literal: string(formerCharacter)+ string(latterCharacter) + string(lex.character)}
			} else {
				tok = token.Token{Type: token.BANGASSIGN, Literal: string(formerCharacter) + string(lex.character)}
			}
		} else {
			tok = newToken(token.BANG, lex.character)
		}
	case '|':
		if lex.getNextCharacter() == '=' {
			formerCharacter := lex.character
			lex.setNextCharacter()
			tok = token.Token{Type: token.ORASSIGN, Literal: string(formerCharacter) + string(lex.character)}
		} else if lex.getNextCharacter() == '|' {
			formerCharacter := lex.character
			lex.setNextCharacter()
			tok = token.Token{Type: token.OROR, Literal: string(formerCharacter) + string(lex.character)}
		} else {
			tok = newToken(token.OR, lex.character)
		}
	case '&':
		if lex.getNextCharacter() == '=' {
			formerCharacter := lex.character
			lex.setNextCharacter()
			tok = token.Token{Type: token.ANDASSIGN, Literal: string(formerCharacter) + string(lex.character)}
		} else if lex.getNextCharacter() == '&' {
			formerCharacter := lex.character
			lex.setNextCharacter()
			tok = token.Token{Type: token.ANDAND, Literal: string(formerCharacter) + string(lex.character)}
		} else {
			tok = newToken(token.AND, lex.character)
		}
	case '+':
		if lex.getNextCharacter() == '=' {
			formerCharacter := lex.character
			lex.setNextCharacter()
			tok = token.Token{Type: token.PLUSASSIGN, Literal: string(formerCharacter) + string(lex.character)}
		} else if lex.getNextCharacter() == '+' {
			formerCharacter := lex.character
			lex.setNextCharacter()
			tok = token.Token{Type: token.PLUSPLUS, Literal: string(formerCharacter) + string(lex.character)}
		} else {
			tok = newToken(token.PLUS, lex.character)
		}
	case '-':
		if lex.getNextCharacter() == '=' {
			formerCharacter := lex.character
			lex.setNextCharacter()
			tok = token.Token{Type: token.MINUSASSIGN, Literal: string(formerCharacter) + string(lex.character)}
		} else if lex.getNextCharacter() == '-' {
			formerCharacter := lex.character
			lex.setNextCharacter()
			tok = token.Token{Type: token.MINUSMINUS, Literal: string(formerCharacter) + string(lex.character)}
		} else {
			tok = newToken(token.MINUS, lex.character)
		}
	case '*':
		if lex.getNextCharacter() == '=' {
			formerCharacter := lex.character
			lex.setNextCharacter()
			tok = token.Token{Type: token.STARASSIGN, Literal: string(formerCharacter) + string(lex.character)}
		} else {
			tok = newToken(token.STAR, lex.character)
		}
	case '/':
		if lex.getNextCharacter() == '=' {
			formerCharacter := lex.character
			lex.setNextCharacter()
			tok = token.Token{Type: token.DIVASSIGN, Literal: string(formerCharacter) + string(lex.character)}
		} else {
			tok = newToken(token.DIV, lex.character)
		}
	}
	lex.setNextCharacter()
	return tok
}

func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}