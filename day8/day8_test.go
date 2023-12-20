package day8

import (
	aoc "aoc"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var example1 = strings.Split(`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`, "\n")

var example2 = strings.Split(`LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`, "\n")

func TestPart1Example1(t *testing.T) {
	output := Part1(example1)

	assert.Equal(t, 2, output)
}

func TestPart1Example2(t *testing.T) {
	output := Part1(example2)

	assert.Equal(t, 6, output)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputData(7)
	output := Part1(input)

	assert.Equal(t, 0, output)
}

func TestPart2Example(t *testing.T) {
	output := Part2(example1)

	assert.Equal(t, 0, output)
}

func TestPart2(t *testing.T) {
	input, _ := aoc.GetInputData(7)
	output := Part2(input)

	assert.Equal(t, 0, output)
}
