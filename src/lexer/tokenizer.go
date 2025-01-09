package lexer

import "fmt"

// handle the source code and tokenize it
func Tokenize(source string) []Token {
	lex := createLexer(source)

	for !lex.atEof() {
		matched := false

		for _, pattern := range lex.patterns {
			loc := pattern.regex.FindStringIndex(lex.remainingSourceCode())

			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}

		// TODO: update this to print the location and the line it happened
		if !matched {
			panic(fmt.Sprintf("Lexer::Error -> unrecognized token near %s\n", lex.remainingSourceCode()))
		}
	}

	lex.push(newToken(EOF, "EOF"))
	return lex.Tokens
}
