package day3

import (
	aoc "aoc"
	"github.com/stretchr/testify/assert"
	"testing"
)

var example = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
}

func TestPart1Example(t *testing.T) {
	output := Part1(example)

	assert.Equal(t, 4361, output)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputData(3)
	output := Part1(input)

	assert.Equal(t, 525181, output)
}

func TestPart2Example(t *testing.T) {
	output := Part2(example)

	assert.Equal(t, 467835, output)
}

func TestPart2(t *testing.T) {
	input, _ := aoc.GetInputData(3)
	output := Part2(input)

	assert.Equal(t, 84289137, output)
}
