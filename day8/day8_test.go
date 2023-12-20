package day8

import (
	aoc "aoc"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPart1Example1(t *testing.T) {
	var example = strings.Split(`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`, "\n")
	output := Part1(example)

	assert.Equal(t, 2, output)
}

func TestPart1Example2(t *testing.T) {
	var example = strings.Split(`LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`, "\n")
	output := Part1(example)

	assert.Equal(t, 6, output)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputData(8)
	output := Part1(input)

	assert.Equal(t, 19783, output)
}

func TestPart2Example(t *testing.T) {
	var example = strings.Split(`LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`, "\n")
	output := Part2(example)

	assert.Equal(t, 6, output)
}

func TestPart2(t *testing.T) {
	input, _ := aoc.GetInputData(8)
	output := Part2(input)

	assert.Equal(t, 9177460370549, output)
}
