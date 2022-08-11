package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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
	// if args, use them and ignore stdin
	if args := flag.Args(); len(args) != 0 {
		return strings.Join(args, " ")
	}

	// no args, let's use stdin
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	return string(bytes)
}

func getOutput(mode, input string) string {
	fn := modes[mode]
	return fn(input)
}
