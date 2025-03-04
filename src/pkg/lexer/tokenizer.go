package lexer

import (
	"fmt"
	"priagram/src/pkg/id"
	"priagram/src/pkg/lexer/formatter"
	"strings"
)

// handle the source code and tokenize it
func Tokenize(source string) ([]Token, error) {
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

		if !matched {
			return nil, fmt.Errorf("Lexer::Error -> unrecognized token at position (%d) near %s", lex.pos, lex.remainingSourceCode())
		}
	}

	lex.push(newToken(EOF, "EOF"))
	return lex.Tokens, nil
}

func Format(tokens []Token) ([]formatter.DiagramData, []formatter.Relation) {
	var formatedDiagram []formatter.DiagramData
	var relations []formatter.Relation

	modelGroups := splitIntoModels(tokens)

	for _, group := range modelGroups {
		modelName, modelContents, modelRelations := processModel(group)

		relations = append(relations, modelRelations...)

		formatedDiagram = append(formatedDiagram, formatter.NewDiagramData(
			modelName,
			"prisma-table",
			modelContents,
		))
	}

	return formatedDiagram, relations
}

func splitIntoModels(tokens []Token) [][]Token {
	var groups [][]Token
	var currentGroup []Token
	inModel := false

	for _, tk := range tokens {
		if tk.Type == MODEL {
			inModel = true
			currentGroup = []Token{}
			continue
		}

		if inModel {
			currentGroup = append(currentGroup, tk)
		}

		if tk.Type == CLOSE_CURLY && inModel {
			groups = append(groups, currentGroup)
			inModel = false
		}
	}

	return groups
}

func processModel(tokens []Token) (string, []formatter.ModelContent, []formatter.Relation) {
	var modelName string
	var colNames, colTypes, colRelations []string

	for _, tk := range tokens {
		switch tk.Type {
		case MODEL_NAME:
			modelName = tk.Value
		case COLUMN_NAME:
			colNames = append(colNames, tk.Value)
		case COLUMN_TYPE:
			colTypes = append(colTypes, tk.Value)
		case RELATION:
			colRelations = append(colRelations, tk.Value)
		}
	}

	modelContents := createModelContents(colNames, colTypes)
	modelRelations := createRelations(modelName, colRelations)

	return modelName, modelContents, modelRelations
}

func createModelContents(colNames, colTypes []string) []formatter.ModelContent {
	var contents []formatter.ModelContent

	for i := 0; i < len(colNames); i++ {
		formatedId := id.FmtId(colNames[i], colTypes[i])
		mc := formatter.NewModelContent(formatedId, colNames[i], colTypes[i])
		contents = append(contents, mc)
	}

	return contents
}

func createRelations(modelName string, relations []string) []formatter.Relation {
	var modelRelations []formatter.Relation

	for _, relation := range relations {
		formatedId := id.FmtId(strings.ToLower(modelName), relation)
		newRelation := formatter.NewRelation(formatedId, strings.ToLower(relation), strings.ToLower(modelName))
		modelRelations = append(modelRelations, newRelation)
	}

	return modelRelations
}
