package lexer

import (
	"fmt"
	"priagram/src/pkg/id"
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
	var modelRelations []formatter.Relation
	var modelName string
	var colName, colType, colRelations []string

	for _, tk := range tokens {
		if tk.isOneOfMany(MODEL, EOF, OPEN_CURLY, CLOSE_CURLY) {
			continue
		}

		switch tk.Type {
		case MODEL_NAME:
			modelName = tk.Value

		case COLUMN_TYPE:
			colType = append(colType, tk.Value)

		case COLUMN_NAME:
			colName = append(colName, tk.Value)

		case RELATION:
			colRelations = append(colRelations, tk.Value)
		}
	}

	for i := 0; i < len(colName); i++ {
		formatedId := id.FmtId(colName[i], colType[i])
		mc := formatter.NewModelContent(formatedId, colName[i], colType[i])
		modelContents = append(modelContents, mc)
	}

	for _, relation := range colRelations {
		formatedId := id.FmtId(modelName, relation)
		newRelation := formatter.NewRelation(formatedId, modelName, relation)
		modelRelations = append(modelRelations, newRelation)
	}

	return formatter.NewDiagramData(modelName, "prisma-table", modelContents, modelRelations)
}
