package main

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
	strs := strings.Split(str, "")

	for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
		strs[i], strs[j] = strs[j], strs[i]
	}

	str = strings.Join(strs, "")

	return str
}

func Swapcase(str string) string {
	runes := []rune(str)
	for idx, rune := range runes {
		if !unicode.IsLetter(rune) {
			continue
		}

		if unicode.IsUpper(rune) {
			runes[idx] = unicode.ToLower(rune)
			continue
		}

		runes[idx] = unicode.ToUpper(rune)
	}

	return string(runes)
}

// TODO
func Smallcaps(str string) string {
	return str
}
