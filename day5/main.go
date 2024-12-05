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
func partOne(graph map[int]ds.Set[int], updates [][]int) (int, [][]int) {
	var invalidUpdates [][]int
	validRes := 0
	preceding := ds.NewSet[int]()

	for _, update := range updates {
		for i, page := range update {
			if preceding.HasIntersection(graph[page]) {
				invalidUpdates = append(invalidUpdates, update)
				break
			}

			preceding.Add(page)

			if i == len(update)-1 {
				validRes += update[len(update)/2]
			}
		}
		clear(preceding)
	}
	return validRes, invalidUpdates
}

// day5 part 2
// create graph / adjacency list of all page edges
// get topological order for verteces in update
// sort update list by topological order conserving order of update
// for items that have equal or no ingoing edges
// accumulate middle value of corrected updates
func partTwo(graph ds.Graph[int], updates [][]int) int {
	res := 0

	for _, update := range updates {
		order := graph.TopologicalOrder(update)
		corrected := utils.SortListByOrder(update, order)

		res += corrected[len(corrected)/2]
	}

	return res
}

func main() {
	edges, updates, err := getIntMatricesFromContents("./input.txt")
	utils.Check(err)
	// all items in set must be after key of map
	graph := ds.NewGraph(edges, []int{})

	validRes, invalidUpdates := partOne(graph, updates)
	fmt.Printf("Valid updates %d\n", validRes)

	invalidRes := partTwo(graph, invalidUpdates)
	fmt.Printf("Invalid updates %d\n", invalidRes)
}
