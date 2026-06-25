package main

import (
	"fmt"
	"headache/internal"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		help()
		return
	}

	if strings.EqualFold(args[1], "help") {
		help()
		return
	}

	parseArgs(args)
}

func parseArgs(args []string) {
	if args[1] == "run" {
		runFile(args)
		return
	}

	compileFile(args)
}

func runFile(args []string) {
	if len(args) != 3 {
		help()
		return
	}

	inputPath := args[2]
	ext := strings.ToLower(filepath.Ext(inputPath))

	switch ext {
	case ".bf":
		bfSource, err := os.ReadFile(inputPath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		internal.Interpreter(string(bfSource))

	case ".ha":
		bfOutput := internal.Translator(inputPath)
		internal.Interpreter(bfOutput)

	default:
		help()
	}
}

func compileFile(args []string) {
	inputPath := args[1]

	if !strings.EqualFold(filepath.Ext(inputPath), ".ha") {
		help()
		return
	}

	var outputPath string
	if len(args) > 2 {
		outputPath = args[2]
	}

	bfOutput := internal.Translator(inputPath)

	if outputPath == "" {
		fmt.Println(bfOutput)
		return
	}

	err := os.WriteFile(outputPath, []byte(bfOutput), 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func help() {
	fmt.Print(
		"Usage:\n" +
			"headache run <example.ha> (Run Headache)\n" +
			"headache run <example.bf> (Run Brainfuck)\n" +
			"headache <example.ha> (Compile to Brainfuck and print to terminal)\n" +
			"headache <example.ha> <output.bf> (Compile to a Brainfuck file)\n",
	)
}
