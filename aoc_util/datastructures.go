package aoc_util

import (
	"cmp"
	"fmt"
)

// Tree cmp.Ordered to give each node its own cost that is comparable
type Tree[W cmp.Ordered] struct {
	Data   W
	Left   *Tree[W]
	Right  *Tree[W]
	Parent *Tree[W]
}

type Pair[T, U any] struct {
	First  T
	Second U
}

type Point struct {
	X, Y int
}

type DirPoint struct {
	X, Y, Dir int
}

// ===

/*
FlattenedGrid

Traversing the grid
up/left and right/down inverse
0 = up, 1 = right, 2 = down, 3 = left
Up => n > offset -> n - offset
Down => n < len(data) - offset -> n + offset
Left => n % offset != 0 -> n - 1
Right => (n + 1) % offset -> n + 1

grid[x][y]
*/
type FlattenedGrid[T any] struct {
	Data     []T
	Offset   int
	RowCount int
}

// Get value at grid[x][y]
func (g *FlattenedGrid[T]) Get(x, y int) T {
	return g.Data[x*g.Offset+y]
}

// Set grid[x][y] to data
func (g *FlattenedGrid[T]) Set(x, y int, data T) {
	g.Data[x*g.Offset+y] = data
}

// Index of grid[x][y] in Data
func (g *FlattenedGrid[T]) Index(x, y int) int {
	val := x*g.Offset + y
	if val >= len(g.Data) {
		return -1
	}
	return val
}

// IndexTo convert Index to (x, y) in grid[x][y]
func (g *FlattenedGrid[T]) IndexTo(index int) (int, int) {
	if index < 0 || index >= len(g.Data) {
		return -1, -1
	}
	return index / g.Offset, index % g.Offset
}

// Len - length of the Data
func (g *FlattenedGrid[T]) Len() int {
	return len(g.Data)
}

func (g *FlattenedGrid[T]) Print(format, separator string) {
	fmt.Printf("[%s", separator)
	for n := 0; n < g.Len(); n++ {
		fmt.Printf(format, g.Data[n])
		if (n+1)%g.Offset != 0 {
			fmt.Printf(separator)
		}
		if (n+1)%g.Offset == 0 {
			fmt.Printf("%s]\n", separator)
			if n != g.Len()-1 {
				fmt.Printf("[%s", separator)
			}
		}
	}
}

func (g *FlattenedGrid[T]) Fill(grid [][]T) {
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			g.Data[x*g.Offset+y] = grid[x][y]
		}
	}
}

// NewFlattenedGrid creates a new instance
func NewFlattenedGrid[T any](rowCount, offset int) *FlattenedGrid[T] {
	var g FlattenedGrid[T]
	g.Data = make([]T, rowCount*offset)
	g.RowCount = rowCount
	g.Offset = offset
	return &g
}

// ===

type Graph struct {
}

// ===

type Queue[T any] struct {
}

// ===
