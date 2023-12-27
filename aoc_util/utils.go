package aoc_util

import (
	"fmt"
	"slices"
	"time"
)

func Zero[T any]() T {
	return *new(T)
}

func IsZero[T comparable](t T) bool {
	return t == Zero[T]()
}

// Timer credits: https://stackoverflow.com/questions/45766572/is-there-an-efficient-way-to-calculate-execution-time-in-golang
//
// Learned more about defer and how you could use/abuse it like this
//
// timer returns a function that prints the name argument and
// the elapsed time between the call to timer and the call to
// the returned function. The returned function is intended to
// be used in a defer statement:
//
// defer timer("sum")()
func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func Repeat(orig, sep string, times int) string {
	result := orig
	for i := 1; i < times; i++ {
		result += sep + orig
	}
	return result
}

func JoinSlices[T any](separator T, slices ...[]T) []T {
	var newSlice []T
	for i, slice := range slices {
		newSlice = append(newSlice, slice...)
		if i < len(slices)-1 {
			newSlice = append(newSlice, separator)
		}
	}
	return newSlice
}

// SlicesEqual checks that the slices are the same length and the elements are in the same order
func SlicesEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func FlipMatrix[T any](matrix [][]T, horizontal bool) {
	if !horizontal {
		slices.Reverse(matrix)
	} else {
		for i := range matrix {
			slices.Reverse(matrix[i])
		}
	}
}

func RotateMatrixCW[T any](matrix [][]T) {
	m := len(matrix[0]) - 1 // col
	// split matrix into 4 squares
	for y := 0; y < (m+1)/2; y++ {
		for x := y; x < m-y; x++ {
			// shift bottom left up
			// shift bottom right left
			// shift top right down
			// shift top left right
			matrix[y][x], matrix[m-x][y], matrix[m-y][m-x], matrix[x][m-y] = matrix[m-x][y], matrix[m-y][m-x], matrix[x][m-y], matrix[y][x]
		}
	}
}

func RotateMatrixCCW[T any](matrix [][]T) {
	m := len(matrix[0]) - 1 // col
	// split matrix into 4 squares
	for y := 0; y < (m+1)/2; y++ {
		for x := y; x < m-y; x++ {
			temp := matrix[x][y]
			// shift top right left
			matrix[x][y] = matrix[y][m-x]
			// shift bottom right up
			matrix[y][m-x] = matrix[m-x][m-y]
			// shift bottom left right
			matrix[m-x][m-y] = matrix[m-y][x]
			// shift top left down
			matrix[m-y][x] = temp
		}
	}
}

func TransposeMatrix[T any](matrix [][]T) {
	x := len(matrix)
	for i := 0; i < x; i++ {
		for j := i + 1; j < x; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

/*
TransposeNewMatrix

	transposes a matrix
	does not do any validations
	assumes all rows are equal column length
	TODO: Harden this to fail if matrix is not balanced?
*/
func TransposeNewMatrix[T any](matrix [][]T) [][]T {
	x := len(matrix[0])
	y := len(matrix)
	newMatrix := make([][]T, x)
	for i := 0; i < x; i++ {
		newMatrix[i] = make([]T, y)
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			newMatrix[i][j] = matrix[j][i]
		}
	}
	return newMatrix
}

/*
RotateNewMatrixCW

	Previously was: flips a matrix via transpose, then reversing each row
	does not do any validations
	assumes all rows are equal column length
	TODO: Harden this to fail if matrix is not balanced?
*/
func RotateNewMatrixCW[T any](matrix [][]T) [][]T {
	x := len(matrix)    // row
	y := len(matrix[0]) // col
	newMatrix := make([][]T, y)
	for j := 0; j < y; j++ {
		col := make([]T, x)
		for i := x - 1; i >= 0; i-- {
			// insert from bottom to top
			col[x-i-1] = matrix[i][j]
		}
		newMatrix[j] = col
	}
	return newMatrix
}

/*
RotateNewMatrixCCW

	does not do any validations
	assumes all rows are equal column length
	TODO: Harden this to fail if matrix is not balanced?
*/
func RotateNewMatrixCCW[T any](matrix [][]T) [][]T {
	j := len(matrix)
	i := len(matrix[0])
	newMatrix := make([][]T, i)
	for x := 0; x < i; x++ {
		newMatrix[x] = make([]T, j)
	}
	for x := 0; x < j; x++ {
		for y := i - 1; y >= 0; y-- {
			newMatrix[i-y-1][x] = matrix[x][y]
		}
	}
	return newMatrix
}

func CopyMatrix[T any](matrix [][]T) [][]T {
	j := len(matrix)
	i := len(matrix[0])
	newMatrix := make([][]T, j)
	for x := 0; x < j; x++ {
		newMatrix[x] = make([]T, i)
		for y := 0; y < i; y++ {
			newMatrix[x][y] = matrix[x][y]
		}
	}
	return newMatrix
}

func PrintMatrix[T any](matrix [][]T) {
	for x := 0; x < len(matrix); x++ {
		fmt.Printf("[")
		for y := 0; y < len(matrix[x]); y++ {
			if b, ok := any(matrix[x][y]).(byte); ok {
				fmt.Printf("%c", b)
			} else {
				fmt.Printf("%+v", matrix[x][y])
			}
			if y < len(matrix[x])-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("]\n")
	}
}
