package day16

import (
	"aoc/aoc_util"
)

/*
Can bitwise pack this into the top half of the int
i |= dir << 4 - store direction bit
i >> 4 - pull out direction bit
i &= 0xF - bottom 4 bytes
1 << 4 == 10000 == 16 == 2^(4-1)
*/
const (
	// Storing these in the upper 4 bits of the byte i |= west << 4, so we could store all 4 states
	north = 1 << iota
	east
	south
	west
)

var charBits = map[byte]int{
	'.':  1,
	'-':  2,
	'|':  3,
	'\\': 4,
	'/':  5,
	'<':  6,
	'>':  7,
	'^':  8,
	'V':  9,
}

var cardinalMirror = map[int]int{
	north: south,
	south: north,
	east:  west,
	west:  east,
}

var rtlCardinals = map[int]int{
	north: east,
	east:  north,
	south: west,
	west:  south,
}

var ltrCardinals = map[int]int{
	north: west,
	west:  north,
	south: east,
	east:  south,
}

func Process(lines []string, part int) int {
	gridLen := len(lines)
	grid := make([][]int, gridLen)
	originalGrid := make([][]int, gridLen)
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		originalGrid[i] = make([]int, len(line))
		for byteIdx, b := range []byte(line) {
			grid[i][byteIdx] = charBits[b]
			originalGrid[i][byteIdx] = charBits[b]
		}
	}

	count, maxCount, x := 0, 0, 0
	if part == 2 {
		x = len(grid) - 2
	}
	for i := 0; i < x+1; i++ {
		for j := 0; j < x+1; j++ {
			for n := 0; n < len(originalGrid[0]); n++ {
				copy(grid[n], originalGrid[n])
			}
			var nodes []aoc_util.CardinalNode

			iterate := func() {
				count = 0
				for len(nodes) > 0 {
					nNodes, b := iterateNodes(nodes[0], grid, gridLen)
					nodes = append(nodes[1:], nNodes...)
					if b {
						count++
					}
				}
				maxCount = max(maxCount, count)
			}

			// This is stupid lmao
			if part == 2 {
				if i == 0 {
					if j == 0 {
						nodes = []aoc_util.CardinalNode{{X: i, Y: j, Cardinal: east}}
						iterate()
					}
					nodes = []aoc_util.CardinalNode{{X: i, Y: j, Cardinal: south}}
				} else if i == x {
					if j == x {
						nodes = []aoc_util.CardinalNode{{X: i, Y: j, Cardinal: west}}
					}
					nodes = []aoc_util.CardinalNode{{X: i, Y: j, Cardinal: north}}
				} else if j == 0 {
					if i == 0 {
						nodes = []aoc_util.CardinalNode{{X: i, Y: j, Cardinal: south}}
						iterate()
					}
					nodes = []aoc_util.CardinalNode{{X: i, Y: j, Cardinal: east}}
				} else if j == x {
					if i == x {
						nodes = []aoc_util.CardinalNode{{X: i, Y: j, Cardinal: north}}
					}
					nodes = []aoc_util.CardinalNode{{X: i, Y: j, Cardinal: west}}
				}
			} else {
				nodes = []aoc_util.CardinalNode{{X: 0, Y: 0, Cardinal: east}}
			}
			iterate()
		}
	}

	if part == 1 {
		return count
	} else {
		return maxCount
	}
}

/*
Part 2 naive approach would just be iterate row * col for O(n^2)
*/

func iterateNodes(node aoc_util.CardinalNode, grid [][]int, gridLen int) ([]aoc_util.CardinalNode, bool) {
	if node.X < 0 || node.X >= gridLen ||
		node.Y < 0 || node.Y >= gridLen ||
		hasVisited(grid, node.X, node.Y, node.Cardinal) {
		return nil, false
	}
	previouslyVisited := false
	var nodes []aoc_util.CardinalNode
	if grid[node.X][node.Y]>>4 != 0 {
		previouslyVisited = true
	}

	lower := grid[node.X][node.Y] & 0xF
	grid[node.X][node.Y] |= node.Cardinal << 4

	switch lower {
	case charBits['.']:
		nodes = append(nodes, possibleNodes(grid, node.X, node.Y, node.Cardinal)...)
	case charBits['-']:
		nodes = append(nodes, possibleNodes(grid, node.X, node.Y, east, west)...)
	case charBits['|']:
		nodes = append(nodes, possibleNodes(grid, node.X, node.Y, south, north)...)
	case charBits['\\']: // ccw 90rot
		nodes = append(nodes, possibleNodes(grid, node.X, node.Y, ltrCardinals[node.Cardinal])...)
	case charBits['/']: // clockwise 90rot
		nodes = append(nodes, possibleNodes(grid, node.X, node.Y, rtlCardinals[node.Cardinal])...)
	}
	return nodes, len(nodes) > 0 && !previouslyVisited
}

func possibleNodes(grid [][]int, x, y int, cardinals ...int) []aoc_util.CardinalNode {
	var nodes []aoc_util.CardinalNode
	for _, cardinal := range cardinals {
		nextX, nextY := nextPos(x, y, cardinal)
		if !hasVisited(grid, nextX, nextY, cardinal) {
			nodes = append(nodes, aoc_util.CardinalNode{X: nextX, Y: nextY, Cardinal: cardinal})
		}
	}
	return nodes
}

func hasVisited(grid [][]int, x, y, cardinal int) bool {
	return x >= 0 &&
		y >= 0 &&
		y < len(grid) &&
		x < len(grid) &&
		(grid[x][y]>>4)&cardinal == cardinal
}

func nextPos(x, y, cardinal int) (int, int) {
	// progress to next Node based on cardinal and x/y
	switch cardinal {
	case north:
		return x - 1, y
	case south:
		return x + 1, y
	case east:
		return x, y + 1
	case west:
		return x, y - 1
	}
	return -1, -1
}
