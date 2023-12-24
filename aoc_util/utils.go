package aoc_util

import (
	"fmt"
	"time"
)

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

/*
TransposeMatrix

	transposes a matrix
	does not do any validations
	assumes all rows are equal column length
	TODO: Harden this to fail if matrix is not balanced?
*/
func TransposeMatrix[T any](matrix [][]T) [][]T {
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
RotateMatrixCW

	Previously was: flips a matrix via transpose, then reversing each row
	does not do any validations
	assumes all rows are equal column length
	TODO: Harden this to fail if matrix is not balanced?
*/
func RotateMatrixCW[T any](matrix [][]T) [][]T {
	x := len(matrix)    // row
	y := len(matrix[0]) // col
	newMatrix := make([][]T, x)
	for j := 0; j < y; j++ {
		col := make([]T, y)
		for i := x - 1; i >= 0; i-- {
			// insert from bottom to top
			col[x-i-1] = matrix[i][j]
		}
		newMatrix[j] = col
	}
	return newMatrix
}

/*
RotateMatrixCCW

	does not do any validations
	assumes all rows are equal column length
	TODO: Harden this to fail if matrix is not balanced?
*/
func RotateMatrixCCW[T any](matrix [][]T) [][]T {
	x := len(matrix[0])
	y := len(matrix)
	newMatrix := make([][]T, x)
	for i := 0; i < x; i++ {
		newMatrix[i] = make([]T, y)
	}
	for j := 0; j < y; j++ {
		for i := x - 1; i >= 0; i-- {
			newMatrix[x-i-1][j] = matrix[j][i]
		}
	}
	return newMatrix
}

func PrintMatrix[T any](matrix [][]T) {
	for i := 0; i < len(matrix); i++ {
		fmt.Printf("[")
		for j := 0; j < len(matrix[i]); j++ {
			if b, ok := any(matrix[i][j]).(byte); ok {
				fmt.Printf("%c", b)
			} else {
				fmt.Printf("%+v", matrix[i][j])
			}
			if j < len(matrix[i])-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("]\n")
	}
}
