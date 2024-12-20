package main

import (
	"container/heap"
	"fmt"

	ds "github.com/JohnBra/aoc-2024/internal/datastructures"
	"github.com/JohnBra/aoc-2024/internal/utils"
)

func race(matrix [][]rune, r, c int) int {
	res := 0

	h := &ds.IntHeap{}
	heap.Init(h)
	heap.Push(h, []int{0, r, c})
	visit := ds.NewSet([2]int{r, c})

	for h.Len() > 0 {
		item := heap.Pop(h).([]int)

		if matrix[item[1]][item[2]] == 'E' {
			res = item[0]
			break
		}

		for _, dir := range utils.Axes {
			nr, nc := item[1]+dir[0], item[2]+dir[1]
			next := [2]int{nr, nc}

			if matrix[nr][nc] != '#' && !visit.Contains(next) {

				visit.Add(next)
				heap.Push(h, []int{item[0] + 1, nr, nc})
			}
		}
	}
	return res
}

func partOne(matrix [][]rune) int {
	// res, start r/c, race time without cheat
	res, sr, sc, baseline := 0, 0, 0, 0
	// all walls that are not the outside border
	walls := ds.NewSet[[2]int]()

	// find start/end and all inner walls
	for r := range len(matrix) {
		for c := range len(matrix[0]) {
			if matrix[r][c] == 'S' {
				sr, sc = r, c
			}

			if r > 0 && r < len(matrix)-1 && c > 0 && c < len(matrix[0])-1 && matrix[r][c] == '#' {
				walls.Add([2]int{r, c})
			}
		}
	}

	baseline = race(matrix, sr, sc)

	for _, w := range walls.Members() {
		matrix[w[0]][w[1]] = '.'

		time := race(matrix, sr, sc)
		if baseline-time >= 100 {
			res += 1
		}

		matrix[w[0]][w[1]] = '#'
	}

	return res
}

func parseLine(line string) ([]rune, error) {
	return []rune(line), nil
}

func main() {
	matrix, err := utils.GetSliceOfSlicesFromFile(utils.GetPuzzleInputSrc(), parseLine)
	utils.Check(err)

	partOneRes := partOne(matrix)
	fmt.Println("Part one res", partOneRes)
}
