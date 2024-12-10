package main

import (
	"fmt"
	"math/big"

	"github.com/JohnBra/aoc-2024/internal/utils"
)

func dfs(nums []int, i int, total int) bool {
	if i == len(nums) {
		return total == nums[0]
	}

	return dfs(nums, i+1, total+nums[i]) || dfs(nums, i+1, total*nums[i])
}

func dfsBigInt(nums []string, i int, total string) bool {
	if i == len(nums) {
		return total == nums[0]
	}

	t := new(big.Int)
	t, tOk := t.SetString(total, 10)
	if !tOk {
		panic(fmt.Sprintf("couldn't convert %s to bigint", total))
	}

	n := new(big.Int)
	n, nOk := n.SetString(nums[i], 10)
	if !nOk {
		panic(fmt.Sprintf("couldn't convert %s to bigint", nums[i]))
	}

	return (dfsBigInt(nums, i+1, big.NewInt(0).Add(t, n).String()) ||
		dfsBigInt(nums, i+1, big.NewInt(0).Mul(t, n).String()) ||
		dfsBigInt(nums, i+1, total+nums[i]))
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

func partTwo(equations [][]string) string {
	res := big.NewInt(0)

	for _, equation := range equations {
		if dfsBigInt(equation, 1, "0") {
			eq := new(big.Int)
			eq, ok := eq.SetString(equation[0], 10)
			if !ok {
				panic(fmt.Sprintf("couldn't convert %s to bigint", eq))
			}
			res = res.Add(res, eq)
		}
	}
	return res.String()
}

func main() {
	puzzleInput := utils.GetPuzzleInputSrc()
	intMatrix, err := utils.GetFileContentsAsIntMatrix(puzzleInput)
	utils.Check(err)

	partOneRes := partOne(intMatrix)
	fmt.Println("Total calibration result: ", partOneRes)

	bigintMatrix, err := utils.GetFileContentsAsNumberMatrix(puzzleInput)
	utils.Check(err)

	partTwoRes := partTwo(bigintMatrix)
	fmt.Println("Total calibration result (for real): ", partTwoRes)
}
