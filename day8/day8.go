package day8

import (
	"regexp"
)

var nodeRegex = regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)

type treeNode struct {
	key         string
	left, right *treeNode
}

type node struct {
	key, left, right string
}

type instruction struct {
	instr string
	index int
}

func (i *instruction) next() byte {
	value := i.instr[i.index]
	i.index++
	if i.index >= len(i.instr) {
		i.index = 0
	}
	return value
}

func Part1(lines []string) int {
	instr, nodeMap, _ := parseNodes(lines, func(node) bool { return false })
	node := nodeMap["AAA"]
	count := 0
	for node.key != "ZZZ" {
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
	}

	return count
}

func Part2(lines []string) int {
	count := 0
	instr, nodeMap, nodes := parseNodes(lines, func(n node) bool { return n.key[2] == 'A' })
	result := -1

	for _, node := range nodes {
		instr.index = 0
		count = 0
		for node.key[2] != 'Z' {
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
		}
		if result == -1 {
			result = count
		} else {
			result = LCM(result, count)
		}
	}

	return result
}

func parseNodes(lines []string, matcher func(node) bool) (instruction, map[string]node, []node) {
	nodeMap := make(map[string]node)
	var nodes []node
	instr := instruction{
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
		node := node{
			groups[0],
			groups[1],
			groups[2],
		}
		nodeMap[groups[0]] = node
		if matcher(node) {
			nodes = append(nodes, node)
		}
	}

	return instr, nodeMap, nodes
}

func nodeGroups(line string) []string {
	matches := nodeRegex.FindAllStringSubmatch(line, -1)
	if len(matches) == 0 {
		return nil
	}
	return matches[0][1:]
}

// GCD credits: https://en.wikipedia.org/wiki/Greatest_common_divisor#Euclidean_algorithm
// (a, b) -> (b, a mod b) repeat until (d, 0) where d = GCD
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

// LCM credits: https://en.wikipedia.org/wiki/Least_common_multiple#Using_the_greatest_common_divisor
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}
