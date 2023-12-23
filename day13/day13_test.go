package day13

import (
	aoc "aoc"
	"github.com/stretchr/testify/assert"
	"testing"
)

var example = []string{
	"#.##..##.",
	"..#.##.#.",
	"##......#",
	"##......#",
	"..#.##.#.",
	"..##..##.",
	"#.#.##.#.",
	"",
	"#...##..#",
	"#....#..#",
	"..##..###",
	"#####.##.",
	"#####.##.",
	"..##..###",
	"#....#..#",
}

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
     98765432
*/

// invertIndex
// s := "12345678"
// fmt.Println(s)
// fmt.Printf("%c\n", s[0])
// fmt.Printf("%c\n", s[len(s)-0-1])

func TestPart1Example(t *testing.T) {
	output := Process(example)
	assert.Equal(t, 405, output)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputData(13)
	output := Process(input)

	assert.Equal(t, 6949, output)
}
