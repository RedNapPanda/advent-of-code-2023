package day2

import (
	aoc "aoc"
	"github.com/stretchr/testify/assert"
	"testing"
)

var example = []string{
	"game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	"game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	"game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	"game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	"game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
}

func TestPart1Example(t *testing.T) {
	maxRed, maxGreen, maxBlue := 12, 13, 14
	sum := Part1(example, maxRed, maxGreen, maxBlue)
	assert.Equal(t, 8, sum)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputData(2)
	maxRed, maxGreen, maxBlue := 12, 13, 14
	sum := Part1(input, maxRed, maxGreen, maxBlue)
	assert.Equal(t, 3035, sum)
}

func TestPart2Example(t *testing.T) {
	sum := Part2(example)
	assert.Equal(t, 2286, sum)
}

func TestPart2(t *testing.T) {
	input, _ := aoc.GetInputData(2)
	sum := Part2(input)
	assert.Equal(t, 66027, sum)
}
