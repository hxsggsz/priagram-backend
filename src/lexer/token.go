package lexer

import (
	"fmt"
)

type Token struct {
	Type  TokenType
	Value string
}

func newUniqueToken(tp TokenType, value string) Token {
	return Token{
		tp, value,
	}
}

func (token Token) Debug() {
	fmt.Printf("%s()\n", formatTokenToString(token.Type))
}

func formatTokenToString(tp TokenType) string {
	if value, found := ReservedWords[tp]; found {
		return value
	}

	return fmt.Sprintf("unknown(%d)", tp)
}
