package lexer

import (
	"io"
	"lla/pkg/types"

	"golang.org/x/exp/slices"
)

type Lexer struct {
	Tokens            []*types.Token
	IsLastCharNumeric bool
	IsLastCharFloat   bool
	DigitsBuffer      []rune
	StringBuffer      []rune
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
