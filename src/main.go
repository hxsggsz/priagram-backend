package main

import (
	"fmt"
	"os"
	"priagram/src/lexer"
)

func main() {
	bytes, err := os.ReadFile("./examples/test.lang")
	if err != nil {
		panic(fmt.Sprintf("Main::Error -> failed to read file %s\n", err))
	}

	tokens := lexer.Tokenize(string(bytes))
	for _, token := range tokens {
		token.Debug()
	}
}
