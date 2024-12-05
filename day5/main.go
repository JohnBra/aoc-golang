package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	ds "github.com/JohnBra/aoc-2024/internal/datastructures"
	"github.com/JohnBra/aoc-2024/internal/utils"
)

func getIntMatricesFromContents(filepath string) ([][]int, [][]int, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return nil, nil, err
	}

	defer file.Close()

	re := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(file)
	var pageOrder [][]int
	var updates [][]int
	section := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			section += 1
			continue
		}

		strings := re.FindAllString(line, -1)
		nums := make([]int, 0, len(strings))
		for _, v := range strings {
			num, err := strconv.Atoi(v)
			if err != nil {
				return nil, nil, err
			}
			nums = append(nums, num)
		}

		if section == 0 {
			pageOrder = append(pageOrder, nums)
		} else if section == 1 {
			updates = append(updates, nums)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return pageOrder, updates, nil
}

// day5 part 1
// update is correct if x comes before y
// iterate over all updates
// init new set for each update containing pages (int values)
// iterate over update slice
// check if there is an overlap between ordering set for cur val
// and preceding pages added to set, if yes break (ignore this update)
// if no add current number to set
func partOne(ordering map[int]ds.Set[int], updates [][]int) int {
	res := 0
	preceding := ds.NewSet[int]()
	for _, update := range updates {
		for i, page := range update {
			if preceding.HasIntersection(ordering[page]) {
				break
			}

			preceding.Add(page)

			if i == len(update)-1 {
				res += update[len(update)/2]
			}
		}
		preceding.Clear()
	}
	return res
}

func main() {
	o, updates, err := getIntMatricesFromContents("./input.txt")
	utils.Check(err)
	// all items in set must be after key of map
	ordering := map[int]ds.Set[int]{}

	for _, tuple := range o {
		_, ok := ordering[tuple[0]]
		if !ok {
			ordering[tuple[0]] = ds.NewSet(tuple[1])
		} else {
			ordering[tuple[0]].Add(tuple[1])
		}
	}

	validUpdates := partOne(ordering, updates)
	fmt.Printf("Valid updates %d\n", validUpdates)
}
