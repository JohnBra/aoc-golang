package main

import (
	"fmt"
	"github.com/JohnBra/aoc-2024/internal/utils"
	"regexp"
	"strconv"
	"strings"
)

// takes in text, extracts nums, converts and multiplies them
func multiplyNumbersInText(text string) int {
	res := 1
	numsRe := regexp.MustCompile(`\d+`)
	nums := numsRe.FindAllString(text, -1)

	if len(nums) == 0 {
		return 0
	}

	for _, v := range nums {
		n, err := strconv.Atoi(v)
		utils.Check(err)
		res *= n
	}

	return res
}

// day3 part 1
// extract all "mul" product tuples from text with regex
// multiply all numbers in product and add to res
func partOne(contents string) int {
	res := 0
	productsRe := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))`)
	products := productsRe.FindAllString(contents, -1)

	for i := 0; i < len(products); i++ {
		res += multiplyNumbersInText(products[i])
	}

	return res
}

// day3 part 2
// extract all enablers "do()" and disablers "don't()" indices
// from text with regex
// create intervals for enabled multiplication tuples in text (start, stop)
// for each interval call partOne
func partTwo(contents string) int {
	res := 0
	dontRe := regexp.MustCompile(`(don't\(\))`)
	doRe := regexp.MustCompile(`(do\(\))`)

	startIndices := doRe.FindAllStringIndex(contents, -1)
	stopIndices := dontRe.FindAllStringIndex(contents, -1)
	stopIndices = append(stopIndices, []int{len(contents), len(contents)})

	var intervals [][]int
	intervals = append(intervals, []int{0, stopIndices[0][0]})
	i := 1
	for j, v := range startIndices {
		// skip this start value if it is within the last valid interval
		if v[0] < intervals[len(intervals)-1][1] {
			continue
		}

		// increment pointer to valid stop value as long as it is smaller
		// than current valid start value
		for i < len(stopIndices) && stopIndices[i][1] < v[0] {
			i += 1
		}

		// add valid interval
		intervals = append(intervals, []int{v[0], stopIndices[i][1]})

		if i == len(stopIndices) && j < len(startIndices) {
			break
		}
	}

	for _, v := range intervals {
		res += partOne(strings.TrimSpace(contents[v[0]:v[1]]))
	}

	return res
}

func parseLine(line string) (string, error) {
	return line, nil
}

func main() {
	// load file contents as one line string
	contents, err := utils.AccumulateLineResultFromFile(utils.GetPuzzleInputSrc(), parseLine)
	utils.Check(err)
	multiplications := partOne(contents)
	fmt.Printf("Multiplications: %d\n", multiplications)

	enabledMultiplications := partTwo(contents)
	fmt.Printf("Enabled multiplications: %d\n", enabledMultiplications)
}
