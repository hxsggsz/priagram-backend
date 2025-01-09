package lexer

const (
	EOF TokenType = iota
	NULL
)

var ReservedWords map[TokenType]string = map[TokenType]string{
	EOF:  "eof",
	NULL: "null",
}
