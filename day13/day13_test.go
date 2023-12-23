package day13

import (
	aoc "aoc"
	"github.com/stretchr/testify/assert"
	"testing"
)

var example = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

/*
Goal to find the reflection points that falls within the body of the data points

Seems this dataset is bound to a single reflection?  At least from example input
Doesn't seem to mention multiple reflections..

1 in this 1st pattern also has nothing to reflect to
    123456789                   1 #...##..# 1 Since 8 is missing, there's no need to check this
        ><                      2 #....#..#v2 2 matches 7
    #.##..##.                   3 ..##..###v3 3 matches 6
    ..#.##.#.                   4v#####.##.v4 5 matches 5
    ##......#                   5^#####.##.^5
    ##......#                   6 ..##..###^6
    ..#.##.#.                   7 #....#..#^7
    ..##..##.
    #.#.##.#.
    >>>>><<<<
    123456789

When parsing, convert '.' to 1 and '#' to 0 (or vice versa)
For rows (columns is just rows transposed, but same checks)
for i to grid X dim
    Reverse slice of first i lines as upper mirror (to what it would look like mirrored on the lower half)
    slice of remaining lines after upper
    row count bounded to smallest set of rows
    Check to see if both lists match, if so, i + 1 is the mirror line
*/

func TestPart1Example(t *testing.T) {
	output := Process(example)
	assert.Equal(t, 405, output)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputDataString(13)
	output := Process(input)

	assert.Equal(t, 36015, output)
}
