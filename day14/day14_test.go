package day14

import (
	aoc "aoc"
	"github.com/stretchr/testify/assert"
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
	output := Process(example)
	assert.Equal(t, 405, output)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputDataString(14)
	output := Process(input)

	assert.Equal(t, 36015, output)
}
