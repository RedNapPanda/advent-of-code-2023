package day18

import (
	"strconv"
	"strings"
)

const (
	right = iota
	down
	left
	up
)

type ColoredPoint struct {
	x, y int
	c    string
}

/*
Process

This problem was pretty obvious it needed either shoelace or pick's theorem
Only knew those formula from recent other related research from earlier days
Took longer to implement and debug what the difference between the area + perimeter and the result
*/
func Process(lines []string, part int) int {

	x, y := 0, 0
	count := 0
	points := make([]ColoredPoint, len(lines))
	for i, l := range lines {
		points[i] = ColoredPoint{x: x, y: y, c: ""}
		dir, steps, _ := parse(l, part)

		switch dir {
		case right:
			y += steps
		case down:
			x += steps
		case left:
			y -= steps
		case up:
			x -= steps
		}
		count += steps
	}

	// result 42, off by 20 from example answer
	// perimeter is 38, p / 2 = 19, which is 1 off..
	// This is from pick's theorem -> i + 0.5p - 1
	v := shoelace(points) + count/2 + 1

	return v
}

func Process2(lines []string, part int) int {

	x, y := 0, 0
	count := 0
	points := make([]ColoredPoint, len(lines))
	for i, l := range lines {
		points[i] = ColoredPoint{x: x, y: y, c: ""}
		dir, steps, _ := parse(l, part)

		switch dir {
		case right:
			y += steps
		case down:
			x += steps
		case left:
			y -= steps
		case up:
			x -= steps
		}
		count += steps
	}

	// result 42, off by 20 from example answer
	// perimeter is 38, p / 2 = 19, which is 1 off..
	// This is from pick's theorem -> i + 0.5p - 1
	// the trench is inside the polygon and therefore the polygon is actually half a step larger
	// difference between the 2 polygons is perimeter/2 + 1
	return shoelace(points) + count/2 + 1
}

func parse(l string, part int) (int, int, string) {
	s := strings.Split(l, " ")
	var d int
	switch s[0] {
	case "R":
		d = right
	case "D":
		d = down
	case "L":
		d = left
	case "U":
		d = up
	}
	i, _ := strconv.Atoi(s[1])

	// p2 is hex
	// 1st 5 digits is hex steps
	// last hex is direction
	hex := s[2][2 : len(s[2])-1]
	if part == 2 {
		d = int(hex[len(hex)-1] - '0')
		v, _ := strconv.ParseInt(hex[:len(hex)-1], 16, 64)
		i = int(v)
	}

	return d, i, s[1]
}

func shoelace(ps []ColoredPoint) int {
	a := 0
	for i := 0; i < len(ps); i++ {
		j := (i + 1) % len(ps)
		a += ps[i].x*ps[j].y - ps[j].x*ps[i].y
	}
	if a < 0 {
		a *= -1
	}
	return a / 2
}
