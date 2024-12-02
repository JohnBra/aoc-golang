package main

import (
	"fmt"
	"github.com/JohnBra/aoc-2024/internal/utils"
	"math"
	"regexp"
	"slices"
	"strconv"
)

// panics on error
func getLeftAndRightList(contents [][]rune) ([]int, []int) {
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
	return left, right
}

func partOne(left []int, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	distance := 0
	for i := 0; i < len(left); i++ {
		distance += int(math.Abs(float64(left[i] - right[i])))
	}
	return distance
}

func partTwo(left []int, right []int) int {
	count := make(map[int]int)
	for i := 0; i < len(right); i++ {
		count[right[i]] += 1
	}

	similarity := 0
	for i := 0; i < len(left); i++ {
		similarity += left[i] * count[left[i]]
	}

	return similarity
}

func main() {
	// for both parts
	// read lines from input
	// create two lists left / right respectively

	contents, err := utils.GetFileContentsAsRunes("./input.txt")
	utils.Check(err)
	left, right := getLeftAndRightList(contents)

	// day1 part 1
	// sort both lists asc
	// init distance = 0
	// iterate over list abs subtract both values and add to distance
	distance := partOne(left, right)
	fmt.Printf("Distance is: %d\n", distance)

	// day1 part 2
	// create map for right list {num -> count}
	// init similarity = 0
	// iterate over left list and add leftVal * map[leftVal] to similarity
	similarity := partTwo(left, right)
	fmt.Printf("Similarity score is: %d\n", similarity)

}
