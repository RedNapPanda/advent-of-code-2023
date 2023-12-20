package advent_of_code

import (
	"bufio"
	"fmt"
	"os"
)

func GetInputData(day int) ([]string, error) {
	if day < 1 || day > 25 {
		return nil, fmt.Errorf("incorrect day or part number: day %d", day)
	}
	path := fmt.Sprintf("C:/Users/excel/Documents/projects/dedi-servers/home-dedi/go/advent-of-code/inputs/day%d.txt", day)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
