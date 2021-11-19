package main

import "fmt"

type NodeType uint8

const (
	LoopNode = iota
	AddNode
	SubNode
	MoveLeftNode
	MoveRightNode
	InputNode
	OutputNode
)

type Node struct {
	Children []Node
	Type     NodeType
}

type Parser struct {
	Content  string
	Line     int
	Col      int
	Index    int
	Nodes    []Node
	Brackets int
	Error    error
}

func NewParser(content string) *Parser {
	return &Parser{
		Content: content,
		Line:    1,
		Col:     1,
		Index:   0,
	}
}

func (p *Parser) Advance() {
	p.Index++
	p.Col++
}

func (p *Parser) Parse(start int, brackets int) []Node {
	nodes := []Node{}

	for p.Index < len(p.Content) {
		c := p.Content[p.Index]

		switch c {
		case '+':
			nodes = append(nodes, Node{Type: AddNode, Children: nil})
			p.Advance()
		case '-':
			nodes = append(nodes, Node{Type: SubNode, Children: nil})
			p.Advance()
		case '>':
			nodes = append(nodes, Node{Type: MoveRightNode, Children: nil})
			p.Advance()
		case '<':
			nodes = append(nodes, Node{Type: MoveLeftNode, Children: nil})
			p.Advance()
		case '.':
			nodes = append(nodes, Node{Type: OutputNode, Children: nil})
			p.Advance()
		case ',':
			nodes = append(nodes, Node{Type: InputNode, Children: nil})
			p.Advance()
		case '[':
			p.Brackets++
			p.Advance()
			nodes = append(nodes, p.MakeLoop())
		case ']':
			if p.Brackets == 0 {
				p.Error = fmt.Errorf("unexpected ']' at line %d, col %d", p.Line, p.Col)
				return nil
			}
			p.Brackets--
			p.Advance()
			if p.Brackets == brackets-1 {
				return nodes
			}
		case '\n':
			p.Line++
			p.Col = 1
			p.Advance()
		}
	}

	if p.Brackets != 0 {
		p.Error = fmt.Errorf("expected ']' at line %d, col %d", p.Line, p.Col-1)
		return nil
	}

	return nodes
}

func (p *Parser) MakeLoop() Node {
	node := Node{Type: LoopNode}
	node.Children = p.Parse(p.Index, p.Brackets)
	return node
}
