package day15

import (
	aoc "aoc"
	"aoc/aoc_util"
	"github.com/stretchr/testify/assert"
	"testing"
)

var example = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"

func TestPart1Example(t *testing.T) {
	defer aoc_util.Timer("TestPart1Example")()
	output := Process(example)
	assert.Equal(t, 1320, output)
}

func TestPart1ExampleRecursive(t *testing.T) {
	assert.Equal(t, 129, calculate('q', 0))

	defer aoc_util.Timer("TestPart1ExampleRecursive")()
	output := RecursiveProcess(example)
	assert.Equal(t, 1320, output)
}

func TestPart1ExampleHASH(t *testing.T) {
	output := Process("HASH")
	assert.Equal(t, 52, output)
}

func TestPart1ExampleHASHRecursive(t *testing.T) {
	assert.Equal(t, 129, calculate('q', 0))

	output := RecursiveProcess("HASH")
	assert.Equal(t, 52, output)
}

// for loop is ~500µs
func TestPart1(t *testing.T) {
	defer aoc_util.Timer("TestPart1")()
	input, _ := aoc.GetInputDataString(15)
	output := Process(input)

	assert.Equal(t, 508498, output)
}

// recursion is ~5ms, cache brings this down to 2.7ms lmao
func TestPart1Recursive(t *testing.T) {
	defer aoc_util.Timer("TestPart1Recursive")()
	input, _ := aoc.GetInputDataString(15)
	output := RecursiveProcess(input)

	assert.Equal(t, 508498, output)
}

func TestPart2Example(t *testing.T) {
	defer aoc_util.Timer("TestPart2Example")()
	output := Process2(example)
	assert.Equal(t, 145, output)
}

func TestPart2(t *testing.T) {
	defer aoc_util.Timer("TestPart2")()
	input, _ := aoc.GetInputDataString(15)
	output := Process2(input)

	assert.Equal(t, 279116, output)
}
