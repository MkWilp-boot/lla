package helpers

import "strconv"

func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func IsSpace(char rune) bool {
	return char == ' ' || char == '\n' || char == '\t' || char == '\r' || char == '\v' || char == '\f'
}
