package day12

import (
	aoc "aoc"
	"aoc/aoc_util"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var example = []string{
	"???.### 1,1,3",             // 1
	".??..??...?##. 1,1,3",      // 4
	"?#?#?#?#?#?#?#? 1,3,1,6",   // 1
	"????.#...#... 4,1,1",       // 1
	"????.######..#####. 1,6,5", // 4
	"?###???????? 3,2,1",        // 10
}

func TestPart1Example(t *testing.T) {
	output := Process(example)
	assert.Equal(t, 21, output)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputData(12)
	output := Process(input)

	assert.Equal(t, 6949, output)
}

func TestPart2(t *testing.T) {
	// day12 took 12.1716ms
	defer aoc_util.Timer("day12")()
	input, _ := aoc.GetInputData(12)
	for i, s := range input {
		split := strings.Split(s, " ")
		input[i] = fmt.Sprintf("%s %s",
			aoc_util.Repeat(split[0], "?", 5),
			aoc_util.Repeat(split[1], ",", 5),
		)
	}
	output := Process(input)

	assert.Equal(t, 51456609952403, output)
}

func TestPart3Example(t *testing.T) {
	output := RecursiveProcess(example)
	assert.Equal(t, 21, output)
}

func TestPart3(t *testing.T) {
	input, _ := aoc.GetInputData(12)
	output := RecursiveProcess(input)

	assert.Equal(t, 6949, output)
}

func TestPart4(t *testing.T) {
	// day12 took 88.9144ms
	defer aoc_util.Timer("day12")()
	input, _ := aoc.GetInputData(12)
	for i, s := range input {
		split := strings.Split(s, " ")
		input[i] = fmt.Sprintf("%s %s",
			aoc_util.Repeat(split[0], "?", 5),
			aoc_util.Repeat(split[1], ",", 5),
		)
	}
	output := RecursiveProcess(input)

	assert.Equal(t, 51456609952403, output)
}
