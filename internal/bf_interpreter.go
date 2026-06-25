package internal

import (
	"bufio"
	"fmt"
	"os"
)

const datasize = 65535

func Interpreter(program string) {
	data := make([]byte, datasize)

	dataPtr := 0
	instrPtr := 0

	brackets := make(map[int]int)
	stack := []int{}

	for i, c := range program {
		switch c {
		case '[':
			stack = append(stack, i)
		case ']':
			if len(stack) == 0 {
				panic("unmatched ]")
			}

			open := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			brackets[open] = i
			brackets[i] = open
		}
	}

	if len(stack) != 0 {
		panic("unmatched [")
	}

	reader := bufio.NewReader(os.Stdin)

	for instrPtr < len(program) {
		switch program[instrPtr] {

		case '>':
			dataPtr++
			if dataPtr >= datasize {
				panic("data pointer out of bounds")
			}

		case '<':
			dataPtr--
			if dataPtr < 0 {
				panic("data pointer out of bounds")
			}

		case '+':
			data[dataPtr]++

		case '-':
			data[dataPtr]--

		case '.':
			fmt.Printf("%c", data[dataPtr])

		case ',':
			b, err := reader.ReadByte()
			if err != nil {
				data[dataPtr] = 0
			} else {
				data[dataPtr] = b
			}

		case '[':
			if data[dataPtr] == 0 {
				instrPtr = brackets[instrPtr]
			}

		case ']':
			if data[dataPtr] != 0 {
				instrPtr = brackets[instrPtr]
			}
		}

		instrPtr++
	}
}
