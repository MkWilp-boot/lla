package lexer

import (
	"io"
	"lla/pkg/helpers"
	"lla/pkg/types"

	"golang.org/x/exp/slices"
)

var isMeaningfulChar = []rune{',', '.', ';'}

type Lexer struct {
	Tokens            []*types.Token
	IsLastCharNumeric bool
	IsLastCharFloat   bool
	DigitsBuffer      []rune
	StringBuffer      []rune
}

func (lexer *Lexer) Tokenize(content string) {
	for _, char := range content {
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
		} else if helpers.IsNumeric(string(char)) {
			lexer.DigitsBuffer = append(lexer.DigitsBuffer, char)
			lexer.IsLastCharNumeric = true
		}

		if helpers.IsSpace(char) {
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

		if (lexer.IsLastCharNumeric && !helpers.IsNumeric(string(char))) && !slices.Contains(isMeaningfulChar, char) {
			panic("cannot start an identifier with a number")
		}
	}
}

func (lexer *Lexer) ResetNumericFlags() {
	lexer.IsLastCharFloat = false
	lexer.IsLastCharNumeric = false
}

func (lexel *Lexer) PopStringBuffer() []rune {
	content := lexel.StringBuffer
	if len(lexel.StringBuffer) > 0 {
		lexel.StringBuffer = slices.Delete(lexel.StringBuffer, 0, len(lexel.StringBuffer))
	}
	return content
}

func (lexel *Lexer) PopDigitBuffer() []rune {
	content := lexel.DigitsBuffer
	if len(lexel.DigitsBuffer) > 0 {
		lexel.DigitsBuffer = slices.Delete(lexel.DigitsBuffer, 0, len(lexel.DigitsBuffer))
	}
	return content
}

func (lexer *Lexer) LastDigit() (rune, error) {
	if len(lexer.DigitsBuffer) == 0 {
		return 0, io.EOF
	}

	return lexer.DigitsBuffer[len(lexer.DigitsBuffer)-1], nil
}

func (lexer *Lexer) LastString() (rune, error) {
	if len(lexer.StringBuffer) == 0 {
		return 0, io.EOF
	}

	return lexer.StringBuffer[len(lexer.StringBuffer)-1], nil
}
