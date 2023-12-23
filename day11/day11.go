package day11

import "fmt"

type coord struct {
	x, y int
}

func Part1(lines []string) int {
	galaxies := parseUniverse(lines, 2)

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i; j < len(galaxies); j++ {
			a, b := galaxies[i], galaxies[j]
			if a != b {
				sum += manhattanDist(a, b)
			}
		}
	}

	fmt.Printf("\n")
	for _, galaxy := range galaxies {
		fmt.Printf("%+v\n", galaxy)
	}

	return sum
}

func Part2(lines []string, expansion int) int {
	galaxies := parseUniverse(lines, expansion)

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i; j < len(galaxies); j++ {
			a, b := galaxies[i], galaxies[j]
			if a != b {
				sum += manhattanDist(a, b)
			}
		}
	}

	return sum
}

func parseUniverse(lines []string, expand int) []coord {
	expand = expand - 1
	var galaxies []coord
	var matrix [][]byte
	inc := 0
	// parse rows and insert extra empty rows
	for x, line := range lines {
		matrix = append(matrix, make([]byte, len(line)))
		isEmpty := true
		nilCoord := coord{-1, -1}
		galaxy := nilCoord
		for y, c := range []byte(line) {
			matrix[x][y] = c
			if c == '#' {
				isEmpty = false
				galaxy = coord{x + inc, y}
				galaxies = append(galaxies, galaxy)
			}
		}
		if isEmpty {
			inc += expand
		}
	}
	var emptyColumns []int
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
			emptyColumns = append(emptyColumns, y)
			skipNext = true
		}
	}

	// expand y coords (x was done on creation, could also have been shifted here as well)
	for i, col := range emptyColumns {
		for x := range galaxies {
			expandedCol := col + i*expand
			if galaxies[x].y > expandedCol {
				galaxies[x].y += expand
			}
		}
	}

	return galaxies
}

// Credits: https://en.wikipedia.org/wiki/Taxicab_geometry
// NO hint this year? other than the not so obvious repeating down-right path from galaxy 5 to 9
// Path of down then right from 5 to 9 would have been more obvious
// Found this after going down the Dijkstra's algorithm route....no thanks
func manhattanDist(a, b coord) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}
