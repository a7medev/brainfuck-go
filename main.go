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

	p := NewParser(string(contents))
	nodes := p.Parse(0, 0)

	if p.Error != nil {
		fmt.Println(p.Error)
	}

	memory := [30000]byte{}
	ptr := 0
	Execute(nodes, &memory, &ptr)
}

func Execute(nodes []Node, memory *[30000]uint8, ptr *int) {
	reader := bufio.NewReader(os.Stdin)
	for _, node := range nodes {
		switch node.Type {
		case AddNode:
			memory[*ptr]++
		case SubNode:
			memory[*ptr]--
		case MoveRightNode:
			*ptr++
		case MoveLeftNode:
			*ptr--
		case OutputNode:
			fmt.Printf("%c", memory[*ptr])
		case InputNode:
			result, err := reader.ReadByte()

			if err != nil {
				fmt.Println("Failed to read input from stdin")
				os.Exit(1)
			}

			memory[*ptr] = result
		case LoopNode:
			for memory[*ptr] != 0 {
				Execute(node.Children, memory, ptr)
			}
		}
	}
}
