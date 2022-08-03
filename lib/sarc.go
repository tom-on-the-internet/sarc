package lib

import (
	"strings"
	"unicode"
)

func Sarcastic(str string) string {
	str = strings.ToLower(str)
	runes := []rune(str)
	upperCase := false

	for idx := range runes {
		if !unicode.IsLetter(runes[idx]) {
			continue
		}

		if upperCase {
			runes[idx] = unicode.ToUpper(runes[idx])
		}

		upperCase = !upperCase
	}

	return string(runes)
}

func NoModify(str string) string {
	return str
}

func Uppercase(str string) string {
	return strings.ToUpper(str)
}

func Lowercase(str string) string {
	return strings.ToLower(str)
}

func Reverse(str string) string {
	// TODO:
	return strings.ToLower(str)
}

func Swapcase(str string) string {
	// TODO:
	return strings.ToLower(str)
}
