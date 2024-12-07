package main

import (
	"fmt"

	"github.com/JohnBra/aoc-2024/internal/utils"
)

func dfs(nums []int, i int, total int) bool {
	if i == len(nums) {
		return total == nums[0]
	}

	return dfs(nums, i+1, total+nums[i]) || dfs(nums, i+1, total*nums[i])
}

func partOne(equations [][]int) int {
	// backtrack until result found or all possible variations tested
	res := 0
	for _, equation := range equations {
		if dfs(equation, 1, 0) {
			res += equation[0]
		}
	}
	return res
}

func main() {
	contents, err := utils.GetFileContentsAsIntMatrix("./input.txt")
	utils.Check(err)

	calibration := partOne(contents)
	fmt.Println("Total calibration result: ", calibration)
}
