package lexer

import (
	"github.com/croese/minic-go/token"
)

type Lexer struct {
	input        string
	position     int // points to current char in input
	readPosition int // current reading position (after current char)
	ch           byte
	lineNumber   int
	column       int
}

func New(input string) *Lexer {
	l := &Lexer{
		input:      input,
		lineNumber: 1,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
	l.column += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			tok = l.twoCharacterToken(token.EQ)
		} else {
			tok = l.newToken(token.ASSIGN)
		}
	case ';':
		tok = l.newToken(token.SEMICOLON)
	case '(':
		tok = l.newToken(token.LPAREN)
	case ')':
		tok = l.newToken(token.RPAREN)
	case '{':
		tok = l.newToken(token.LBRACE)
	case '}':
		tok = l.newToken(token.RBRACE)
	case '[':
		tok = l.newToken(token.LBRACKET)
	case ']':
		tok = l.newToken(token.RBRACKET)
	case ',':
		tok = l.newToken(token.COMMA)
	case '!':
		if l.peekChar() == '=' {
			tok = l.twoCharacterToken(token.NOT_EQ)
		} else {
			tok = l.newToken(token.BANG)
		}
	case '+':
		tok = l.newToken(token.PLUS)
	case '-':
		tok = l.newToken(token.MINUS)
	case '/':
		tok = l.newToken(token.SLASH)
	case '*':
		tok = l.newToken(token.ASTERISK)
	case '%':
		tok = l.newToken(token.PERCENT)
	case '<':
		tok = l.newToken(token.LT)
	case '>':
		tok = l.newToken(token.GT)
	case 0:
		tok = l.eofToken()
	default:
		if isLetter(l.ch) {
			return l.identiferToken()
		} else if isDigit(l.ch) {
			return l.numberToken()
		} else {
			tok = l.newToken(token.ILLEGAL)
		}
	}

	l.readChar()
	return tok
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.lineNumber += 1
			l.column = 0
		}
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) newToken(tokenType token.TokenType) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(l.ch),
		Line:    l.lineNumber,
		Column:  l.column,
	}
}

func (l *Lexer) eofToken() token.Token {
	return token.Token{
		Type:    token.EOF,
		Literal: "",
		Line:    l.lineNumber,
		Column:  l.column,
	}
}

func (l *Lexer) identiferToken() token.Token {
	col := l.column
	lit := l.readIdentifier()

	return token.Token{
		Type:    token.LookupIdent(lit),
		Literal: lit,
		Line:    l.lineNumber,
		Column:  col,
	}
}

func (l *Lexer) numberToken() token.Token {
	col := l.column
	lit := l.readNumber()

	return token.Token{
		Type:    token.INT,
		Literal: lit,
		Line:    l.lineNumber,
		Column:  col,
	}
}

func (l *Lexer) twoCharacterToken(tokenType token.TokenType) token.Token {
	ch, col := l.ch, l.column
	l.readChar()
	literal := string(ch) + string(l.ch)
	return token.Token{
		Type:    tokenType,
		Literal: literal,
		Line:    l.lineNumber,
		Column:  col,
	}
}
