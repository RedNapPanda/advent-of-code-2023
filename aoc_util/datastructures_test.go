package aoc_util

import (
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

func TestFlatteningGridIteration(t *testing.T) {
	lines := strings.Split(example, "\n")
	gridLen := len(lines)
	grid := make([][]int, gridLen)
	flattened := NewFlattenedGrid[byte](gridLen*gridLen, gridLen)
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for byteIdx, b := range []byte(line) {
			grid[i][byteIdx] = int(b - 48)
			flattened.Data[(i*flattened.Offset)+byteIdx] = b
		}
	}
	PrintMatrix(grid)
	flattened.Print("%c", " ")
}
