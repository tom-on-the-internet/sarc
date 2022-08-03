package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		os.Stderr.WriteString("Must provide an argument\n")
		os.Exit(1)
	}

	str := os.Args[1]
	sarc := sarc(str)
	fmt.Println(sarc)
}

func sarc(str string) string {
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
