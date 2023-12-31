package day14

import (
	aoc "aoc"
	"aoc/aoc_util"
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
