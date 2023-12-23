package aoc_util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeat(t *testing.T) {
	// No separator
	assert.Equal(t, "1234512345123451234512345", Repeat("12345", "", 5))
	// Separator
	assert.Equal(t, "12345~12345~12345", Repeat("12345", "~", 3))
}

func TestTransposeMatrix2x2(t *testing.T) {
	matrix := [][]int{
		{1, 2},
		{3, 4},
	}
	expected := [][]int{
		{1, 3},
		{2, 4},
	}
	assert.Equal(t, TransposeMatrix(matrix), expected)
}

func TestTransposeMatrix3x3(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expected := [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}
	assert.Equal(t, TransposeMatrix(matrix), expected)
}

func TestTransposeMatrix4x4(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	expected := [][]int{
		{1, 5, 9, 13},
		{2, 6, 10, 14},
		{3, 7, 11, 15},
		{4, 8, 12, 16},
	}
	assert.Equal(t, TransposeMatrix(matrix), expected)
}
