package day18

import (
	aoc "aoc"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var example = `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`

func TestPart1Example(t *testing.T) {
	output := Process(strings.Split(example, "\n"), 1)
	assert.Equal(t, 62, output)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputData(18)
	output := Process(input, 1)

	assert.Equal(t, 48400, output)
}

func TestPart2Example(t *testing.T) {
	output := Process(strings.Split(example, "\n"), 2)
	assert.Equal(t, 952408144115, output)
}
func TestPart2(t *testing.T) {
	input, _ := aoc.GetInputData(18)
	output := Process(input, 2)

	assert.Equal(t, 72811019847283, output)
}
