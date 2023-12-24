package day14

import (
	aoc "aoc"
	"aoc/aoc_util"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var example = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

func TestPart1Example(t *testing.T) {
	output := Part1(strings.Split(example, "\n"))
	assert.Equal(t, 136, output)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputData(14)
	output := Part1(input)

	assert.Equal(t, 113078, output)
}

func TestPart2Example(t *testing.T) {
	defer aoc_util.Timer("example2")()
	output := Part2(strings.Split(example, "\n"), 1_000_000_000)
	assert.Equal(t, 64, output)
}

func TestPart2(t *testing.T) {
	defer aoc_util.Timer("part2")()
	input, _ := aoc.GetInputData(14)
	output := Part2(input, 1_000_000_000)

	assert.Equal(t, 94255, output)
}

func TestPart2ExampleNorth(t *testing.T) {
	defer aoc_util.Timer("north")()
	lines := strings.Split(example, "\n")
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = make([]byte, len(line))
		for byteIdx, b := range []byte(line) {
			grid[i][byteIdx] = b
		}
	}

	aoc_util.PrintMatrix(grid)
	fmt.Printf("\n")
	shiftNorth(grid)

	aoc_util.PrintMatrix(grid)
}

func TestPart2ExampleSouth(t *testing.T) {
	defer aoc_util.Timer("south")()
	lines := strings.Split(example, "\n")
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = make([]byte, len(line))
		for byteIdx, b := range []byte(line) {
			grid[i][byteIdx] = b
		}
	}

	aoc_util.PrintMatrix(grid)
	fmt.Printf("\n")
	shiftSouth(grid)

	aoc_util.PrintMatrix(grid)
}

func TestPart2ExampleEast(t *testing.T) {
	defer aoc_util.Timer("east")()
	lines := strings.Split(example, "\n")
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = make([]byte, len(line))
		for byteIdx, b := range []byte(line) {
			grid[i][byteIdx] = b
		}
	}

	aoc_util.PrintMatrix(grid)
	fmt.Printf("\n")
	shiftEast(grid)

	aoc_util.PrintMatrix(grid)
}

func TestPart2ExampleWest(t *testing.T) {
	defer aoc_util.Timer("west")()
	lines := strings.Split(example, "\n")
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = make([]byte, len(line))
		for byteIdx, b := range []byte(line) {
			grid[i][byteIdx] = b
		}
	}

	aoc_util.PrintMatrix(grid)
	fmt.Printf("\n")
	shiftWest(grid)

	aoc_util.PrintMatrix(grid)
}
