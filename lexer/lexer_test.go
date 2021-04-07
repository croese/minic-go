package lexer

import (
	"testing"

	"github.com/croese/minic-go/token"
	"github.com/stretchr/testify/require"
)

func TestNextToken(t *testing.T) {
	input := `=(){};,+-*/%<>![]`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		expectedLine    int
		expectedColumn  int
	}{
		{token.ASSIGN, "=", 1, 1},
		{token.LPAREN, "(", 1, 2},
		{token.RPAREN, ")", 1, 3},
		{token.LBRACE, "{", 1, 4},
		{token.RBRACE, "}", 1, 5},
		{token.SEMICOLON, ";", 1, 6},
		{token.COMMA, ",", 1, 7},
		{token.PLUS, "+", 1, 8},
		{token.MINUS, "-", 1, 9},
		{token.ASTERISK, "*", 1, 10},
		{token.SLASH, "/", 1, 11},
		{token.PERCENT, "%", 1, 12},
		{token.LT, "<", 1, 13},
		{token.GT, ">", 1, 14},
		{token.BANG, "!", 1, 15},
		{token.LBRACKET, "[", 1, 16},
		{token.RBRACKET, "]", 1, 17},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		require.Equal(t, tt.expectedType, tok.Type, "in tests[%d], tok=%+v", i, tok)
		require.Equal(t, tt.expectedLiteral, tok.Literal, "in tests[%d], tok=%+v", i, tok)
		require.Equal(t, tt.expectedLine, tok.Line, "in tests[%d], tok=%+v", i, tok)
		require.Equal(t, tt.expectedColumn, tok.Column, "in tests[%d], tok=%+v", i, tok)
	}
}
