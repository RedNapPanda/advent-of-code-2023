package day12

import (
	aoc "aoc"
	"github.com/stretchr/testify/assert"
	"testing"
)

var example = []string{
	"???.### 1,1,3",
	".??..??...?##. 1,1,3",
	"?#?#?#?#?#?#?#? 1,3,1,6",
	"????.#...#... 4,1,1",
	"????.######..#####. 1,6,5",
	"?###???????? 3,2,1",
}

func TestPart1Example(t *testing.T) {
	output := Part1(example)
	assert.Equal(t, 21, output)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputData(12)
	output := Part1(input)

	assert.Equal(t, 6949, output)
}
