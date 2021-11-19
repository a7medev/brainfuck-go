package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: brainfuck program.b\n")
		os.Exit(1)
	}

	file := os.Args[1]
	contents, err := os.ReadFile(file)

	if err != nil {
		fmt.Printf("Failed to load file %v\n", file)
		os.Exit(1)
	}

	memory := [30000]byte{}
	ptr := 0

	Execute(string(contents), memory, ptr)
}

func Execute(contents string, memory [30000]uint8, ptr int) {
	reader := bufio.NewReader(os.Stdin)

	for _, c := range contents {
		switch c {
		case '+':
			memory[ptr]++
		case '-':
			memory[ptr]--
		case '<':
			ptr++
		case '>':
			ptr--
		case '.':
			fmt.Printf("%c", memory[ptr])
		case ',':
			result, err := reader.ReadByte()

			if err != nil {
				fmt.Println("Failed to read input from stdin")
				os.Exit(1)
			}

			memory[ptr] = result
		}
	}
}
