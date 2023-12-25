package aoc_util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestSa(t *testing.T) {
	b := byte(0)
	b |= 1 << 4
	b |= 2 << 4
	b |= 4 << 4
	b |= 8 << 4
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", b>>4)
	fmt.Printf("%t\n", (b>>4)&(1|2|4|8) != 0)
	fmt.Printf("%f\n", math.Log2(8))
}

func TestRepeat(t *testing.T) {
	// No separator
	assert.Equal(t, "1234512345123451234512345", Repeat("12345", "", 5))
	// Separator
	assert.Equal(t, "12345~12345~12345", Repeat("12345", "~", 3))
}

func TestTransposeMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 2},
		{3, 4},
	}
	expected := [][]int{
		{1, 3},
		{2, 4},
	}
	assert.Equal(t, expected, TransposeNewMatrix(matrix))
	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expected = [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}
	assert.Equal(t, expected, TransposeNewMatrix(matrix))
	matrix = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	expected = [][]int{
		{1, 5, 9, 13},
		{2, 6, 10, 14},
		{3, 7, 11, 15},
		{4, 8, 12, 16},
	}
	assert.Equal(t, TransposeNewMatrix(matrix), expected)
	TransposeMatrix(matrix)
	assert.Equal(t, expected, matrix)
}

// TODO: Proper assertions
func TestRotateMatrixCW(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("Original\n")
	PrintMatrix(matrix)
	fmt.Printf("Flipped / RotateNewMatrixCW\n")
	matrix = RotateNewMatrixCW(matrix)
	PrintMatrix(matrix)
	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("Original\n")
	PrintMatrix(matrix)
	fmt.Printf("Flipped / RotateMatrixCW\n")
	RotateMatrixCW(matrix)
	PrintMatrix(matrix)
}

// TODO: Proper assertions
func TestRotateMatrixCCW(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("Original\n")
	PrintMatrix(matrix)
	fmt.Printf("Rotated CCW\n")
	matrix = RotateNewMatrixCCW(matrix)
	PrintMatrix(matrix)
	matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("Original\n")
	PrintMatrix(matrix)
	fmt.Printf("Flipped / RotateMatrixCCW\n")
	RotateMatrixCCW(matrix)
	PrintMatrix(matrix)
}
