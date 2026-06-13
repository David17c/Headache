package main

import (
	"fmt"
	"headache/internal"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		Help()
	}

	ParseArgs(args)
}

func ParseArgs(args []string) {
	var outputPath string
	var inputPath string

	// If the user wants to directly run a headache or brainfuck file
	if args[1] == "run" {
		if len(args) < 3 || len(args) > 3 {
			fmt.Printf("Usage: %s run <file>\n", args[0])
			return
		}
		inputPath = args[2]
		if strings.HasSuffix(strings.ToLower(inputPath), ".bf") {
			bf_output, err := os.ReadFile(inputPath)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			internal.Interpreter(string(bf_output))
			return
		}
		bf_output := internal.Translator(inputPath)
		internal.Interpreter(bf_output)
		// If the user wants to compile the headache file to a brainfuck file
	} else {
		inputPath = args[1]
		if len(args) > 2 {
			outputPath = args[2]
		}
		bf_output := internal.Translator(inputPath)
		if outputPath == "" {
			fmt.Println(bf_output)
			return
		}

		err := os.WriteFile(outputPath, []byte(bf_output), 0666)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func Help() {
	// ToDo will fix later
}
