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

func TransposeMatrix[T any](matrix [][]T) [][]T {
	x := len(matrix[0])
	y := len(matrix)
	if x == 0 || y == 0 {
		return nil
	}
	transposed := make([][]T, x)
	for i := 0; i < x; i++ {
		transposed[i] = make([]T, y)
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if len(matrix[j]) == 0 {
				fmt.Printf("Matrix len is zero %d | %+v\n", j, matrix)
			}
			transposed[i][j] = matrix[j][i]
		}
	}
	return transposed
}
