package main

import (
	"fmt"
	"github.com/JohnBra/aoc-2024/utils"
)

func partOne(contents [][]int) int {
	res := 0
	for i := 0; i < len(contents); i++ {
		report := contents[i]
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
			res += 1
		}
	}
	return res
}

func main() {
	contents, err := utils.GetFileContentsAsInts("./input.txt")
	utils.Check(err)
	reports := partOne(contents)
	fmt.Printf("Safe reports: %d", reports)
}
