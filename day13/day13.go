package day13

import (
	"aoc/aoc_util"
	"slices"
	"strings"
)

func Process(data string, part int) int {
	dataSets := strings.Split(data, "\n\n")
	var grid [][]byte

	gridCount, value := 0, 0
	for _, dataSet := range dataSets {
		lines := strings.Split(dataSet, "\n")
		for i, line := range lines {
			grid = append(grid, make([]byte, len(line)))
			for byteIdx, b := range []byte(line) {
				grid[i][byteIdx] = b
			}
		}
		value += findMirror(grid, part-1)
		grid = [][]byte{}
		gridCount++
	}
	return value
}

func findMirror(grid [][]byte, smudges int) int {
	if len(grid) == 0 {
		return 0
	}

	mirrorCol, mirrorRow := 0, 0
	// rows
	for i := 0; i < len(grid)-1; i++ {
		upper := make([][]byte, i+1)
		copy(upper, grid[:i+1])
		slices.Reverse(upper)
		lower := grid[i+1:]
		rows := min(len(upper), len(lower))
		upper = upper[:rows]
		lower = lower[:rows]

		value := 0
	rowMatching:
		for x := 0; x < len(upper); x++ {
			// matches = aoc_util.SlicesEqual(upper[x], lower[x]) // Part 1
			// converted the function that checked for all equality to count for mismatches.  Breaks if > smudges
			for y, v := range upper[x] {
				if lower[x][y] != v {
					value++
					if value > smudges {
						break rowMatching
					}
				}
			}
		}
		if value == smudges {
			mirrorRow = i + 1
			break
		}
	}

	// TransposeMatrix by X, Y
	transposed := aoc_util.TransposeMatrix(grid)
	// columns
	for i := 0; i < len(transposed)-1; i++ {
		upper := transposed[:i+1]
		lower := transposed[i+1:]
		rows := min(len(upper), len(lower))
		upper = upper[len(upper)-rows:]
		lower = lower[:rows]

		value := 0
	colMatching:
		for x := 0; x < len(upper); x++ {
			// matches = aoc_util.SlicesEqual(upper[x], lower[x]) // Part 1
			for y, v := range upper[len(upper)-x-1] {
				if lower[x][y] != v {
					value++
					if value > smudges {
						break colMatching
					}
				}
			}
		}
		if value == smudges {
			mirrorCol = i + 1
			break
		}
	}
	return mirrorCol + 100*mirrorRow
}
