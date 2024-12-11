package main

import (
	"testing"

	"github.com/JohnBra/aoc-2024/internal/utils"
)

// test input
var matrix [][]int = [][]int{
	{8, 9, 0, 1, 0, 1, 2, 3},
	{7, 8, 1, 2, 1, 8, 7, 4},
	{8, 7, 4, 3, 0, 9, 6, 5},
	{9, 6, 5, 4, 9, 8, 7, 4},
	{4, 5, 6, 7, 8, 9, 0, 3},
	{3, 2, 0, 1, 9, 0, 1, 2},
	{0, 1, 3, 2, 9, 8, 0, 1},
	{1, 0, 4, 5, 6, 7, 3, 2},
}

func getInputMatrix(filepath string) [][]int {
	matrix, err := utils.GetSliceOfSlicesFromFile(filepath, parseLine)
	utils.Check(err)
	return matrix
}

func TestHikeRecursive(t *testing.T) {
	scores, ratings := hikeRecursive(matrix)

	if scores != 36 || ratings != 81 {
		t.Errorf("hikeRecursive(testinput) = %d, %d; expected scores == 36 and ratings == 81", scores, ratings)
	}
}

func BenchmarkHikeRecursive(b *testing.B) {
	input := getInputMatrix("./10.in")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = hikeRecursive(input)
	}
}

func TestHikeIterative(t *testing.T) {
	scores, ratings := hikeIterative(matrix)

	if scores != 36 || ratings != 81 {
		t.Errorf("hikeIterative(testinput) = %d, %d; expected scores == 36 and ratings == 81", scores, ratings)
	}
}

func BenchmarkHikeIterative(b *testing.B) {
	input := getInputMatrix("./10.in")

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, _ = hikeIterative(input)
	}
}
