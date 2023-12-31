package day14

import (
	"aoc/aoc_util"
	"crypto/sha256"
	"fmt"
)

var cache = make(map[string]int)
var hashBytes []byte

func Part1(lines []string) int {
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = make([]byte, len(line))
		for byteIdx, b := range []byte(line) {
			grid[i][byteIdx] = b
		}
	}

	// transpose for recursion
	aoc_util.TransposeMatrix(grid)
	rvalue := 0
	for _, bytes := range grid {
		rvalue += scoreTiltWestWithoutShift(bytes, -1)
	}
	aoc_util.TransposeMatrix(grid)
	// transpose after to flip back since it's unmodified
	shiftNorth(grid)

	value := 0
	for n := 0; n < len(grid); n++ {
		count := 0
		for _, b := range grid[n] {
			if b == 'O' {
				count++
			}
		}
		value += count * (len(grid) - n)
	}

	if rvalue != value {
		return -1
	}

	return value
}

/*
value from index i with n rock, 2 rocks, pos 1 aka start, rocks start with index 0
11 - i - n
11 - 1 - 0 => 10 + 0 => 10
11 - 1 - 1 => 11 - 2 => 9
Starts from 0, counts W rocks and finds first # rock.

	sum = calculates the value of the n rocks with that lastRock index
	returns sum + scoreTiltWestWithoutShift, new lastRock index

This doesn't actually shift anything
for n rock

	sum += 11 - i - n - 1
*/
func scoreTiltWestWithoutShift(slice []byte, lastRock int) int {
	count, sum := 0, 0
	for i := lastRock + 1; i < len(slice); i++ {
		switch slice[i] {
		case '#': // matches group slice[:]
			if count > 0 {
				sum += sumGroup(len(slice), lastRock, count)
			}
			return sum + scoreTiltWestWithoutShift(slice, i)
		case 'O':
			count++
		}
	}
	// handle in case it doesn't end with a rock
	return sumGroup(len(slice), lastRock, count)
}

func sumGroup(sliceLen, lastRock, count int) int {
	i := 0
	for n := 0; n < count; n++ {
		i += sliceLen - lastRock - n - 1
	}
	return i
}

func Part2(lines []string, iterations int) int {
	value := 0
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = make([]byte, len(line))
		for byteIdx, b := range []byte(line) {
			grid[i][byteIdx] = b
		}
	}

	cycleTilts(grid, iterations)

	value = 0
	for n := 0; n < len(grid); n++ {
		count := 0
		for _, b := range grid[n] {
			if b == 'O' {
				count++
			}
		}
		value += count * (len(grid) - n)
	}
	return value
}

/*
iterate till n count
if cached, break out of tilting since we have this value already
cycleLen = current iteration - last iteration
iterationsLeft - i % cycleLen = remaining single cycles remaining (not group)
*/
func cycleTilts(grid [][]byte, iterations int) {
	if iterations == 0 {
		return
	}
	cycleLen, i := 0, 0
	var key string
	cycle := func() {
		for n := 0; n < 4; n++ {
			shiftNorth(grid)
			aoc_util.RotateMatrixCW(grid)
		}
	}
	for ; i < iterations; i++ {
		key = hashGrid(grid)
		if _, ok := cache[key]; ok {
			cycleLen = i - cache[key]
			fmt.Printf("Cycle found starting at index %d, length: %d\n", cache[key], cycleLen)
			break
		}
		cache[key] = i
		cycle()
	}
	if cycleLen > 0 {
		remainder := (iterations - i) % cycleLen
		for m := 0; m < remainder; m++ {
			cycle()
		}
		return
	}
}

func hashGrid(grid [][]byte) string {
	hashBytes = make([]byte, len(grid)*len(grid[0]))
	for i := 0; i < len(grid); i++ {
		iShift := i * len(grid[i])
		for j := 0; j < len(grid[i]); j++ {
			hashBytes[iShift+j] = grid[i][j]
		}
	}
	sha := sha256.Sum256(hashBytes)
	return string(sha[:])
}

func shiftNorth(grid [][]byte) {
	for col := 0; col < len(grid); col++ {
		lastRock := 0
		for row := 0; row < len(grid); row++ {
			switch grid[row][col] {
			case '#':
				lastRock = row + 1
			case 'O':
				lastRock += 1
				grid[row][col] = '.'
				grid[lastRock-1][col] = 'O'
			}
		}
	}
}
