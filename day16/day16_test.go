package day16

import (
	aoc "aoc"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var example = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

func TestPart1Example(t *testing.T) {
	output := Process(strings.Split(example, "\n"), 1)
	assert.Equal(t, 46, output)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputData(16)
	output := Process(input, 1)

	assert.Equal(t, 6883, output)
}

func TestPart2Example(t *testing.T) {
	output := Process(strings.Split(example, "\n"), 2)
	assert.Equal(t, 51, output)
}

func TestPart2(t *testing.T) {
	input, _ := aoc.GetInputData(16)
	output := Process(input, 2)

	assert.Equal(t, 7228, output)
}
