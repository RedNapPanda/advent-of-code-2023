package day11

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var example = []string{
	"...#......",
	".......#..",
	"#.........",
	"..........",
	"......#...",
	".#........",
	".........#",
	"..........",
	".......#..",
	"#...#.....",
}

func TestPart1Example(t *testing.T) {
	output := Process(example)

	assert.Equal(t, -1, output)
}
