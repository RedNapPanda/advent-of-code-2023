package day11

import (
	aoc "aoc"
	"github.com/stretchr/testify/assert"
	"testing"
)

var example = []string{
	"...#......",
	".......#..",
	"#.........",
	"..........",
	"......#...",
	".#........",
	".........#",
	"..........",
	".......#..",
	"#...#.....",
}

func TestPart1Example(t *testing.T) {
	output := Part1(example)

	assert.Equal(t, 374, output)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputData(11)
	output := Part1(input)

	assert.Equal(t, 9274989, output)
}

func TestPart2Example1(t *testing.T) {
	output := Part2(example, 10)

	assert.Equal(t, 1030, output)
}

func TestPart2Example2(t *testing.T) {
	output := Part2(example, 100)

	assert.Equal(t, 8410, output)
}

func TestPart2(t *testing.T) {
	input, _ := aoc.GetInputData(11)
	output := Part2(input, 1_000_000)

	assert.Equal(t, 357134560737, output)
}
