package day17

import (
	aoc "aoc"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var example = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

/*
2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533
*/

func TestPart1Example(t *testing.T) {
	output := Process(strings.Split(example, "\n"), 1)
	assert.Equal(t, 102, output)
}

// --- PASS: TestPart1 (18.74s)
func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputData(17)
	output := Process(input, 1)

	assert.Equal(t, 1256, output)
}

func TestPart2Example1(t *testing.T) {
	output := Process(strings.Split(example, "\n"), 2)
	assert.Equal(t, 94, output)
}

func TestPart2Example2(t *testing.T) {
	example = `111111111111
999999999991
999999999991
999999999991
999999999991`
	output := Process(strings.Split(example, "\n"), 2)
	assert.Equal(t, 71, output)
}

func TestPart2(t *testing.T) {
	input, _ := aoc.GetInputData(17)
	output := Process(input, 2)

	assert.Equal(t, 1382, output)
}
