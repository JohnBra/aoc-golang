package main

import (
	"fmt"
	"github.com/JohnBra/aoc-2024/internal/utils"
)

const searchWord string = "XMAS"

var directions = [8][2]int{
	{1, 0},   // down
	{-1, 0},  // up
	{0, 1},   // right
	{0, -1},  // left
	{1, 1},   // down right
	{1, -1},  // down left
	{-1, 1},  // up right
	{-1, -1}, // up left
}

func dfs(matrix [][]string, r int, c int, dir [2]int, i int) int {
	if i == len(searchWord) {
		return 1
	}
	if r < 0 || c < 0 || r == len(matrix) || c == len(matrix[0]) || matrix[r][c] != string(searchWord[i]) {
		return 0
	}

	return dfs(matrix, r+dir[0], c+dir[1], dir, i+1)
}

func partOne(contents [][]string) int {
	rows, cols := len(contents), len(contents[0])
	res := 0

	for r := range rows {
		for c := range cols {
			for _, dir := range directions {
				res += dfs(contents, r, c, dir, 0)
			}
		}
	}

	return res
}

func main() {
	// load contents for both parts as
	// rune matrix
	contents, err := utils.GetFileContentsAsStringMatrix("./input.txt", ``)
	utils.Check(err)

	count := partOne(contents)
	fmt.Println("XMAS count ", count)
}
