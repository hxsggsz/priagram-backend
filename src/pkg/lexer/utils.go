package lexer

import "regexp"

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

func (lex *lexer) removeLastToken() {
	tokens := lex.Tokens

	if len(tokens) > 0 {
		lex.Tokens = tokens[:len(tokens)-1]
	}
}

// gets the remaining source code attached to the lexer struct from the current position to the end with the :
//
// Return:
//   - string: the remaining code according by the current lexer position
func (lex *lexer) remainingSourceCode() string {
	return lex.source[lex.pos:]
}

func (lex *lexer) getLastToken() Token {
	return lex.Tokens[len(lex.Tokens)-1]
}
func (token Token) IsOneOfMany(expectedTokens ...TokenType) bool {
	for _, expected := range expectedTokens {
		return expected == token.Type
	}

	return false
}

func (lex *lexer) atEof() bool {
	return lex.pos >= len(lex.source)
}

func createLexer(source string) *lexer {
	return &lexer{
		pos:    0,
		source: source,
		Tokens: make([]Token, 0),
		patterns: []regexPattern{
			{regexp.MustCompile(`\s+`), skipHandler},                // skips whitespaces
			{regexp.MustCompile(`@\w+(?:\([^)]*\))?`), skipHandler}, // skips `@id` or `@unique`
			{regexp.MustCompile(`\)`), skipHandler},                 // skips `@id` or `@unique`
			{regexp.MustCompile(`\{`), defaultHandler(OPEN_CURLY, "{")},
			{regexp.MustCompile(`\}`), defaultHandler(CLOSE_CURLY, "}")},
			{regexp.MustCompile(`\bmodel\b`), defaultHandler(MODEL, "model")},
			{regexp.MustCompile(`@[^)]+?\)`), columnRelationHandler},
			{regexp.MustCompile(`(.*?)\s*{`), modelNameHandler},
			{regexp.MustCompile(`^[A-Z][a-zA-Z]*(\?|(\[\]))?`), columnTypeHandler},
			{regexp.MustCompile(`^[a-zA-Z]\S*`), columnNameHandler},
		},
	}
}
