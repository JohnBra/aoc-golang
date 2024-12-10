package main

import (
	"fmt"
	"strconv"
	"strings"

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
	visit ds.Set[[2]int],
	r, c, prevVal int,
) (int, int) {

	if r < 0 ||
		c < 0 ||
		r == len(matrix) ||
		c == len(matrix[0]) ||
		visit.Contains([2]int{r, c}) ||
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

	visit.Add([2]int{r, c})

	score, rating := 0, 0
	for _, dir := range dirs {
		sc, ra := dfs(matrix, peaks, head, visit, r+dir[0], c+dir[1], matrix[r][c])
		score += sc
		rating += ra
	}

	visit.Remove([2]int{r, c})

	return score, rating
}

// day10 part 1 and 2 solution
// simple dfs to find each hiking trail
// iterate through matrix, start hiking (dfs) when field == 0
// find every possible path where current matrix val == prevVal + 1
// use map to track if specific trailhead has already hit peak (scores/part one)
// ratings will be every distinct hiking path (ratings/part two)
func hike(matrix [][]int) (int, int) {
	scores, ratings := 0, 0
	rows, cols := len(matrix), len(matrix[0])
	peaks := map[[2]int]ds.Set[[2]int]{}
	visit := ds.NewSet[[2]int]()

	for r := range rows {
		for c := range cols {
			if matrix[r][c] == 0 {
				sc, ra := dfs(matrix, peaks, [2]int{r, c}, visit, r, c, -1)
				scores += sc
				ratings += ra
			}
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

	partOneRes, partTwoRes := hike(matrix)
	fmt.Println("Scores", partOneRes)
	fmt.Println("Ratings", partTwoRes)
}
