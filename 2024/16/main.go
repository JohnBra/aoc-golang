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

func walkMaze(matrix [][]rune) (int, int) {
	resCost := math.MaxInt
	// map with (r, c, o) -> min cost to get there
	lowestCost := ds.Map[[3]int, int]{}
	// record of optimal predecessors to specific location with orientation
	optimalPaths := ds.Map[[3]int, ds.Set[[3]int]]{}
	endStates := ds.Set[[3]int]{}

	// minheap with items [cost, r, c, orientation]
	h := &ds.IntHeap{}
	heap.Init(h)
	heap.Push(h, []int{0, len(matrix) - 2, 1, 1})

	for h.Len() > 0 {
		item := heap.Pop(h).([]int)
		cur := [3]int{item[1], item[2], item[3]}

		if item[0] > lowestCost.Get(cur, math.MaxInt) {
			continue
		}

		if matrix[item[1]][item[2]] == 'E' {
			if item[0] > resCost {
				break
			}
			endStates.Add(cur)
			resCost = item[0]
		}

		for o, dir := range utils.Axes {
			nr, nc := item[1]+dir[0], item[2]+dir[1]
			next := [3]int{nr, nc, o}
			opposite := (item[3] + 2) % 4 // opposite direction of current orientation
			ncost := getCost(item, o)     // cost to get to nr, nc with new orientation
			lcost := lowestCost.Get(next, math.MaxInt)
			if matrix[nr][nc] != '#' && o != opposite && ncost <= lcost {
				if ncost < lcost {
					optimalPaths[next] = ds.NewSet[[3]int]()
					lowestCost[next] = ncost
				}
				optimalPaths[next].Add(cur)
				heap.Push(h, []int{ncost, nr, nc, o})
			}
		}
	}

	states, visit := ds.NewDeque([][3]int{}), ds.NewSet[[3]int]()

	for _, s := range endStates.Members() {
		states.PushBack(s)
		visit.Add(s)
	}

	for states.Len() > 0 {
		key, _ := states.PopFront()
		for prev := range optimalPaths.Get(key, ds.NewSet[[3]int]()) {
			if !visit.Contains(prev) {
				visit.Add(prev)
				states.PushBack(prev)
			}
		}
	}

	seats := ds.NewSet[[2]int]()
	for _, v := range visit.Members() {
		seats.Add([2]int{v[0], v[1]})
	}

	return resCost, len(seats)
}

func parseLine(line string) ([]rune, error) {
	return []rune(line), nil
}

func main() {
	matrix, err := utils.GetSliceOfSlicesFromFile(utils.GetPuzzleInputSrc(), parseLine)
	utils.Check(err)

	partOneRes, partTwoRes := walkMaze(matrix)
	fmt.Println("Part one res", partOneRes)
	fmt.Println("Part two res", partTwoRes)
}
