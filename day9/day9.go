package day9

import (
	"strconv"
	"strings"
)

type History struct {
	history [][]int
}

func (h *History) diffTree() {
	layer := 0
	for length := len(h.history[0]) - 1; length > 0; length-- {
		diff := make([]int, length)
		zeroCount := 0
		for i := 0; i < length; i++ {
			diff[i] = h.history[layer][i+1] - h.history[layer][i]
			if diff[i] == 0 {
				zeroCount++
			}
		}
		h.history = append(h.history, diff)
		layer++
		if zeroCount == length {
			break
		}
	}
}

func (h *History) fillNext() {
	lastLayerIdx := len(h.history) - 1
	for layer := lastLayerIdx; layer >= 0; layer-- {
		if layer == lastLayerIdx {
			h.history[layer] = append(h.history[layer], 0)
			continue
		}
		lastIndex := len(h.history[layer]) - 1
		a := h.history[layer][lastIndex]
		b := h.history[layer+1][lastIndex]
		h.history[layer] = append(h.history[layer], a+b)
	}
}

func Part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		history := parseHistory(line)
		history.diffTree()
		history.fillNext()
		sum += history.history[0][len(history.history[0])-1]
	}

	return sum
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
