package day15

import (
	"fmt"
	"strings"
)

type cameraLens struct {
	label    []byte
	focalLen int
}

func Process(input string) int {
	input = strings.ReplaceAll(input, "\n", "")
	split := strings.Split(input, ",")
	val := 0
	for _, line := range split {
		bytes := []byte(line)
		val += hash(bytes)
	}
	return val
}

func Process2(input string) int {
	input = strings.ReplaceAll(input, "\n", "")
	split := strings.Split(input, ",")
	val := 0

	boxes := make([][]cameraLens, 256)
	for _, sequence := range split {
		bytes := []byte(sequence)

		focalLen := 0
		remove := false
		switch {
		case bytes[len(bytes)-1] == '-':
			remove = true
		case bytes[len(bytes)-2] == '=':
			focalLen = int(bytes[len(bytes)-1]) - 48
		}
		var label []byte
		if remove {
			slice := bytes[:len(bytes)-1]
			label = make([]byte, len(bytes)-1)
			copy(label, slice)
		} else {
			slice := bytes[:len(bytes)-2]
			label = make([]byte, len(bytes)-2)
			copy(label, slice)
		}

		box := hash(label)
		if remove && len(boxes[box]) > 0 {
			for i, l := range boxes[box] {
				if len(l.label) != len(label) {
					continue
				}
				matches := true
				for n := 0; n < len(label); n++ {
					if l.label[n] != label[n] {
						matches = false
						break
					}
				}
				if matches {
					if len(boxes[box]) == 1 {
						boxes[box] = nil
					} else {
						boxes[box] = append(boxes[box][:i], boxes[box][i+1:]...)
					}
					break
				}
			}
		} else if !remove {
			l := cameraLens{label, focalLen}
			if len(boxes[box]) == 0 {
				boxes[box] = []cameraLens{l}
			} else {
				swapped := false
				for i, lensInBox := range boxes[box] {
					if len(lensInBox.label) != len(label) {
						continue
					}
					matches := true
					for n := 0; n < len(label); n++ {
						if lensInBox.label[n] != label[n] {
							matches = false
							break
						}
					}
					if matches {
						boxes[box][i] = l
						swapped = true
						break
					}
				}
				if !swapped {
					boxes[box] = append(boxes[box], l)
				}
			}
		}
	}

	for boxIndex, box := range boxes {
		for i, l := range box {
			val += (boxIndex + 1) * (i + 1) * l.focalLen
		}
	}
	return val
}

func RecursiveProcess(input string) int {
	input = strings.ReplaceAll(input, "\n", "")
	split := strings.Split(input, ",")
	val := 0
	for _, line := range split {
		val += recursiveHash([]byte(line), 0, 0)
	}
	return val
}

/*
HASH steps for each c
	v += ascii code (c)
	v *= 17
	v %= 256

func([]) => nothing
func([x : ]) => fn(x), func([
*/

func hash(bytes []byte) int {
	if len(bytes) == 0 {
		return 0
	}

	val := 0
	for i := 0; i < len(bytes); i++ {
		val = calculate(bytes[i], val)
	}
	return val
}

var recurseCache = make(map[string]int)

func recursiveHash(bytes []byte, x, v int) int {
	if x == len(bytes) {
		return v
	}
	key := fmt.Sprintf("%c,%d", int(bytes[x]), v)
	if h, ok := recurseCache[key]; ok {
		return recursiveHash(bytes, x+1, h)
	}
	return recursiveHash(bytes, x+1, calculate(bytes[x], v))
}

func calculate(c byte, val int) int {
	val += int(c)
	val *= 17
	val %= 256
	return val
}
