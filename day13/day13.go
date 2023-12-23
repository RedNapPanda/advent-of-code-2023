package day13

import (
	"aoc/aoc_util"
	"slices"
	"strings"
)

func Process(data string) int {
	dataSets := strings.Split(strings.TrimSuffix(data, "\n"), "\n\n")
	var grid [][]byte

	i, gridCount, value := 0, 0, 0
	for _, dataSet := range dataSets {
		lines := strings.Split(dataSet, "\n")
		for _, line := range lines {
			grid = append(grid, make([]byte, len(line)))
			for byteIdx, b := range []byte(line) {
				grid[i][byteIdx] = b
			}
			i++
		}
		value += findMirror(grid)
		grid = [][]byte{}
		gridCount++
		i = 0
	}
	return value
}

func findMirror(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	mirrorCol, mirrorRow := 0, 0
	// rows
	for i := 0; i < len(grid); i++ {
		upper := make([][]byte, i+1)
		copy(upper, grid[:i+1])
		slices.Reverse(upper)
		lower := grid[i+1:]
		rows := min(len(upper), len(lower))
		upper = upper[:rows]
		lower = lower[:rows]

		matches := true
		for x := 0; x < len(upper); x++ {
			matches = aoc_util.SlicesEqual(upper[x], lower[x])
			if !matches {
				break
			}
		}
		if matches && i < len(grid)-1 {
			mirrorRow = i + 1
			break
		}
	}

	// TransposeMatrix by X, Y
	transposed := aoc_util.TransposeMatrix(grid)
	// columns
	for i := 0; i < len(transposed); i++ {
		upper := make([][]byte, i+1)
		copy(upper, transposed[:i+1])
		slices.Reverse(upper)
		lower := transposed[i+1:]
		rows := min(len(upper), len(lower))
		upper = upper[:rows]
		lower = lower[:rows]

		matches := true
		for x := 0; x < len(upper); x++ {
			matches = aoc_util.SlicesEqual(upper[x], lower[x])
			if !matches {
				break
			}
		}
		if matches && i < len(transposed)-1 {
			mirrorCol = i + 1
			break
		}
	}
	return mirrorCol + 100*mirrorRow
}
