package main

import (
	"fmt"

	"github.com/JohnBra/aoc-2024/internal/utils"
)

func partOne(robots [][]int) int {
	// map has 103 rows and 101 columns
	rows, cols := 103, 101
	i := 100
	for i > 0 {
		for _, robot := range robots {
			robot[1] = (robot[1] + robot[3] + rows) % rows // row
			robot[0] = (robot[0] + robot[2] + cols) % cols // col
		}
		i--
	}

	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, robot := range robots {
		r, c := robot[1], robot[0]

		if r < rows/2 && c < cols/2 { //q1: r 0-50 c 0-49
			q1 += 1
		} else if r < rows/2 && c > cols/2 { //q2: r 0-50 c 51-101
			q2 += 1
		} else if r > rows/2 && c < cols/2 { //q3: r 52-103 c 0-49
			q3 += 1
		} else if r > rows/2 && c > cols/2 { //q4 r 52-103 c 51-101
			q4 += 1
		}
	}

	return q1 * q2 * q3 * q4
}

func main() {
	// slices of [c, r, cMod, rMod]
	robots, err := utils.GetFileContentsAsIntMatrix(utils.GetPuzzleInputSrc())
	utils.Check(err)

	partOneRes := partOne(robots)
	fmt.Println("Part one res", partOneRes)
}
