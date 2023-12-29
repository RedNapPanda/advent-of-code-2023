package day5

import (
	aoc "aoc"
	"aoc/aoc_util"
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
	output := Process(example, 1)

	assert.Equal(t, 35, output)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputData(5)
	output := Process(input, 1)

	assert.Equal(t, 313045984, output)
}

func TestPart2Example(t *testing.T) {
	output := Process(example, 2)

	assert.Equal(t, 46, output)
}

// No Goroutines: 195.35s
// Goroutine per seedRange: 46.13s
// Not bruteforce: 509.7µs - could probably shave off more if prioritized  lower layers first and merged ranges.
func TestPart2(t *testing.T) {
	defer aoc_util.Timer("part2")()
	input, _ := aoc.GetInputData(5)
	output := Process(input, 2)

	assert.Equal(t, 20283860, output)
}
