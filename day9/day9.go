package day9

import (
	"strconv"
	"strings"
)

type History struct {
	history [][]int
}

func Part1(lines []string) int {
	histories := make([]History, len(lines))
	for i, line := range lines {
		histories[i] = parseHistory(line)
	}
	for _, history := range histories {
		fillHistory(history)
	}

	return -1
}

func Part2(lines []string) int {
	return -1
}

func parseHistory(line string) History {
	pointStrs := strings.Fields(line)
	dataPoints := make([]int, len(pointStrs))
	for i, pointStr := range pointStrs {
		v, _ := strconv.Atoi(pointStr)
		dataPoints[i] = v
	}
	return History{
		[][]int{dataPoints},
	}
}

func fillHistory(history History) {

}
