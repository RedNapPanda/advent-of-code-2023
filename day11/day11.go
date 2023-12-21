package day11

import "fmt"

func Process(lines []string) int {
	var matrix [][]byte
	inc := 0
	// parse rows and insert extra empty rows
	for x, line := range lines {
		matrix = append(matrix, make([]byte, len(line)))
		isEmpty := true
		for y, c := range []byte(line) {
			matrix[x+inc][y] = c
			if c == '#' {
				isEmpty = false
			}
		}
		if isEmpty {
			slice := matrix[x+inc]
			matrix = append(matrix, slice)
			inc++
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

				}
				list[y] = '.'
				matrix[x] = list
			}
			skipNext = true
			inc++
		}
	}

	for _, bytes := range matrix {
		fmt.Printf("%+v\n", string(bytes))
	}
	fmt.Printf("%d\n", len(matrix))

	return 0
}
