package types

import "fmt"

type TokenType uint32

var tokenTypeNames map[TokenType]string

func TranslateTokenTypeToString(key TokenType) string {
	value, ok := tokenTypeNames[key]
	if !ok {
		panic(fmt.Sprintf("[DEV] Undefined entry %d on tokenTypeNames", key))
	}
	return value
}

func init() {
	tokenTypeNames = map[TokenType]string{
		INVALID:         "INVALID",
		LET_NAME:        "LET_NAME",
		RETURN:          "RETURN",
		LET:             "LET",
		ASSIGNER:        "ASSIGNER",
		INTEGER_LITERAL: "INTEGER_LITERAL",
		FLOAT_LITERAL:   "FLOAT_LITERAL",
		SEMICOLUMN:      "SEMICOLUMN",
	}

	if len(tokenTypeNames) != int(_TOTAL) {
		panic("[DEV] Missing entry on tokenTypeNames")
	}
}

const (
	INVALID TokenType = iota
	LET_NAME
	RETURN
	LET
	ASSIGNER
	INTEGER_LITERAL
	FLOAT_LITERAL
	SEMICOLUMN
	_TOTAL
)

type Token struct {
	Type           TokenType
	Representation string
}
