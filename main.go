package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	args := os.Args[1:]
	text := ArgsToString(args)

	if len(text) == 0 {
		fmt.Println("USAGE: sarc The text you want to be sarcastic")
		os.Exit(1)
	}

	sarcasticText := ToSarcastic(text)
	fmt.Println(sarcasticText)
}

func ArgsToString(args []string) string {
	return strings.Trim(strings.Join(args, " "), " ")
}

func ToSarcastic(text string) string {
	var output []rune
	shouldUpperCase := true

	for _, character := range text {
		if unicode.IsLetter(character) {
			shouldUpperCase = !shouldUpperCase
		}

		if shouldUpperCase {
			output = append(output, unicode.ToUpper(character))
		} else {
			output = append(output, unicode.ToLower(character))
		}
	}

	return string(output)
}
