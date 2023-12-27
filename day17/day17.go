package day17

import (
	"aoc/aoc_util/astar"
)

const (
	east = 1 << iota
	south
	north
	west
)

type DistNode struct {
	X, Y, Dir, Dist int
}

/*
Process

https://en.wikipedia.org/wiki/A*_search_algorithm

# Heuristic of 0 since limitations are only on direction and cumulative weights, not ensuring shorter dist to goal

Start at (0,0) going east or south - Need to track
Mark Visited
Get neighbors (left, right, forward if dist < 3)

	Cost is currentCost + weight[neighbor]
	Push neighbor, cost, heuristic to queue

This takes ~200s for part 2... Needs some optimization, think it's the super simple priority queue sorting that's causing issues
TODO: Reimplement this so that it takes < 1s
*/
func Process(lines []string, part int) int {
	gridLen := len(lines)
	xLen := gridLen
	yLen := len(lines[0])

	withinGrid := func(x, y int) bool {
		return x >= 0 &&
			y >= 0 &&
			x < xLen &&
			y < yLen
	}

	neighbors := func(node *DistNode) []DistNode {
		var neighbors []DistNode
		var x, y int
		if node.Dist < 3 {
			x, y = nextNode(node.X, node.Y, node.Dir)
			if withinGrid(x, y) {
				neighbors = append(neighbors, DistNode{X: x, Y: y, Dir: node.Dir, Dist: node.Dist + 1})
			}
		}
		left, right := turnLeftRight(node.Dir)
		x, y = nextNode(node.X, node.Y, left)
		if withinGrid(x, y) {
			neighbors = append(neighbors, DistNode{X: x, Y: y, Dir: left, Dist: 1})
		}
		x, y = nextNode(node.X, node.Y, right)
		if withinGrid(x, y) {
			neighbors = append(neighbors, DistNode{X: x, Y: y, Dir: right, Dist: 1})
		}
		return neighbors
	}

	ultraNeighbors := func(node *DistNode) []DistNode {
		var ultraNeighbor []DistNode
		var x, y int
		if node.Dist < 10 {
			x, y = nextNode(node.X, node.Y, node.Dir)
			if withinGrid(x, y) {
				ultraNeighbor = append(ultraNeighbor, DistNode{X: x, Y: y, Dir: node.Dir, Dist: node.Dist + 1})
			}
		}
		if node.Dist >= 4 || node.Dist == 0 {
			left, right := turnLeftRight(node.Dir)
			x, y = nextNode(node.X, node.Y, left)
			if withinGrid(x, y) {
				ultraNeighbor = append(ultraNeighbor, DistNode{X: x, Y: y, Dir: left, Dist: 1})
			}
			x, y = nextNode(node.X, node.Y, right)
			if withinGrid(x, y) {
				ultraNeighbor = append(ultraNeighbor, DistNode{X: x, Y: y, Dir: right, Dist: 1})
			}
		}
		return ultraNeighbor
	}
	partNeighbors := neighbors
	if part == 2 {
		partNeighbors = ultraNeighbors
	}
	endFunc := func(node *DistNode) bool {
		return node.X+1 == xLen && node.Y+1 == yLen
	}
	if part == 2 {
		endFunc = func(node *DistNode) bool {
			return node.X+1 == xLen && node.Y+1 == yLen && node.Dist >= 4
		}
	}

	sum, _, _ := astar.AStar[DistNode](
		DistNode{Dir: east},
		xLen*yLen,
		endFunc,
		partNeighbors,
		func(node *DistNode) int {
			return int(lines[node.X][node.Y] - '0')
		},
		func(node *DistNode) int {
			return 0 // no directional weight applied
		},
	)

	return sum
}

func nextNode(x, y, dir int) (int, int) {
	switch dir {
	case east:
		return x, y + 1
	case west:
		return x, y - 1
	case north:
		return x - 1, y
	case south:
		return x + 1, y
	}
	return -1, -1
}

func turnLeftRight(dir int) (int, int) {
	switch dir {
	case east:
		return north, south
	case west:
		return south, north
	case north:
		return west, east
	case south:
		return east, west
	}
	return -1, -1
}
