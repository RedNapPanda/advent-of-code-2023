package day9

import (
	aoc "aoc"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Calculating possibilities from the end since we have a set number of broken spring groups to match
?#?#?#?#?#?#?#? | (4, 0) ? | 1 [1 3 1 6]
[? # ? # ? # ? # ? # ? # ? # ? _]
[0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 1]
[0 0 0 0 0 0 1 1 2 1 0 0 0 0 0 0]
[0 0 0 0 2 1 3 1 0 0 0 0 0 0 0 0]
[3 1 4 1 0 0 0 0 0 0 0 0 0 0 0 0]
[5 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
*/
var example = []string{
	"0 3 6 9 12 15",
	"1 3 6 10 15 21",
	"10 13 16 21 30 45",
}

func TestPart1Example(t *testing.T) {
	output, _ := Process(example)

	assert.Equal(t, 114, output)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputData(9)
	output, _ := Process(input)

	assert.Equal(t, 1938731307, output)
}

func TestPart2Example(t *testing.T) {
	_, output := Process(example)

	assert.Equal(t, 2, output)
}

func TestPart2(t *testing.T) {
	input, _ := aoc.GetInputData(9)
	_, output := Process(input)

	assert.Equal(t, 948, output)
}
