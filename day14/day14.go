package day14

import (
	"aoc/aoc_util"
	"fmt"
)

func Process(lines []string) int {
	value := 0
	grid := make([][]byte, len(lines)+2)
	// Prefixing an suffixing entire grid with a wall of #
	grid[0] = make([]byte, len(lines[0]))
	grid[len(grid)-1] = make([]byte, len(lines[0]))
	for i, _ := range grid[0] {
		grid[0][i] = '#'
		grid[len(grid)-1][i] = '#'
	}
	for x, line := range lines {
		grid[x+1] = make([]byte, len(line))
		for byteIdx, b := range []byte(line) {
			grid[x+1][byteIdx] = b
		}
	}

	aoc_util.PrintMatrix(grid)

	// grid = aoc_util.TransposeMatrix(grid)
	grid = aoc_util.RotateMatrixCCW(grid)

	aoc_util.PrintMatrix(grid)

	for i, bytes := range grid {
		val := recurse(bytes, 0)
		fmt.Printf("%d | Value %d\n", i, val)
		value += val
	}

	return value
}

/*
value from index i with n rock, 2 rocks, position 1 aka start, rocks start with index 0
11 - i - n
11 - 1 - 0 => 10 + 0 => 10
11 - 1 - 1 => 11 - 2 => 9
Starts from 0, counts O rocks and finds first # rock.

	sum = calculates the value of the n rocks with that lastRock index
	returns sum + recurse, new lastRock index

This doesn't actually shift anything
for n rock

	sum += 11 - i - n - 1
*/
func recurse(slice []byte, lastRock int) int {
	count, sum := 0, 0
	sumGroup := func() int {
		i := 0
		for n := 0; n < count; n++ {
			i += len(slice) - lastRock - n - 2
		}
		return i
	}
	for i := lastRock + 1; i < len(slice); i++ {
		switch slice[i] {
		case '#': // matches group slice[:]
			if count > 0 {
				sum += sumGroup()
			}
			return sum + recurse(slice, i)
		case 'O':
			count++
		}
	}
	// handle in case it doesn't end with a rock
	return sum + sumGroup()
}

func shiftGrid(grid [][]byte) {
	/*
	   		   issue shifting like this is now I have to iterate through again to calculate...

	   		   Should probably append a # stone so that there is not the issue of not ending on a closed rock (or put a barrier)
	   		   OO.O.O....#..OO -> OO.O.O....# & ..OO
	   		   both OO.O.O....# and OOO.....O# are the same 'value'. So would OOOO.# but that's a different subproblem of length

	           **go from RTL and prefix with #
	   		   they both shift to OOOO......#
	*/

	// shift north
	count, cubeRockPos := 0, 0
	for y := 0; y < len(grid[0]); y++ {
		for x := 0; x < len(grid); x++ {
			// copyBytes := make([]byte, len(grid))
			switch grid[x][y] {
			case 'O':
				if cubeRockPos+1 < y {
					grid[cubeRockPos+1][y] = 'O'
					grid[x][y] = '.'
				}
				cubeRockPos++
				count++
			case '#':
				cubeRockPos = x
			}
		}
	}

	// This assumes the grid is shifting to the left
	// for x := 0; x < len(grid); x++ {
	// 	row := grid[x]
	// 	for y := 0; y < len(row); y++ {
	// 		switch row[y] {
	// 		case 'O':
	// 			if cubeRockPos+1 < y {
	// 				row[cubeRockPos+1] = 'O'
	// 				row[y] = '.'
	// 			}
	// 			cubeRockPos++
	// 		case '#':
	// 			cubeRockPos = y
	// 		}
	// 	}
	// }
}
