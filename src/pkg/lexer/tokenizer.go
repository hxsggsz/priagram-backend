package lexer

import (
	"fmt"
	"priagram/src/pkg/lexer/formatter"
)

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
			panic(fmt.Sprintf("Lexer::Error -> unrecognized token at position (%d) near %s\n", lex.pos, lex.remainingSourceCode()))
		}
	}

	lex.push(newToken(EOF, "EOF"))
	return lex.Tokens
}

func Format(tokens []Token) formatter.DiagramData {
	var modelContents []formatter.ModelContent
	var modelName, colRelation string
	var colName, colType []string

	for _, tk := range tokens {
		if tk.isOneOfMany(MODEL, EOF, OPEN_CURLY, CLOSE_CURLY) {
			continue
		}

		fmt.Println("current token:", tk.Type, tk.Value)

		switch tk.Type {
		case MODEL_NAME:
			fmt.Println("current model name: ", tk.Type, tk.Value)
			modelName = tk.Value

		case COLUMN_TYPE:
			fmt.Println("current column type: ", tk.Type, tk.Value)
			colType = append(colType, tk.Value)

		case COLUMN_NAME:
			fmt.Println("current column name: ", tk.Type, tk.Value)
			colName = append(colName, tk.Value)

		case RELATION:
			fmt.Println("current column name: ", tk.Type, tk.Value)
			colRelation = tk.Value
		}
	}

	for i := 0; i < len(colName); i++ {
		mc := formatter.NewModelContent(colName[i], colName[i], colType[i])
		modelContents = append(modelContents, mc)
	}

	return formatter.NewDiagramData(modelName, "prisma-table", modelContents)
}
