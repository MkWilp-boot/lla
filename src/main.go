package main

import (
	"fmt"
	"lla/pkg/lexer"
	"lla/pkg/types"
	"os"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}()

	file, err := os.ReadFile("../main.lla")
	if err != nil {
		panic(err)
	}

	fileString := string(file)

	lexer := lexer.Lexer{
		Tokens:       make([]*types.Token, 0),
		DigitsBuffer: make([]rune, 0),
		StringBuffer: make([]rune, 0),
	}

	lexer.Tokenize(fileString)

	for _, v := range lexer.Tokens {
		PrintToken(v)
	}
}

func PrintToken(token *types.Token) {
	fmt.Printf("\nType: %q\nRepresentation: %q\n", types.TranslateTokenTypeToString(token.Type), token.Representation)
}
