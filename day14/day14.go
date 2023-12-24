package day14

import (
	"aoc/aoc_util"
	"fmt"
	"strings"
)

func Process(data string) int {
	dataSets := strings.Split(data, "\n\n")

	gridCount, value := 0, 0
	for _, dataSet := range dataSets {
		var grid [][]byte
		lines := strings.Split(dataSet, "\n")
		for i, line := range lines {
			grid = append(grid, make([]byte, len(line)))
			for byteIdx, b := range []byte(line) {
				grid[i][byteIdx] = b
			}
		}
		fmt.Printf("Original\n")
		aoc_util.PrintMatrix(grid)

		grid = aoc_util.TransposeMatrix(grid)

		fmt.Printf("Transposed\n")
		aoc_util.PrintMatrix(grid)

		/*
		   		   issue shifting like this is now I have to iterate through again to calculate...

		   		   Should probably append a # stone so that there is not the issue of not ending on a closed rock (or put a barrier)
		   		   OO.O.O....#..OO -> OO.O.O....# & ..OO
		   		   both OO.O.O....# and OOO.....O# are the same 'value'. So would OOOO.# but that's a different subproblem of length

		           **go from RTL and prefix with #
		   		   they both shift to OOOO......#
		*/
		for x := 0; x < len(grid); x++ {
			cubeRockPos := -1
			row := grid[x]
			for y := 0; y < len(row); y++ {
				switch row[y] {
				case 'O':
					if cubeRockPos+1 < y {
						row[cubeRockPos+1] = 'O'
						row[y] = '.'
					}
					cubeRockPos++
				case '#':
					cubeRockPos = y
				}
			}

		}

		fmt.Printf("Shifted\n")
		aoc_util.PrintMatrix(grid)

		grid = aoc_util.TransposeMatrix(grid)
		fmt.Printf("Transposed Shifted\n")
		aoc_util.PrintMatrix(grid)

		// value += result
		grid = [][]byte{}
		gridCount++
	}
	return value
}

func recurse(slice []byte, i, j int) {

}
