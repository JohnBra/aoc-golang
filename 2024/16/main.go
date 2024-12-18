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
	r, c := len(matrix)-2, 1
	res := math.MaxInt
	visit := ds.NewSet[[3]int]()

	// cost, r, c, orientation
	h := &ds.MinHeap{}
	heap.Init(h)
	heap.Push(h, []int{0, r, c, 1})

	for h.Len() > 0 {
		item := heap.Pop(h).([]int)

		if matrix[item[1]][item[2]] == 'E' {
			fmt.Println("found path with cost", item[0])
			if item[0] > res {
				break
			}
			res = utils.Min(res, item[0])
		}

		for o, dir := range utils.Axes {
			nextR, nextC := item[1]+dir[0], item[2]+dir[1]
			if matrix[nextR][nextC] != '#' && !visit.Contains([3]int{nextR, nextC, o}) {
				visit.Add([3]int{nextR, nextC, o})
				heap.Push(h, []int{getCost(item, o), nextR, nextC, o})
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
