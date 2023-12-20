package advent_of_code

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1Example(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	output := Process(input, ProcessPart1)
	assert.Equal(t, 142, output)
}

func TestPart1(t *testing.T) {
	input, _ := GetInputData(1)
	output := Process(input, ProcessPart1)
	fmt.Printf("Sum: %d\n", output)
	assert.Equal(t, 54450, output)
}

func TestPart2Example(t *testing.T) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	output := Process(input, ProcessPart2)
	assert.Equal(t, 281, output)
}

func TestPart2(t *testing.T) {
	input, _ := GetInputData(1)
	output := Process(input, ProcessPart2)
	fmt.Printf("Sum: %d\n", output)
	assert.Equal(t, 54265, output)
}
