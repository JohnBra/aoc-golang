package main

import (
	"container/heap"
	"fmt"
	"math"

	ds "github.com/JohnBra/aoc-2024/internal/datastructures"
	"github.com/JohnBra/aoc-2024/internal/utils"
)

func getCost(item []int, orientation int) int {
	diff := utils.Abs(item[3] - orientation)
	if diff > 2 {
		diff = 1
	}
	return item[0] + 1 + diff*1000
}

func walkMaze(matrix [][]rune) int {
	res := math.MaxInt
	visit := ds.NewSet[[3]int]()

	// minheap with items [cost, r, c, orientation]
	h := &ds.IntHeap{}
	heap.Init(h)
	heap.Push(h, []int{0, len(matrix) - 2, 1, 1})

	for h.Len() > 0 {
		item := heap.Pop(h).([]int)
		visit.Add([3]int{item[1], item[2], item[3]})

		if matrix[item[1]][item[2]] == 'E' {
			if item[0] > res {
				break
			}
			res = utils.Min(res, item[0])
		}

		for o, dir := range utils.Axes {
			nr, nc := item[1]+dir[0], item[2]+dir[1]
			opposite := (item[3] + 2) % 4 // opposite direction to current orientation
			if matrix[nr][nc] != '#' && o != opposite && !visit.Contains([3]int{nr, nc, o}) {
				heap.Push(h, []int{getCost(item, o), nr, nc, o})
			}
		}
	}

	return res
}

func parseLine(line string) ([]rune, error) {
	return []rune(line), nil
}

func main() {
	matrix, err := utils.GetSliceOfSlicesFromFile(utils.GetPuzzleInputSrc(), parseLine)
	utils.Check(err)

	partOneRes := walkMaze(matrix)
	fmt.Println("Part one res", partOneRes)
}
