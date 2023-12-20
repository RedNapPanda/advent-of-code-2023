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
	pwd, _ := os.Getwd()
	path := pwd + fmt.Sprintf("/inputs/day%d/input.txt", day)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
