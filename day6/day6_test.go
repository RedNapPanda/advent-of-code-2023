package day6

import (
	aoc "aoc"
	"github.com/stretchr/testify/assert"
	"testing"
)

var example = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

func TestPart1Example(t *testing.T) {
	output := Part1(example)

	assert.Equal(t, 288, output)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputData(6)
	output := Part1(input)

	assert.Equal(t, 449550, output)
}

func TestPart2Example(t *testing.T) {
	output := Part2(example)

	assert.Equal(t, 71503, output)
}

func TestPart2(t *testing.T) {
	input, _ := aoc.GetInputData(6)
	output := Part2(input)

	assert.Equal(t, 28360140, output)
}
