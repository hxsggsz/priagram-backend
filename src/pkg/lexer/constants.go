package lexer

const (
	EOF TokenType = iota
	OPEN_CURLY
	CLOSE_CURLY
	MODEL
	ENUM
	MODEL_NAME
	COLUMN_ENUM
	COLUMN_NAME
	COLUMN_TYPE
	RELATION
)

var ReservedWords map[TokenType]string = map[TokenType]string{
	EOF:         "eof",
	CLOSE_CURLY: "close curly",
	OPEN_CURLY:  "open curly",
	ENUM:        "enum",
	MODEL:       "model",
	MODEL_NAME:  "model name",
	COLUMN_ENUM: "column enum",
	COLUMN_NAME: "column name",
	COLUMN_TYPE: "column type",
	RELATION:    "realtion",
}
