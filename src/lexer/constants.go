package lexer

const (
	EOF TokenType = iota
	OPEN_CURLY
	CLOSE_CURLY
)

var ReservedWords map[TokenType]string = map[TokenType]string{
	EOF:         "eof",
	CLOSE_CURLY: "close curly",
	OPEN_CURLY:  "open curly",
}
