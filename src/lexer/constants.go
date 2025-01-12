package lexer

const (
	EOF TokenType = iota
	OPEN_CURLY
	CLOSE_CURLY
	MODEL
	MODEL_NAME
	COLUMN_NAME
	COLUMN_TYPE
)

var ReservedWords map[TokenType]string = map[TokenType]string{
	EOF:         "eof",
	CLOSE_CURLY: "close curly",
	OPEN_CURLY:  "open curly",
	MODEL:       "model",
	MODEL_NAME:  "model name",
	COLUMN_NAME: "column name",
	COLUMN_TYPE: "column type",
}
