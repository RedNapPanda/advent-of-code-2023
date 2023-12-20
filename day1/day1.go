package day1

import (
	"runtime"
	"strings"
)

type processor func(c chan<- int, line string)

func Process(lines []string, fn processor) int {
	runtime.GOMAXPROCS(8)
	c := make(chan int, 16)
	defer close(c)
	sum := 0
	for index := range lines {
		go fn(c, lines[index])
		sum += <-c
	}
	return sum
}

func Part1(c chan<- int, line string) {
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
	c <- (first * 10) + second
}

func Part2(c chan<- int, line string) {
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
	c <- (first * 10) + second
}

func isAsciiDigit(char uint8) bool {
	return char >= 48 && char <= 57
}

func toAsciiDigit(char uint8) int {
	return int(char - 48)
}

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

func isTextNumber(s string, fn func(string, string) bool) int {
	for index, number := range numbers {
		if fn(s, number) {
			return index + 1
		}
	}
	return 0
}
