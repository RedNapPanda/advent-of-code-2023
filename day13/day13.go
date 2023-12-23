package day13

import "fmt"

func Process(lines []string) int {

	// var colCache map[int]int
	// var rowCache map[int]int

	grids := parseLines(lines)
	fmt.Printf("grids: %d\n", len(grids))

	return 0
}

/*
parseLines
splits lines into i j k where

	i => grid ordinal
	j => row
	k => column
*/
func parseLines(lines []string) [][][]byte {
	i := 0
	var results [][][]byte
	var grid [][]byte

	flushGrid := func() {
		results = append(results, grid)
		grid = [][]byte{}
	}

	for _, line := range lines {
		grid = append(grid, make([]byte, len(lines[0])))
		for byteIdx, b := range []byte(line) {
			grid[i][byteIdx] = b
		}
		i++
		if line == "" {
			i = 0
			flushGrid()
		}
	}
	flushGrid()
	return results
}
