package day8

import (
	"regexp"
)

var nodeRegex = regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)

type Node struct {
	key, left, right string
}

type Instruction struct {
	instr string
	index int
}

func (i *Instruction) next() byte {
	value := i.instr[i.index]
	i.index++
	if i.index >= len(i.instr) {
		i.index = 0
	}
	return value
}

func Part1(lines []string) int {
	count := 0
	instr, nodeMap := parseNodes(lines)
	node := nodeMap["AAA"]
	for {
		side := instr.next()
		var next string
		switch side {
		case 'L':
			next = node.left
		case 'R':
			next = node.right
		}
		node = nodeMap[next]
		count++

		if node.key == "ZZZ" {
			break
		}
	}

	return count
}

func Part2(lines []string) int {
	return -1
}

func parseNodes(lines []string) (Instruction, map[string]Node) {
	nodeMap := make(map[string]Node)
	instr := Instruction{
		lines[0],
		0,
	}

	for i := 1; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}
		groups := nodeGroups(lines[i])
		if groups == nil {
			continue
		}
		node := Node{
			groups[0],
			groups[1],
			groups[2],
		}
		nodeMap[groups[0]] = node
	}

	return instr, nodeMap
}

func nodeGroups(line string) []string {
	matches := nodeRegex.FindAllStringSubmatch(line, -1)
	if len(matches) == 0 {
		return nil
	}
	return matches[0][1:]
}
