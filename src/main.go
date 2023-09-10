package main

import (
	"fmt"
	"lla/pkg/lexer"
	"lla/pkg/types"
	"os"
	"strconv"

	"golang.org/x/exp/slices"
)

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func isSpace(char rune) bool {
	return char == ' ' || char == '\n' || char == '\t' || char == '\r' || char == '\v' || char == '\f'
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}()

	isMeaningfulChar := []rune{',', '.', ';'}

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

	for _, char := range fileString {
		token := new(types.Token)
		switch string(lexer.StringBuffer) {
		case "return":
			lexer.PopStringBuffer()
			token.Type = types.RETURN
			lexer.Tokens = append(lexer.Tokens, token)
		case "let":
			lexer.PopStringBuffer()
			token.Type = types.LET
			lexer.Tokens = append(lexer.Tokens, token)
		default:
			if len(lexer.Tokens) != 0 {
				if lexer.Tokens[len(lexer.Tokens)-1].Type == types.LET {
					token.Representation = string(lexer.PopStringBuffer())
					token.Type = types.LET_NAME
					lexer.Tokens = append(lexer.Tokens, token)
				}
			}
		}

		if lexer.IsLastCharNumeric && char == '.' {
			lexer.DigitsBuffer = append(lexer.DigitsBuffer, char)
			lexer.IsLastCharFloat = true
		} else if isNumeric(string(char)) {
			lexer.DigitsBuffer = append(lexer.DigitsBuffer, char)
			lexer.IsLastCharNumeric = true
		}

		if isSpace(char) {
			lexer.ResetNumericFlags()
			lexer.PopDigitBuffer()
			lexer.PopStringBuffer()
		} else if char == ';' {
			if lexer.IsLastCharFloat {
				lexer.ResetNumericFlags()
				token.Type = types.FLOAT_LITERAL
				token.Representation = string(lexer.PopDigitBuffer())
				lexer.Tokens = append(lexer.Tokens, token)
			} else if lexer.IsLastCharNumeric {
				lexer.ResetNumericFlags()
				token.Type = types.INTEGER_LITERAL
				token.Representation = string(lexer.PopDigitBuffer())
				lexer.Tokens = append(lexer.Tokens, token)
			}

			token = new(types.Token)
			lexer.PopStringBuffer()
			token.Type = types.SEMICOLUMN
			lexer.Tokens = append(lexer.Tokens, token)
		} else {
			lexer.StringBuffer = append(lexer.StringBuffer, char)
		}

		if (lexer.IsLastCharNumeric && !isNumeric(string(char))) && !slices.Contains(isMeaningfulChar, char) {
			panic("cannot start an identifier with a number")
		}
	}

	for _, v := range lexer.Tokens {
		PrintToken(v)
	}
}

func PrintToken(token *types.Token) {
	fmt.Printf("\nType: %q\nRepresentation: %q\n", types.TranslateTokenTypeToString(token.Type), token.Representation)
}
