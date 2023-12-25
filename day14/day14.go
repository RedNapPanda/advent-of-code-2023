package day14

import (
	"aoc/aoc_util"
	"crypto/sha256"
	"fmt"
)

var cache = make(map[string]int)
var hashBytes []byte

func Part1(lines []string) int {
	value := 0
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = make([]byte, len(line))
		for byteIdx, b := range []byte(line) {
			grid[i][byteIdx] = b
		}
	}

	// transpose for recursion
	aoc_util.TransposeMatrix(grid)
	value = 0
	for _, bytes := range grid {
		value += scoreTiltWestWithoutShift(bytes, -1)
	}
	fmt.Printf("Recursive Value: %d\n", value)
	aoc_util.TransposeMatrix(grid)
	// transpose after to flip back since it's unmodified

	shiftNorth(grid)

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
	fmt.Printf("Transpose Shifted Value: %d\n", value)

	return value
}

/*
value from index i with n rock, 2 rocks, pos 1 aka start, rocks start with index 0
11 - i - n
11 - 1 - 0 => 10 + 0 => 10
11 - 1 - 1 => 11 - 2 => 9
Starts from 0, counts O rocks and finds first # rock.

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
	fmt.Printf("Value: %d\n", value)
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
	// key := keyGrid(grid)
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

func shiftSouth(grid [][]byte) {
	for col := 0; col < len(grid); col++ {
		lastRock := len(grid) - 1
		for row := len(grid) - 1; row >= 0; row-- {
			switch grid[row][col] {
			case '#':
				lastRock = row - 1
			case 'O':
				lastRock -= 1
				grid[row][col] = '.'
				grid[lastRock+1][col] = 'O'
			}
		}
	}
}

func shiftWest(grid [][]byte) {
	openSpace := 0
	var bytes []byte
	for row := 0; row < len(grid); row++ {
		bytes = grid[row]
		for col := 0; col < len(bytes); col++ {
			switch bytes[col] {
			case 'O':
				if col != openSpace {
					bytes[openSpace] = 'O'
					bytes[col] = '.'
				}
				openSpace = max(openSpace+1, 0)
			case '#':
				openSpace = max(col+1, 0)
			}
		}
	}
}

func shiftEast(grid [][]byte) {
	var openSpace int
	var bytes []byte
	for row := 0; row < len(grid); row++ {
		bytes = grid[row]
		openSpace = len(bytes) - 1
		for col := len(bytes) - 1; col >= 0; col-- {
			switch bytes[col] {
			case 'O':
				if col != openSpace {
					bytes[openSpace] = 'O'
					bytes[col] = '.'
				}
				openSpace = max(openSpace-1, 0)
			case '#':
				openSpace = max(col-1, 0)
			}
		}
	}
}
