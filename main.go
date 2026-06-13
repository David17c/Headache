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
		return
	}

	if strings.ToLower(args[1]) == "help" {
		Help()
		return
	}

	ParseArgs(args)
}

func ParseArgs(args []string) {
	var outputPath string
	var inputPath string

	// If the user wants to directly run a headache or brainfuck file
	if args[1] == "run" {
		if len(args) < 3 || len(args) > 3 {
			Help()
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
	fmt.Printf(
		"Usage:\n" +
			"headache run <brainfuck or headache file> (runs code directly)\n" +
			"headache <example.ha> (compile and print bf to terminal)\n" +
			"headache <example.ha> <output.bf> (write compiled brainfuck to output.bf)\n",
	)
}
