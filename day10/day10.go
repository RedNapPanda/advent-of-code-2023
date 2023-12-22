package day10

type coord struct {
	x, y int
}

var pipeCheck = func(tile, prev coord) bool { return tile.y == prev.y }
var dashCheck = func(tile, prev coord) bool { return tile.x == prev.x }
var lCheck = func(tile, prev coord) bool {
	return (tile.x > prev.x && tile.y == prev.y) || (tile.x == prev.x && tile.y < prev.y)
}
var jCheck = func(tile, prev coord) bool {
	return (tile.x > prev.x && tile.y == prev.y) || (tile.x == prev.x && tile.y > prev.y)
}
var sevenCheck = func(tile, prev coord) bool {
	return (tile.x < prev.x && tile.y == prev.y) || (tile.x == prev.x && tile.y > prev.y)
}
var fCheck = func(tile, prev coord) bool {
	return (tile.x < prev.x && tile.y == prev.y) || (tile.x == prev.x && tile.y < prev.y)
}

func Process(lines []string) (int, int) {
	var start coord
	var tiles [][]byte
	cleanTiles := make([][]byte, len(lines))
	for x, line := range lines {
		bytes := []byte(line)
		tiles = append(tiles, bytes)

		for y := 0; y < len(bytes); y++ {
			if bytes[y] == 'S' {
				start = coord{x, y}
			}
		}
		cleanTiles[x] = make([]byte, len(line))
	}
	x1 := max(0, start.x-1)
	y1 := max(0, start.y-1)
	x2 := min(len(tiles), start.x+1)
	y2 := min(len(tiles[start.x]), start.y+1) // data assumption that all rows are the same length

	var checks []coord
	if start.x-1 >= 0 {
		checks = append(checks, coord{x1, start.y})
	}
	if start.y-1 >= 0 {
		checks = append(checks, coord{start.x, y1})
	}
	if start.x+1 <= len(tiles) {
		checks = append(checks, coord{x2, start.y})
	}
	if start.y+1 <= len(tiles[start.x]) {
		checks = append(checks, coord{start.x, y2})
	}

	prev := start
	var next coord
	count := 1
	for _, tile := range checks {
		if p, _, c, ok, _ := nextTile(tiles, start, tile, count); ok {
			next = p
			count = c
			cleanTiles[prev.x][prev.y] = tiles[prev.x][prev.y]
			cleanTiles[next.x][next.y] = tiles[next.x][next.y]
			break
		}
	}

	finished := false
	for !finished {
		prev, next, count, _, finished = nextTile(tiles, prev, next, count)
		cleanTiles[next.x][next.y] = tiles[next.x][next.y]
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

func nextTile(tiles [][]byte, prev coord, target coord, count int, ) (coord, coord, int, bool, bool) {
	next := target
	b := tiles[target.x][target.y]

	x, y := target.x, target.y
	switch {
	case b == '|' && pipeCheck(target, prev):
		if target.x > prev.x {
			x = target.x + 1
		} else {
			x = target.x - 1
		}
	case b == '-' && dashCheck(target, prev):
		if target.y > prev.y {
			y = target.y + 1
		} else {
			y = target.y - 1
		}
	case b == 'L' && lCheck(target, prev):
		if target.x > prev.x && target.y == prev.y {
			y = target.y + 1
		} else {
			x = target.x - 1
		}
	case b == 'J' && jCheck(target, prev):
		if target.x > prev.x && target.y == prev.y {
			y = target.y - 1
		} else {
			x = target.x - 1
		}
	case b == '7' && sevenCheck(target, prev):
		if target.x < prev.x && target.y == prev.y {
			y = target.y - 1
		} else {
			x = target.x + 1
		}
	case b == 'F' && fCheck(target, prev):
		if target.x < prev.x && target.y == prev.y {
			y = target.y + 1
		} else {
			x = target.x + 1
		}
	case b == '.':
		return prev, target, count, false, false
	}
	next = coord{x, y}
	notTarget := next != target
	if notTarget {
		count++
		prev = target
	}

	return prev, next, count, notTarget, b == 'S'
}
