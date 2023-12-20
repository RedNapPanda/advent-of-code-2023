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

func (h *History) fill() {
	lastLayer := len(h.history) - 1
	for layer := lastLayer; layer >= 0; layer-- {
		if layer == lastLayer {
			// Prepend/Append hack since this layer is purely zeroes
			h.history[layer] = append(h.history[layer], 0, 0)
			continue
		}
		lastIndex := len(h.history[layer]) - 1
		nextA := h.history[layer][lastIndex]
		nextB := h.history[layer+1][lastIndex+1]
		prevA := h.history[layer][0]
		prevB := h.history[layer+1][0]
		h.history[layer] = append(h.history[layer], nextA+nextB)
		h.history[layer] = append([]int{prevA - prevB}, h.history[layer]...)
	}
}

func Process(lines []string) (int, int) {
	next, prev := 0, 0
	for _, line := range lines {
		history := parseHistory(line)
		history.diffTree()
		history.fill()
		next += history.history[0][len(history.history[0])-1]
		prev += history.history[0][0]
	}

	return next, prev
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
