package advent_of_code

import (
	"runtime"
	"strings"
	"sync"
)

type processor func(c chan<- int, line string)

func Process(lines []string, fn processor) int {
	runtime.GOMAXPROCS(8)
	var wg sync.WaitGroup
	c := make(chan int, 16)
	sum := 0
	for index := range lines {
		wg.Add(1)
		go fn(c, lines[index])
		sum += <-c
	}
	return sum
}

func ProcessPart1(c chan<- int, line string) {
	first, second := 0, 0
	for j := 0; j < len(line); j++ {
		if first != 0 && second != 0 {
			break
		}
		char := line[j]
		if first == 0 && char >= 48 && char <= 57 {
			first = int(char - 48)
		}
		char = line[len(line)-(j+1)]
		if second == 0 && char >= 48 && char <= 57 {
			second = int(char - 48)
		}
	}
	c <- (first * 10) + second
}

func ProcessPart2(c chan<- int, line string) {
	first, second := 0, 0
	for j := 0; j < len(line); j++ {
		if first != 0 && second != 0 {
			break
		}
		char := line[j]

		if first == 0 {
			if char >= 48 && char <= 57 {
				first = int(char - 48)
			} else if digit := isTextNumber(line[j:], strings.HasPrefix); digit != 0 {
				first = digit
			}
		}
		secondIndex := len(line) - (j + 1)
		char = line[secondIndex]
		if second == 0 {
			if char >= 48 && char <= 57 {
				second = int(char - 48)
			} else if digit := isTextNumber(line[:secondIndex+1], strings.HasSuffix); digit != 0 {
				second = digit
			}
		}
	}
	c <- (first * 10) + second
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
