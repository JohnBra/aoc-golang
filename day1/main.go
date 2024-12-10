package main

import (
	"fmt"
	"github.com/JohnBra/aoc-2024/internal/utils"
	"math"
	"slices"
)

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

	contents, err := utils.GetFileContentsAsIntMatrix(utils.GetPuzzleInputSrc())
	utils.Check(err)
	lists, err := utils.Zip(contents)
	utils.Check(err)

	if len(lists) != 2 {
		panic(fmt.Errorf("Input should have two columns exactly"))
	}

	// day1 part 1
	// sort both lists asc
	// init distance = 0
	// iterate over list abs subtract both values and add to distance
	distance := partOne(lists[0], lists[1])
	fmt.Printf("Distance is: %d\n", distance)

	// day1 part 2
	// create map for right list {num -> count}
	// init similarity = 0
	// iterate over left list and add leftVal * map[leftVal] to similarity
	similarity := partTwo(lists[0], lists[1])
	fmt.Printf("Similarity score is: %d\n", similarity)

}
