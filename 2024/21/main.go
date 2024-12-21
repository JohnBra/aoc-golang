package main

import (
	"fmt"

	"github.com/JohnBra/aoc-2024/internal/utils"
)

func partOne(input [][]rune) int {
	res := 0
	return res
}

func main() {
	input, err := utils.GetFileContentsAsRuneMatrix(utils.GetPuzzleInputSrc())
	utils.Check(err)

	partOneRes := partOne(input)
	fmt.Println("Part one res", partOneRes)
}
