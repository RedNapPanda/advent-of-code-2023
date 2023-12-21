package day11

import (
	"fmt"
)

type coord struct {
	x, y int
}

func Process(lines []string) int {
	matrix, galaxies := parseUniverse(lines)

	fmt.Printf("\n")
	for _, galaxy := range galaxies {
		fmt.Printf("%+v\n", galaxy)
	}

	fmt.Printf("\n")
	for _, bytes := range matrix {
		fmt.Printf("%+v\n", string(bytes))
	}
	return 0
}

func parseUniverse(lines []string) ([][]byte, []coord) {
	var galaxies []coord
	var matrix [][]byte
	var columnsToInc []int
	inc := 0
	// parse rows and insert extra empty rows
	for x, line := range lines {
		matrix = append(matrix, make([]byte, len(line)))
		isEmpty := true
		nilCoord := coord{-1, -1}
		galaxy := nilCoord
		for y, c := range []byte(line) {
			matrix[x+inc][y] = c
			if c == '#' {
				isEmpty = false
				galaxy = coord{x + inc, y}
			}
		}
		if isEmpty {
			slice := matrix[x+inc]
			matrix = append(matrix, slice)
			inc++
		}
		if galaxy != nilCoord {
			galaxies = append(galaxies, galaxy)
		}
	}
	inc = 0
	skipNext := false
	// find empty columns
	for y := 0; y < len(matrix[0]); y++ {
		if skipNext {
			skipNext = false
			continue
		}
		isEmpty := true
		for x := 0; x < len(matrix); x++ {
			c := matrix[x][y]
			if c == '#' {
				isEmpty = false
			}
		}
		if isEmpty {
			for x := 0; x < len(matrix); x++ {
				var list []byte
				if y == len(matrix[x]) {
					list = append(matrix[x], '.')
				} else {
					pre := matrix[x][:y+1]
					post := matrix[x][y:]
					list = append(pre, post...)

					columnsToInc = append(columnsToInc, y)
				}
				list[y] = '.'
				matrix[x] = list
			}
			for i, _ := range galaxies {
				galaxy := galaxies[i]
				if galaxy.y > y+1 {
					galaxy.y += 1
					galaxies[i] = galaxy
				}
			}
			skipNext = true
			inc++
		}
	}

	return matrix, galaxies
}
