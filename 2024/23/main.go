package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	ds "github.com/JohnBra/aoc-2024/internal/datastructures"
	"github.com/JohnBra/aoc-2024/internal/utils"
)

func search(graph ds.Graph[string], interconnected ds.Set[[520]string], node string, req ds.Set[string]) {
	key := sortedTuple(req.Members())
	if interconnected.Contains(key) {
		return
	}
	interconnected.Add(key)

	for _, nei := range graph[node].Members() {
		if req.Contains(nei) {
			continue
		}

		fulfills := true
		for _, q := range req.Members() {
			fulfills = fulfills && graph[q].Contains(nei)
		}

		if !fulfills {
			continue
		}

		nreq := ds.NewSet(req.Members()...)
		nreq.Add(nei)
		search(graph, interconnected, nei, nreq)
	}
}

func partOne(graph ds.Graph[string], eset ds.Set[[2]string]) int {
	res := 0
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

func partTwo(graph ds.Graph[string]) string {
	interconnected := ds.NewSet[[520]string]()

	for node := range graph {
		search(graph, interconnected, node, ds.NewSet(node))
	}

	res := []string{}
	for _, arr := range interconnected.Members() {
		slice := []string{}
		for _, v := range arr {
			if v != "" {
				slice = append(slice, v)
			}
		}
		if len(slice) > len(res) {
			res = slice
		}
	}

	slices.Sort(res)

	return strings.Join(res, ",")
}

func parseLine(line string) ([2]string, error) {
	re := regexp.MustCompile(`[a-z]+`)
	return [2]string(re.FindAllString(line, -1)), nil
}

func main() {
	input, err := utils.GetSliceOfSlicesFromFile(utils.GetPuzzleInputSrc(), parseLine)
	utils.Check(err)

	graph := ds.NewGraphUndirected(input, []string{})
	eset := ds.NewSet(input...) // set of all edges

	partOneRes := partOne(graph, eset)
	fmt.Println("Part one res", partOneRes)

	partTwoRes := partTwo(graph)
	fmt.Println("Part two res", partTwoRes)

}

// creates tuple with 520 strings (max interconnected nodes in input)
//
// a bit hacky, but works
func sortedTuple(arr []string) [520]string {
	slices.Sort(arr)

	narr := [520]string{}
	copy(narr[:], arr)
	return narr
}
