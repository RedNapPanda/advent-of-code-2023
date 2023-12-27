package day10

import "aoc/aoc_util"

var pipeCheck = func(tile, prev aoc_util.Point) bool { return tile.Y == prev.Y }
var dashCheck = func(tile, prev aoc_util.Point) bool { return tile.X == prev.X }
var lCheck = func(tile, prev aoc_util.Point) bool {
	return (tile.X > prev.X && tile.Y == prev.Y) || (tile.X == prev.X && tile.Y < prev.Y)
}
var jCheck = func(tile, prev aoc_util.Point) bool {
	return (tile.X > prev.X && tile.Y == prev.Y) || (tile.X == prev.X && tile.Y > prev.Y)
}
var sevenCheck = func(tile, prev aoc_util.Point) bool {
	return (tile.X < prev.X && tile.Y == prev.Y) || (tile.X == prev.X && tile.Y > prev.Y)
}
var fCheck = func(tile, prev aoc_util.Point) bool {
	return (tile.X < prev.X && tile.Y == prev.Y) || (tile.X == prev.X && tile.Y < prev.Y)
}

func Process(lines []string) (int, int) {
	var start aoc_util.Point
	var tiles [][]byte
	cleanTiles := make([][]byte, len(lines))
	for x, line := range lines {
		bytes := []byte(line)
		tiles = append(tiles, bytes)

		for y := 0; y < len(bytes); y++ {
			if bytes[y] == 'S' {
				start = aoc_util.Point{X: x, Y: y}
			}
		}
		cleanTiles[x] = make([]byte, len(line))
	}
	x1 := max(0, start.X-1)
	y1 := max(0, start.Y-1)
	x2 := min(len(tiles), start.X+1)
	y2 := min(len(tiles[start.X]), start.Y+1) // data assumption that all rows are the same length

	var checks []aoc_util.Point
	if start.X-1 >= 0 {
		checks = append(checks, aoc_util.Point{X: x1, Y: start.Y})
	}
	if start.Y-1 >= 0 {
		checks = append(checks, aoc_util.Point{X: start.X, Y: y1})
	}
	if start.X+1 <= len(tiles) {
		checks = append(checks, aoc_util.Point{X: x2, Y: start.Y})
	}
	if start.Y+1 <= len(tiles[start.X]) {
		checks = append(checks, aoc_util.Point{X: start.X, Y: y2})
	}

	prev := start
	var next aoc_util.Point
	count := 1
	for _, tile := range checks {
		if p, _, c, ok, _ := nextTile(tiles, start, tile, count); ok {
			next = p
			count = c
			cleanTiles[prev.X][prev.Y] = tiles[prev.X][prev.Y]
			cleanTiles[next.X][next.Y] = tiles[next.X][next.Y]
			break
		}
	}

	finished := false
	for !finished {
		prev, next, count, _, finished = nextTile(tiles, prev, next, count)
		cleanTiles[next.X][next.Y] = tiles[next.X][next.Y]
	}

	area := 0
	for _, row := range cleanTiles {
		inArea := false
		for i, c := range row {
			if c == '|' || c == '7' || c == 'F' {
				inArea = !inArea
			}
			if i != 0 && i+1 != len(row) && inArea && c == 0 {
				area++
			}
		}
	}

	return (count - 1) / 2, area
}

func nextTile(tiles [][]byte, prev aoc_util.Point, target aoc_util.Point, count int) (aoc_util.Point, aoc_util.Point, int, bool, bool) {
	next := target
	b := tiles[target.X][target.Y]

	x, y := target.X, target.Y
	switch {
	case b == '|' && pipeCheck(target, prev):
		if target.X > prev.X {
			x = target.X + 1
		} else {
			x = target.X - 1
		}
	case b == '-' && dashCheck(target, prev):
		if target.Y > prev.Y {
			y = target.Y + 1
		} else {
			y = target.Y - 1
		}
	case b == 'L' && lCheck(target, prev):
		if target.X > prev.X && target.Y == prev.Y {
			y = target.Y + 1
		} else {
			x = target.X - 1
		}
	case b == 'J' && jCheck(target, prev):
		if target.X > prev.X && target.Y == prev.Y {
			y = target.Y - 1
		} else {
			x = target.X - 1
		}
	case b == '7' && sevenCheck(target, prev):
		if target.X < prev.X && target.Y == prev.Y {
			y = target.Y - 1
		} else {
			x = target.X + 1
		}
	case b == 'F' && fCheck(target, prev):
		if target.X < prev.X && target.Y == prev.Y {
			y = target.Y + 1
		} else {
			x = target.X + 1
		}
	case b == '.':
		return prev, target, count, false, false
	}
	next = aoc_util.Point{X: x, Y: y}
	notTarget := next != target
	if notTarget {
		count++
		prev = target
	}

	return prev, next, count, notTarget, b == 'S'
}
