package day6

import (
	aoc "aoc"
	"github.com/stretchr/testify/assert"
	"testing"
)

var example = []string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
}

func TestPart1Example(t *testing.T) {
	output := Part1(example)

	assert.Equal(t, 6440, output)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputData(7)
	output := Part1(input)

	assert.Equal(t, 248422077, output)
}

func TestPart2Example(t *testing.T) {
	output := Part2(example)

	assert.Equal(t, 5905, output)
}

func TestPart2(t *testing.T) {
	input, _ := aoc.GetInputData(7)
	output := Part2(input)

	assert.Equal(t, 249817836, output)
}
