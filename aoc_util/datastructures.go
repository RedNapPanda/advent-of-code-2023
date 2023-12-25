package aoc_util

type Tree[T any] struct {
	Data   T
	Left   *Tree[T]
	Right  *Tree[T]
	Parent *Tree[T]
}

type Coordinate struct {
	X, Y int
}
type CardinalNode struct {
	X, Y, Cardinal int
}
