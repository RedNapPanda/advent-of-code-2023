package day6

import (
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	var races [][]int
	timeSplit := strings.Fields(lines[0][5:])
	distSplit := strings.Fields(lines[1][9:])

	for i := 0; i < len(timeSplit); i++ {
		time, _ := strconv.Atoi(timeSplit[i])
		dist, _ := strconv.Atoi(distSplit[i])
		races = append(races, []int{time, dist})
	}

	product := 1
	for _, race := range races {
		wins := calculateWins(race[0], race[1])
		product *= wins
	}

	return product
}

func Part2(lines []string) int {
	time, _ := strconv.Atoi(strings.ReplaceAll(lines[0][5:], " ", ""))
	dist, _ := strconv.Atoi(strings.ReplaceAll(lines[1][9:], " ", ""))
	return calculateWins(time, dist)
}

func calculateWins(time int, dist int) int {
	wins := 0
	for speed := time; speed >= 0; speed-- {
		remainingTime := time - speed
		if remainingTime*speed > dist {
			wins++
		}
	}
	return wins
}
