package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Translator(input_path string) string {
	var output string
	var lineNumber int

	var open_loops_counter int

	file, err := os.Open(input_path)
	if err != nil {
		fmt.Printf("Error: unable to open %s\n", input_path)
		return ""
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		lineNumber++

		// Skip comments
		if idx := strings.Index(line, "//"); idx != -1 {
			line = strings.TrimSpace(line[:idx])
		}

		// Skip empty lines
		if line == "" {
			continue
		}

		line = strings.ToLower(line)
		fields := strings.Fields(line)

		switch fields[0] {
		case "forward":
			s, err := repeatCommand(fields, ">")
			if err != nil {
				fmt.Println(err)
				return ""
			}
			output += s

		case "back":
			s, err := repeatCommand(fields, "<")
			if err != nil {
				fmt.Println(err)
				return ""
			}
			output += s

		case "increment":
			s, err := repeatCommand(fields, "+")
			if err != nil {
				fmt.Println(err)
				return ""
			}
			output += s

		case "decrement":
			s, err := repeatCommand(fields, "-")
			if err != nil {
				fmt.Println(err)
				return ""
			}
			output += s

		case "input":
			s, err := repeatCommand(fields, ",")
			if err != nil {
				fmt.Println(err)
				return ""
			}
			output += s

		case "print":
			s, err := repeatCommand(fields, ".")
			if err != nil {
				fmt.Println(err)
				return ""
			}

			output += s

		case "loop":
			output += "["
			open_loops_counter++
		case "end":
			if open_loops_counter <= 0 {
				fmt.Printf("Error in line %d: ending a loop that doesn't exist\n", lineNumber)
				return ""
			}
			output += "]"
			open_loops_counter--
		}
	}
	if open_loops_counter > 0 {
		fmt.Printf("Error: you forgot to close a loop\n")
		return ""
	}
	return output
}

func repeatCommand(fields []string, symbol string) (string, error) {
	if len(fields) < 2 {
		return symbol, nil
	}

	n, err := strconv.Atoi(fields[1])
	if err != nil {
		return "", fmt.Errorf("invalid number: %s", fields[1])
	}

	return strings.Repeat(symbol, n), nil
}
