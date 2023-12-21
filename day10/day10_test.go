package day10

import (
	aoc "aoc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1Example1(t *testing.T) {
	example := []string{
		".....",
		".S-7.",
		".|.|.",
		".L-J.",
		".....",
	}
	output, _ := Process(example)
	assert.Equal(t, 4, output)
}

func TestPart1Example2(t *testing.T) {
	example := []string{
		"-L|F7",
		"7S-7|",
		"L|7||",
		"-L-J|",
		"L|-JF",
	}
	output, _ := Process(example)
	assert.Equal(t, 4, output)
}

func TestPart1Example3(t *testing.T) {
	example := []string{
		"..F7.",
		".FJ|.",
		"SJ.L7",
		"|F--J",
		"LJ...",
	}
	output, _ := Process(example)
	assert.Equal(t, 8, output)
}

func TestPart1Example4(t *testing.T) {
	example := []string{
		"7-F7-",
		".FJ|7",
		"SJLL7",
		"|F--J",
		"LJ.LJ",
	}
	output, _ := Process(example)
	assert.Equal(t, 8, output)
}

func TestPart1(t *testing.T) {
	input, _ := aoc.GetInputData(10)
	output, _ := Process(input)

	assert.Equal(t, 7012, output)
}

func TestPart2Example1(t *testing.T) {
	example := []string{
		"...........",
		".S-------7.",
		".|F-----7|.",
		".||.....||.",
		".||.....||.",
		".|L-7.F-J|.",
		".|..|.|..|.",
		".L--J.L--J.",
		"...........",
	}
	_, output := Process(example)
	assert.Equal(t, 4, output)
}

func TestPart2Example2(t *testing.T) {
	example := []string{
		"..........",
		".S------7.",
		".|F----7|.",
		".||....||.",
		".||....||.",
		".|L-7F-J|.",
		".|..||..|.",
		".L--JL--J.",
		"..........",
	}
	_, output := Process(example)
	assert.Equal(t, 4, output)
}

func TestPart2Example3(t *testing.T) {
	example := []string{
		".F----7F7F7F7F-7....",
		".|F--7||||||||FJ....",
		".||.FJ||||||||L7....",
		"FJL7L7LJLJ||LJ.L-7..",
		"L--J.L7...LJS7F-7L7.",
		"....F-J..F7FJ|L7L7L7",
		"....L7.F7||L7|.L7L7|",
		".....|FJLJ|FJ|F7|.LJ",
		"....FJL-7.||.||||...",
		"....L---J.LJ.LJLJ...",
	}
	_, output := Process(example)
	assert.Equal(t, 8, output)
}

func TestPart2Example4(t *testing.T) {
	example := []string{
		"FF7FSF7F7F7F7F7F---7",
		"L|LJ||||||||||||F--J",
		"FL-7LJLJ||||||LJL-77",
		"F--JF--7||LJLJ7F7FJ-",
		"L---JF-JLJ.||-FJLJJ7",
		"|F|F-JF---7F7-L7L|7|",
		"|FFJF7L7F-JF7|JL---7",
		"7-L-JL7||F7|L7F-7F7|",
		"L.L7LFJ|||||FJL7||LJ",
		"L7JLJL-JLJLJL--JLJ.L",
	}
	_, output := Process(example)
	assert.Equal(t, 10, output)
}

func TestPart2(t *testing.T) {
	input, _ := aoc.GetInputData(10)
	_, output := Process(input)

	assert.Equal(t, 395, output)
}
