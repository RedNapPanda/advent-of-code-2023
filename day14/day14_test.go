package day14

import (
	aoc "aoc"
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

/*
##########
O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
*/
func TestPart1Example(t *testing.T) {
	output := Process(strings.Split(example, "\n"))
	assert.Equal(t, 136, output)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputData(14)
	output := Process(input)

	assert.Equal(t, 113078, output)
}
