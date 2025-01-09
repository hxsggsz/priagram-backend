package lexer

import (
	"fmt"
	"regexp"
)

type regexHandler func(lex *lexer, regex *regexp.Regexp)

type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

type lexer struct {
	patterns []regexPattern
	Tokens   []Token
	source   string
	pos      int
}

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

// advance the position of the code that it's currently tokenizing in N positions
// it takes one integer as a parameter
//   - n: integer - the number of positions to advance
func (lex *lexer) advancePosition(n int) {
	lex.pos += n
}

// add a new token to the end of the slice of tokens
// it takes a Token struct as a parameter
//   - token: Token - the token to add to the slice
func (lex *lexer) push(token Token) {
	lex.Tokens = append(lex.Tokens, token)
}

// gets the remaining source code attached to the lexer struct from the current position to the end with the :
//
// Return:
//   - string: the remaining code according by the current lexer position
func (lex *lexer) remainingSourceCode() string {
	return lex.source[lex.pos:]
}

func (lex *lexer) atEof() bool {
	return lex.pos >= len(lex.source)
}

// the default handler for simple tokens like { and }
// it takes the token type and the value of the code as paramters
//   - tp: TokenType - the type of the token by the enum
//   - value: string - the value of the code that it's currently executing
//
// Return:
//   - regexHandler: a function to handle the token
func defaultHandler(tp TokenType, value string) regexHandler {
	return func(lex *lexer, regex *regexp.Regexp) {
		lex.advancePosition(len(value))
		lex.push(newToken(tp, value))
	}
}
func skipHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainingSourceCode())
	lex.advancePosition(match[1])
}

//TODO: create the handlers for the model, the model name and etc here

func createLexer(source string) *lexer {
	return &lexer{
		pos:    0,
		source: source,
		Tokens: make([]Token, 0),
		patterns: []regexPattern{
			{regexp.MustCompile(`\{`), defaultHandler(OPEN_CURLY, "{")},
			{regexp.MustCompile(`\}`), defaultHandler(CLOSE_CURLY, "}")},
			{regexp.MustCompile(`\s+`), skipHandler}, // skips whitespaces
		},
	}
}
