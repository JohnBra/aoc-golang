package main

import (
	"container/heap"
	"fmt"
	"math"

	ds "github.com/JohnBra/aoc-2024/internal/datastructures"
	"github.com/JohnBra/aoc-2024/internal/utils"
)

func makeMatrix(memory int, corrupt [][]int, bytes int) [][]rune {
	matrix := make([][]rune, memory+1)
	for r := range memory + 1 {
		col := make([]rune, memory+1)
		for c := range memory + 1 {
			col[c] = '.'
		}
		matrix[r] = col
	}

	for i := 0; i < bytes; i++ {
		matrix[corrupt[i][1]][corrupt[i][0]] = '#'
	}

	return matrix
}

func partOne(matrix [][]rune) int {
	res := math.MaxInt

	// cost, r, c
	h := &ds.MinHeap{}
	heap.Init(h)
	heap.Push(h, []int{0, 0, 0})
	visit := ds.NewSet([2]int{0, 0})

	for h.Len() > 0 {
		cur := heap.Pop(h).([]int)

		if cur[1] == len(matrix)-1 && cur[2] == len(matrix)-1 {
			res = utils.Min(res, cur[0])
			break
		}

		for _, dir := range utils.Axes {
			nr, nc := cur[1]+dir[0], cur[2]+dir[1]
			if !utils.IsOutOfBounds(matrix, nr, nc) && matrix[nr][nc] != '#' && !visit.Contains([2]int{nr, nc}) {
				heap.Push(h, []int{cur[0] + 1, nr, nc})
				visit.Add([2]int{nr, nc})
			}
		}
	}

	return res
}

func main() {
	corrupt, err := utils.GetFileContentsAsIntMatrix(utils.GetPuzzleInputSrc())
	utils.Check(err)

	matrix := makeMatrix(70, corrupt, 1024)
	partOneRes := partOne(matrix)
	fmt.Println("Part one res:", partOneRes)
}
