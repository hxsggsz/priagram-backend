package lexer

import (
	"regexp"
	"strings"
)

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

func modelNameHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainingSourceCode())
	modelName := strings.TrimSpace(lex.remainingSourceCode()[:match[1]-1])

	lex.advancePosition(len(modelName))
	lex.push(newToken(MODEL_NAME, modelName))
}

func columnEnumHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringSubmatch(lex.remainingSourceCode())

	lex.push(newToken(COLUMN_ENUM, match[0]))
	lex.advancePosition(len(match[0]))
}

func columnNameHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringSubmatch(lex.remainingSourceCode())

	lex.push(newToken(COLUMN_NAME, match[0]))
	lex.advancePosition(len(match[0]))
}

func columnTypeHandler(lex *lexer, regex *regexp.Regexp) {
	matches := regex.FindStringSubmatch(lex.remainingSourceCode())
	match := matches[0]

	lastToken := lex.getLastToken()
	isTypeEnum := lex.isTypeAEnum(match)

	if strings.Contains(match, lastToken.Value) {
		lex.advancePosition(len(match))
		lex.removeLastToken()

		return
	}

	lex.advancePosition(len(match))
	lex.push(newToken(COLUMN_TYPE, match))

	if isTypeEnum {
		lex.push(newToken(RELATION, match))
	}
}

func columnRelationHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringSubmatch(lex.remainingSourceCode())
	re := regexp.MustCompile(`\[([a-zA-Z]+?)(?:Id)?\]`)
	matches := re.FindStringSubmatch(match[0])

	lex.push(newToken(RELATION, matches[1]))
	lex.advancePosition(len(match[0]))
}
