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

// TODO: create the handlers for the model, the model name and etc here

func modelNameHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainingSourceCode())
	modelName := strings.TrimSpace(lex.remainingSourceCode()[:match[1]-1])

	err := validateModelName(modelName)

	if err != nil {
		panic(err)
	}

	lex.advancePosition(len(modelName))
	lex.push(newToken(MODEL_NAME, modelName))
}

func columnHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringSubmatch(lex.remainingSourceCode())
	// TODO: get the column name and type here

	lex.advancePosition(len(match[0]))
	lex.push(newToken(ID, "id"))
}
