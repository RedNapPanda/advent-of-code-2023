package day3

import (
	"fmt"
	"strconv"
)

type Number struct {
	value  int
	row    int
	index  int
	length int
}

type Pair struct {
	first  int
	second int
}

func Part1(lines []string) int {
	sum := 0
	coords, matrix := fillMatrixAndCoords(lines)

	for _, coord := range coords {
		row := max(0, coord.row-1)
		for ; row < min(coord.row+2, len(matrix)); row++ {
			startIndex := max(0, coord.index-1)
			endIndex := min(len(matrix[row]), coord.index+coord.length+1)
			if hasSymbol(matrix[row][startIndex:endIndex]) {
				sum += coord.value
				break
			}
		}
	}

	return sum
}

func Part2(lines []string) int {
	sum := 0
	gearMap := make(map[string]*Pair)
	coords, matrix := fillMatrixAndCoords(lines)

	for _, coord := range coords {
		row := max(0, coord.row-1)
		for ; row < min(coord.row+2, len(matrix)); row++ {
			startIndex := max(0, coord.index-1)
			endIndex := min(len(matrix[row]), coord.index+coord.length+1)
			if exists, index := hasGear(matrix[row][startIndex:endIndex]); exists {
				index += startIndex
				key := fmt.Sprintf("%d;%d", row, index)
				gear, ok := gearMap[key]
				if ok {
					gear.second = coord.value
				} else {
					gearMap[key] = &Pair{first: coord.value}
				}
			}
		}
	}

	for _, pair := range gearMap {
		if pair.first != 0 && pair.second != 0 {
			sum += pair.first * pair.second
		}
	}

	return sum
}

func fillMatrixAndCoords(lines []string) ([]Number, [][]byte) {
	var matrix [][]byte
	var coords []Number
	for row := 0; row < len(lines); row++ {
		bytes := []byte(lines[row])
		matrix = append(matrix, bytes)
		lineCoords := parseLine(bytes, row)
		coords = append(coords, lineCoords...)
	}
	return coords, matrix
}

func parseLine(bytes []byte, row int) []Number {
	var coords []Number
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == '.' {
			continue
		}
		var digits []byte = nil
		startIndex := 0
		if isAsciiDigit(bytes[i]) {
			startIndex = i
			digits = []byte{bytes[i]}
			for index := i + 1; index < len(bytes) && isAsciiDigit(bytes[index]); index++ {
				digits = append(digits, bytes[index])
			}
		}
		if digits != nil {
			digitStr := string(digits)
			value, _ := strconv.Atoi(digitStr)
			coords = append(coords, Number{
				value:  value,
				row:    row,
				index:  startIndex,
				length: len(digitStr),
			})
			i += len(digitStr) - 1
		}
	}
	return coords
}

func hasGear(bytes []byte) (bool, int) {
	for index, b := range bytes {
		if isGear(b) {
			return true, index
		}
	}
	return false, -1
}

func hasSymbol(bytes []byte) bool {
	for _, b := range bytes {
		if isSymbol(b) {
			return true
		}
	}
	return false
}

func isGear(char byte) bool {
	return char == '*'
}

func isSymbol(char byte) bool {
	return char != '.' && !isAsciiDigit(char)
}

func isAsciiDigit(char byte) bool {
	return char >= 48 && char <= 57
}
