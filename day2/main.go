package main

import (
	"fmt"
	"github.com/JohnBra/aoc-2024/internal/utils"
)

// day2 part 1
// init asc/desc flag
// iterate through each report (int slice)
// check if num is asc/desc and in range for safe series (+-3 asc/desc respectively)
// accumulate safe reports and return
func partOne(contents [][]int) int {
	res := 0
	for i := 0; i < len(contents); i++ {
		if isSafe(contents[i]) {
			res += 1
		}
	}
	return res
}

func isSafe(report []int) bool {
	asc, desc := true, true
	for j := 1; j < len(report); j++ {
		if report[j] <= report[j-1] || report[j] > report[j-1]+3 {
			asc = false
		}

		if report[j] >= report[j-1] || report[j] < report[j-1]-3 {
			desc = false
		}
	}

	if asc || desc {
		return true
	}
	return false
}

// day2 part 2
// for each report check if is safe
// if not create all possible variations missing one level
// check if that is safe
// accumulate result and return
func partTwo(contents [][]int) int {
	res := 0
	for i := 0; i < len(contents); i++ {
		report := contents[i]
		// create all possible variations of report
		// if one of variations is safe, add 1+ to res
		if isSafe(report) {
			res += 1
			continue
		} else {
			for j := 0; j < len(report); j++ {
				var tmp []int
				for t := 0; t < len(report); t++ {
					if t == j {
						continue
					}
					tmp = append(tmp, report[t])
				}
				// take all values except report[j]
				if isSafe(tmp) {
					res += 1
					break
				}
			}
		}
	}
	return res
}

func main() {
	// for both parts
	// read lines from input
	// get input as variable length slice of slices of ints
	contents, err := utils.GetFileContentsAsInts("./input.txt")
	utils.Check(err)

	reports := partOne(contents)
	fmt.Printf("Safe reports: %d\n", reports)

	reportsDampened := partTwo(contents)
	fmt.Printf("Safe reports dampened: %d\n", reportsDampened)

}
