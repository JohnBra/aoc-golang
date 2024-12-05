package main

import (
	"fmt"
	"time"

	"github.com/JohnBra/aoc-2024/internal/utils"
)

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

/*
{1, 1},	// [0] ↘ => 2 down | 2 right from origin and then [2] ↗ | ↙️ [1]
{1, -1},	// [1] ↙️ => 2 down | 2 left from origin and then [3] ↖ | ↘ [0]
{-1, 1}, // [2] ↗ => 2 up | 2 right from origin and then [0] ↘ | ↖ [3]
{-1, -1}, // [3] ↖ => 2 up | 2 left from origin and then [1] ↙️ | ↗ [2]
*/
var diagonals = [4][2]int{
	{1, 1},   // ↘
	{1, -1},  // ↙️
	{-1, 1},  // ↗
	{-1, -1}, // ↖
}

type DirectionModifier struct {
	Row, Col, Dir int
}

/*
If "MAS" found do the following for diagonal direction

Note: map key and last value of tuple are index of direction
in diagonals array

0: (r+2, c+0, [2]) + (r+0, c+2, [1])
1: (r+2, c+0, [3]) + (r+0, c-2, [0])
2: (r-2, c+0, [0]) + (r+0, c+2, [3])
3: (r-2, c+0, [1]) + (r+0, c-2, [2])
*/
var crossCheck = map[int][2]DirectionModifier{
	0: {{2, 0, 2}, {0, 2, 1}},
	1: {{2, 0, 3}, {0, -2, 0}},
	2: {{-2, 0, 0}, {0, 2, 3}},
	3: {{-2, 0, 1}, {0, -2, 2}},
}

func dfs(matrix [][]string, r int, c int, dir [2]int, i int, word string) int {
	if i == len(word) {
		return 1
	}
	if r < 0 ||
		c < 0 ||
		r == len(matrix) ||
		c == len(matrix[0]) ||
		matrix[r][c] != string(word[i]) {
		return 0
	}

	return dfs(matrix, r+dir[0], c+dir[1], dir, i+1, word)
}

// day4 part 1
// iterate through each field in xmas matrix
// dfs recursively in each direction
// dfs will return 1 if xmas string was found for field
func partOne(contents [][]string) int {
	start := time.Now()

	defer utils.TimeTrack(start, "partOne")
	rows, cols := len(contents), len(contents[0])
	res := 0

	for r := range rows {
		for c := range cols {
			for _, dir := range directions {
				res += dfs(contents, r, c, dir, 0, "XMAS")
			}
		}
	}

	fmt.Printf("Binomial took %s\n", time.Since(start))
	return res
}

// day 4 part 2
// iterate through each field in xmas matrix
// dfs recursively in all diagonal directions
// if dfs returns 1 check corresponding cross starts
func partTwo(contents [][]string) int {
	defer utils.TimeTrack(time.Now(), "partTwo")
	rows, cols := len(contents), len(contents[0])
	res := 0

	for r := range rows {
		for c := range cols {
			for i, dir := range diagonals {
				if dfs(contents, r, c, dir, 0, "MAS") == 1 {
					for _, modifier := range crossCheck[i] {
						res += dfs(
							contents,
							r+modifier.Row,
							c+modifier.Col,
							diagonals[modifier.Dir],
							0,
							"MAS")
					}
				}
			}
		}
	}

	// divide res by 2 since we find each cross twice
	// probably room for some improvement
	return res / 2
}

func main() {
	// load contents for both parts as
	// rune matrix
	contents, err := utils.GetFileContentsAsStringMatrix("./input.txt", ``)
	utils.Check(err)

	xMasCount := partOne(contents)
	fmt.Println("XMAS count ", xMasCount)

	crossMasCount := partTwo(contents)
	fmt.Println("Cross MAS count ", crossMasCount)
}
