package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	ds "github.com/JohnBra/aoc-2024/internal/datastructures"
	"github.com/JohnBra/aoc-2024/internal/utils"
)

func partOne(edges [][2]string) int {
	res := 0

	graph := ds.NewGraphUndirected(edges, []string{})
	eset := ds.NewSet(edges...)
	triplets := ds.Set[[3]string]{}

	for e1, e1Neighbors := range graph {
		for _, e2 := range e1Neighbors.Members() {
			for _, e3 := range graph[e2].Members() {
				if eset.Contains([2]string{e1, e3}) || eset.Contains([2]string{e3, e1}) {
					triplet := []string{e1, e2, e3}
					slices.Sort(triplet)
					triplets.Add([3]string(triplet))
				}
			}
		}
	}

	for _, t := range triplets.Members() {
		if strings.HasPrefix(t[0], "t") || strings.HasPrefix(t[1], "t") || strings.HasPrefix(t[2], "t") {
			res += 1
		}
	}

	return res
}

func parseLine(line string) ([2]string, error) {
	re := regexp.MustCompile(`[a-z]+`)
	return [2]string(re.FindAllString(line, -1)), nil
}

func main() {
	input, err := utils.GetSliceOfSlicesFromFile(utils.GetPuzzleInputSrc(), parseLine)
	utils.Check(err)

	partOneRes := partOne(input)
	fmt.Println("Part one res", partOneRes)
}
