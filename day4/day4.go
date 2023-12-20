package day4

import (
	"strings"
)

type Hand struct {
	cards   []string
	winners []string
}

func Part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		hand := parseLine(line)
		intersect := make(map[string]string)
		for _, v := range hand.winners {
			intersect[v] = v
		}
		value := 0
		for _, v := range hand.cards {
			if _, ok := intersect[v]; ok {
				if value == 0 {
					value = 1
				} else {
					value *= 2
				}
			}
		}
		sum += value
	}
	return sum
}

func Part2(lines []string) int {
	copies := make(map[int]int)
	sum := 0
	for i, line := range lines {
		hand := parseLine(line)
		intersect := make(map[string]string)
		for _, v := range hand.winners {
			intersect[v] = v
		}

		add(i, 1, copies)
		matches := 0
		for _, v := range hand.cards {
			if _, ok := intersect[v]; ok {
				matches++
			}
		}
		if currentCopies, ok := copies[i]; ok {
			for j := 1; j <= matches; j++ {
				add(i+j, currentCopies, copies)
			}
		}
	}
	for _, v := range copies {
		sum += v
	}
	return sum
}

func parseLine(line string) Hand {
	prefix := strings.Split(line, ":")
	pipe := strings.Split(prefix[1], "|")
	var cards, winners []string
	for _, winner := range strings.Fields(pipe[0]) {
		winners = append(winners, winner)
	}
	for _, card := range strings.Fields(pipe[1]) {
		cards = append(cards, card)
	}

	return Hand{
		cards:   cards,
		winners: winners,
	}
}

func add(i int, value int, copies map[int]int) {
	if _, ok := copies[i]; ok {
		copies[i] += value
	} else {
		copies[i] = value
	}
}
