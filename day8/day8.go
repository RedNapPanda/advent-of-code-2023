package day8

type Node struct {
	left  string
	right string
}

func Part1(lines []string) int {
	parseNodes(lines)
	return -1
}

func Part2(lines []string) int {
	return -1
}

func parseNodes(lines []string) (string, *map[string]Node) {
	nodes := make(map[string]Node)
	instructions := lines[0]

	return instructions, &nodes
}
