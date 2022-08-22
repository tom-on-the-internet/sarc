package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

//nolint
var formats = map[string]func(string) string{
	"ball":       Ball,
	"cursive":    Cursive,
	"lowercase":  Lowercase,
	"nomodify":   NoModify,
	"reverse":    Reverse,
	"rot13":      Rot13,
	"sarcastic":  Sarcastic,
	"smallcaps":  Smallcaps,
	"stencil":    Stencil,
	"swapcase":   Swapcase,
	"uppercase":  Uppercase,
	"upsidedown": UpsideDown,
}

type options struct {
	interactive bool
	format      string
}

func main() {
	options := getOptions()
	input := getInput()

	if options.interactive {
		handleInteractive(options.format, input)
		return
	}

	output := getOutput(options.format, input)

	fmt.Println(output)
}

func handleInteractive(format, input string) {
	p := tea.NewProgram(initialModel(format, input))

	teaModel, err := p.StartReturningModel()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	finalModel, ok := teaModel.(model)
	if !ok {
		panic("model isn't a model. this should never happen.")
	}

	output := "\n\n" + getOutput(finalModel.formats[finalModel.cursor], input)
	fmt.Println(output)
}

func getOptions() options {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "USAGE %s [options] input\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "EXAMPLE %s -f sarcastic 'You are so smart.'\n\n", os.Args[0])

		fmt.Fprint(os.Stderr, "OPTIONS\n\n")

		fmt.Fprint(os.Stderr, "-f, --format         format style\n")
		fmt.Fprint(os.Stderr, "-i, --interactive    interactive mode\n")
		fmt.Fprint(os.Stderr, "\n")

		fmt.Fprint(os.Stderr, "FORMATS\n\n")

		for format := range formats {
			fmt.Fprintf(os.Stderr, "%s\n", format)
		}
	}

	format := ""

	flag.StringVar(&format, "format", "", "")
	flag.StringVar(&format, "f", "", "")

	interactive := flag.Bool("i", false, "Interactive")

	flag.Parse()

	format = strings.TrimSpace(format)

	// default format
	if format == "" {
		format = "sarcastic"
	}

	if _, ok := formats[format]; !ok {
		fmt.Fprintf(os.Stderr, "invalid format: %s\n\n", format)
		flag.Usage()
		os.Exit(1)
	}

	return options{format: format, interactive: *interactive}
}

func getInput() string {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	return string(bytes)
}

func getOutput(format, input string) string {
	fn := formats[format]
	return fn(input)
}

func getFormats() []string {
	keys := make([]string, 0, len(formats))
	for k := range formats {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	return keys
}
