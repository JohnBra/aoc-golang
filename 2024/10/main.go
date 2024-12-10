package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	ds "github.com/JohnBra/aoc-2024/internal/datastructures"
	"github.com/JohnBra/aoc-2024/internal/utils"
)

var dirs = [4][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func dfs(
	matrix [][]int,
	peaks map[[2]int]ds.Set[[2]int],
	head [2]int,
	r, c, prevVal int,
) (int, int) {

	if r < 0 ||
		c < 0 ||
		r == len(matrix) ||
		c == len(matrix[0]) ||
		matrix[r][c] != prevVal+1 {
		return 0, 0
	}

	if matrix[r][c] == 9 {
		score, rating := 0, 1
		key := [2]int{r, c}
		_, ok := peaks[key]
		if !ok {
			peaks[key] = ds.NewSet[[2]int]()
		}

		if !peaks[key].Contains(head) {
			peaks[key].Add(head)
			score = 1
		}
		return score, rating
	}

	score, rating := 0, 0
	for _, dir := range dirs {
		sc, ra := dfs(matrix, peaks, head, r+dir[0], c+dir[1], matrix[r][c])
		score += sc
		rating += ra
	}

	return score, rating
}

// day10 part 1 and 2 solution
// simple dfs to find each hiking trail
// iterate through matrix, start hiking (dfs) when field == 0
// find every possible path where current matrix val == prevVal + 1
// use map to track if specific trailhead has already hit peak (scores/part one)
// ratings will be every distinct hiking path (ratings/part two)
func hikeRecursive(matrix [][]int) (int, int) {
	defer utils.TimeTrack(time.Now(), "hike recursive")
	scores, ratings := 0, 0
	rows, cols := len(matrix), len(matrix[0])
	peaks := map[[2]int]ds.Set[[2]int]{}

	for r := range rows {
		for c := range cols {
			if matrix[r][c] == 0 {
				sc, ra := dfs(matrix, peaks, [2]int{r, c}, r, c, -1)
				scores += sc
				ratings += ra
			}
		}
	}
	return scores, ratings
}

// Iterative version to the solution for the fun of it
func hikeIterative(matrix [][]int) (int, int) {
	defer utils.TimeTrack(time.Now(), "hike iterative")
	scores, ratings := 0, 0
	rows, cols := len(matrix), len(matrix[0])
	// stack item: [r, c]
	stack := ds.NewStack[[2]int]()
	peaks := map[[2]int]ds.Set[[2]int]{}

	for r := range rows {
		for c := range cols {
			if matrix[r][c] == 0 {
				stack.Push([2]int{r, c})
			}
		}
	}

	_, ok := stack.Top()
	if !ok {
		fmt.Println("no trail heads found")
		return 0, 0
	}

	var head [2]int

	for !stack.IsEmpty() {
		item, _ := stack.Pop()
		r, c := item[0], item[1]

		if matrix[r][c] == 9 { // add score and rating if peak
			key := [2]int{r, c}
			_, ok := peaks[key]
			if !ok {
				peaks[key] = ds.NewSet[[2]int]()
			}

			if !peaks[key].Contains(head) {
				peaks[key].Add(head)
				scores += 1
			}

			ratings += 1
			continue
		}

		if matrix[r][c] == 0 { // set current trail head
			head[0], head[1] = r, c
		}

		// only add valid items to stack (i.e. not out of bounds and in order of hike)
		for _, dir := range dirs {
			nextR, nextC := r+dir[0], c+dir[1]
			if nextR < 0 ||
				nextC < 0 ||
				nextR == len(matrix) ||
				nextC == len(matrix[0]) ||
				matrix[nextR][nextC] != matrix[r][c]+1 {
				continue
			}
			stack.Push([2]int{nextR, nextC})
		}
	}

	return scores, ratings
}

func parseLine(line string) ([]int, error) {
	res := make([]int, len(line))
	strings := strings.Split(line, ``)

	for i, s := range strings {
		num, err := strconv.Atoi(s)
		if err != nil {
			return res, err
		}
		res[i] = num
	}
	return res, nil
}

func main() {
	src := utils.GetPuzzleInputSrc()
	matrix, err := utils.GetSliceOfSlicesFromFile(src, parseLine)
	utils.Check(err)

	partOneResRec, partTwoResRec := hikeRecursive(matrix)
	fmt.Println("Scores (recursive)", partOneResRec)
	fmt.Println("Ratings (recursive)", partTwoResRec)

	fmt.Println("")

	partOneResIter, partTwoResIter := hikeIterative(matrix)
	fmt.Println("Scores (iterative)", partOneResIter)
	fmt.Println("Ratings (iterative)", partTwoResIter)
}
