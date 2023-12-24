package day14

import (
	"aoc/aoc_util"
	"fmt"
)

func Part1(lines []string) int {
	value := 0
	// grid := make([][]byte, len(lines)+2)
	grid := make([][]byte, len(lines))
	// Prefixing and suffixing entire grid with a wall of #
	// grid[0] = make([]byte, len(lines[0]))
	// grid[len(grid)-1] = make([]byte, len(lines[0]))
	// for i, _ := range grid[0] {
	// 	grid[0][i] = '#'
	// 	grid[len(grid)-1][i] = '#'
	// }
	for i, line := range lines {
		grid[i] = make([]byte, len(line))
		for byteIdx, b := range []byte(line) {
			grid[i][byteIdx] = b
		}
	}

	// transposed := aoc_util.TransposeNewMatrix(grid)

	groups := make([]int, len(grid)*len(grid[0])*4) // 4 rotations
	groupLen := 0
	for y := 0; y < len(grid); y++ {
		// flattening the horizontal groups for later processing
		unclosedGroup := true
		for x := 0; x < len(grid[0]); x++ {
			if unclosedGroup && grid[y][x] != '#' {
				groups[groupLen] = x
				groups[groupLen+1] = y
				groupLen += 2
				unclosedGroup = !unclosedGroup
			}
		}
		if unclosedGroup { // is this closing off unclosed rows?
			groups[groupLen] = len(grid[0])
			groups[groupLen+1] = y
			groupLen += 2
		}
	}

	// grid = aoc_util.RotateNewMatrixCCW(grid)
	// value = shiftGrid(grid)

	// for _, bytes := range grid {
	// 	value += recurse(bytes, 0)
	// }
	// fmt.Printf("Value: %d\n", value)
	return value
}

func Part2(lines []string) int {
	value := 0
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = make([]byte, len(line))
		for byteIdx, b := range []byte(line) {
			grid[i][byteIdx] = b
		}
	}

	// transposed := aoc_util.TransposeNewMatrix(grid)

	// grid = aoc_util.RotateNewMatrixCCW(grid)
	// value += shiftGrid(grid)

	for iterations := 0; iterations < 1; iterations++ {
		if iterations%1_000_000 == 0 {
			fmt.Printf("iteration %d\n", iterations)
		}

		// tilt north and west
		// rockCol, rockRow := len(grid[0])-1, 0
		// nextCol := make([]byte, len(grid[0]))
		// for i := 0; i < len(grid[0]); i++ {
		//
		// }
		aoc_util.PrintMatrix(grid)
		fmt.Printf("Transpose\n")
		aoc_util.TransposeMatrix(grid)
		aoc_util.PrintMatrix(grid)
		fmt.Printf("Shift\n")
		shiftGrid(grid, true)
		// shiftGrid(grid, false)
		aoc_util.PrintMatrix(grid)

		// // technically a flip
		// aoc_util.RotateMatrixCW(grid)
		// shiftGrid(grid)
		// aoc_util.RotateMatrixCW(grid)
		// shiftGrid(grid)
		// aoc_util.RotateMatrixCW(grid)
		// shiftGrid(grid)
		// aoc_util.RotateMatrixCW(grid)
		// value += shiftGrid(grid)
	}

	fmt.Printf("Value: %d\n", value)
	return value
}

/*
value from index i with n rock, 2 rocks, pos 1 aka start, rocks start with index 0
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
	for i := lastRock + 1; i < len(slice); i++ {
		switch slice[i] {
		case '#': // matches group slice[:]
			if count > 0 {
				sum += sumGroup(len(slice), lastRock, count)
			}
			return sum + recurse(slice, i)
		case 'O':
			count++
		}
	}
	// handle in case it doesn't end with a rock
	return sum + sumGroup(len(slice), lastRock, count)
}

type cacheData struct {
	original, transformed string
	pos, value            int
}

var cache = make(map[string]cacheData)

func shiftGrid(grid [][]byte, left bool) int {
	cache = make(map[string]cacheData)
	value := 0

	for _, bytes := range grid {
		fmt.Printf("line pre : %s\n", string(bytes))
		if left {
			shiftLeft(bytes)
		} else {
			shiftRight(bytes)
		}
		fmt.Printf("line post: %s\n", string(bytes))
	}

	return value
}

func sumGroup(sliceLen, lastRock, count int) int {
	i := 0
	for n := 0; n < count; n++ {
		i += sliceLen - 2 - lastRock - n
	}
	return i
}

func shiftLeft(bytes []byte) {
	var rocks []int
	lastRock := len(bytes) + 1
	for i := len(bytes) - 1; i >= 0; i-- {
		switch bytes[i] {
		case '#':
			if len(rocks) > 0 {
				for j := i; j < len(bytes) && len(rocks) > 0; j++ {
					rockIndex := rocks[len(rocks)-1]
					switch bytes[j+1] {
					case 'O':
						if rockIndex == j+1 {
							rocks = rocks[:len(rocks)-1]
						}
					case '.':
						bytes[j+1] = 'O'
						bytes[rockIndex] = '.'
						rocks = rocks[:len(rocks)-1]
					}
				}
				rocks = nil
			}
			lastRock = i
			break
		case 'O':
			rocks = append(rocks, i)
		}
	}

	if len(rocks) > 0 {
		for j := 0; j < lastRock && len(rocks) > 0; j++ {
			rockIndex := rocks[len(rocks)-1]
			switch bytes[j] {
			case 'O':
				if rockIndex == j {
					rocks = rocks[:len(rocks)-1]
				}
			case '.':
				bytes[j] = 'O'
				bytes[rockIndex] = '.'
				rocks = rocks[:len(rocks)-1]
			}
		}
		rocks = nil
	}
}

func shiftRight(bytes []byte) {
	var rocks []int
	lastRock := len(bytes) - 1
	for i := 0; i < len(bytes); i++ {
		switch bytes[i] {
		case '#':
			if len(rocks) > 0 {
				shiftedRocks := 0
				for j := i - 1; j >= 0 && shiftedRocks < len(rocks); j-- {
					if bytes[j] == '.' {
						bytes[j] = 'O'
						bytes[rocks[len(rocks)-shiftedRocks-1]] = '.'
						shiftedRocks++
					}
				}
				rocks = []int{}
			}
			lastRock = i
			break
		case 'O':
			rocks = append(rocks, i)
		}
	}

	if len(rocks) > 0 {
		shiftedRocks := 0
		for n := len(bytes) - 1; n >= lastRock && shiftedRocks < len(rocks); n-- {
			rockIndex := rocks[len(rocks)-shiftedRocks-1]
			switch bytes[n] {
			case 'O':
				if rockIndex == n {
					shiftedRocks++
				}
			case '.':
				bytes[n] = 'O'
				bytes[rockIndex] = '.'
				shiftedRocks++
			}
		}
		rocks = nil
	}
}
