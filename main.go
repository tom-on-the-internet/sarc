package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// nolint
var modes = map[string]func(string) string{
	"lowercase": Lowercase,
	"nomodify":  NoModify,
	"reverse":   Reverse,
	"sarcastic": Sarcastic,
	"swapcase":  Swapcase,
	"uppercase": Uppercase,
}

func main() {
	mode := getMode()
	input := getInput()
	output := getOutput(mode, input)

	fmt.Println(output)
}

func getMode() string {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "USAGE %s [-m mode] input\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "EXAMPLE %s -m sarcastic 'You are so smart.'\n\n", os.Args[0])

		fmt.Fprint(os.Stderr, "MODES\n\n")

		for mode := range modes {
			fmt.Fprintf(os.Stderr, "%s\n", mode)
		}
	}

	mode := ""

	flag.StringVar(&mode, "mode", "", "")
	flag.StringVar(&mode, "m", "", "")

	flag.Parse()

	mode = strings.TrimSpace(mode)

	// default mode
	if mode == "" {
		mode = "sarcastic"
	}

	if _, ok := modes[mode]; !ok {
		fmt.Fprintf(os.Stderr, "invalid mode: %s\n\n", mode)
		flag.Usage()
		os.Exit(1)
	}

	return mode
}

func getInput() string {
	if args := flag.Args(); len(args) != 0 {
		return args[0]
	}

	// use stdin
	scanner := bufio.NewScanner(os.Stdin)
	strs := []string{}

	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}

	return strings.Join(strs, "\n")
}

func getOutput(mode, input string) string {
	fn := modes[mode]
	return fn(input)
}
