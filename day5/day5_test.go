package day5

import (
	aoc "aoc"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var example = strings.Split(`seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`, "\n")

func TestPart1Example(t *testing.T) {
	output := Part1(example)

	assert.Equal(t, 35, output)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputData(5)
	output := Part1(input)

	assert.Equal(t, 313045984, output)
}

func TestPart2Example(t *testing.T) {
	output := Part2(example)

	assert.Equal(t, 46, output)
}

// No Goroutines: 195.35s
// Goroutine per seedRange: 46.13s
func TestPart2(t *testing.T) {
	input, _ := aoc.GetInputData(5)
	output := Part2(input)

	assert.Equal(t, 20283860, output)
}

func TestIntersect(t *testing.T) {
	m := mapRange{5, 20, 1}
	s := mapRange{100, 106, 2}
	b, i, rem := m.intersect(s)
	fmt.Printf("%t | %+v | %+v\n", b, i, rem)
}
