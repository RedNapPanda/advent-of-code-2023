package day1

import (
	"strings"
)

type processor func(line string) int
type isNumber func(string, string) bool

var numbers = [...]string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func Process(lines []string, fn processor) int {
	sum := 0
	for index := range lines {
		sum += fn(lines[index])
	}
	return sum
}

func Part1(line string) int {
	first, second := 0, 0
	for j := 0; j < len(line); j++ {
		if first != 0 && second != 0 {
			break
		}
		char := line[j]
		if first == 0 && isAsciiDigit(char) {
			first = toAsciiDigit(char)
		}
		char = line[len(line)-(j+1)]
		if second == 0 && isAsciiDigit(char) {
			second = toAsciiDigit(char)
		}
	}
	return (first * 10) + second
}

func Part2(line string) int {
	first, second := 0, 0
	for j := 0; j < len(line); j++ {
		if first != 0 && second != 0 {
			break
		}
		char := line[j]

		if first == 0 {
			if isAsciiDigit(char) {
				first = toAsciiDigit(char)
			} else if digit := isTextNumber(line[j:], strings.HasPrefix); digit != 0 {
				first = digit
			}
		}
		secondIndex := len(line) - (j + 1)
		if second == 0 {
			char = line[secondIndex]
			if isAsciiDigit(char) {
				second = toAsciiDigit(char)
			} else if digit := isTextNumber(line[:secondIndex+1], strings.HasSuffix); digit != 0 {
				second = digit
			}
		}
	}
	return (first * 10) + second
}

func isAsciiDigit(char byte) bool {
	return char >= 48 && char <= 57
}

func toAsciiDigit(char byte) int {
	return int(char - 48)
}

func isTextNumber(s string, fn isNumber) int {
	for index, number := range numbers {
		if fn(s, number) {
			return index + 1
		}
	}
	return 0
}
