package main

import (
	"fmt"
	"github.com/JohnBra/aoc-2024/utils"
	"math"
	"regexp"
	"slices"
	"strconv"
)

// day1 part 1
// read lines from input
// add two lists left / right respectively
// sort both lists asc
// init res = 0
// iterate over list abs subtract both values and add to res
func main() {
	contents, err := utils.GetFileContentsAsRunes("./input.txt")
	utils.Check(err)
	re := regexp.MustCompile(`\d+`)
	var left []int
	var right []int

	for r := 0; r < len(contents); r++ {
		line := string(contents[r])
		nums := re.FindAllString(line, -1)

		leftVal, err := strconv.Atoi(nums[0])
		utils.Check(err)
		left = append(left, leftVal)

		rightVal, err := strconv.Atoi(nums[1])
		utils.Check(err)
		right = append(right, rightVal)
	}

	slices.Sort(left)
	slices.Sort(right)

	res := 0
	for i := 0; i < len(left); i++ {
		res += int(math.Abs(float64(left[i] - right[i])))
	}

	fmt.Printf("Distance is: %d", res)
}
