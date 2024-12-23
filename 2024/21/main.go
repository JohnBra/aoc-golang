package main

import (
	"container/heap"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	ds "github.com/JohnBra/aoc-2024/internal/datastructures"
	"github.com/JohnBra/aoc-2024/internal/utils"
)

type Coord struct {
	r, c int
}

type HItem struct {
	cost  int
	coord Coord
	path  string
	dir   int
}

func (h HItem) Priority() int {
	return h.cost
}

// key -> shortest paths to key
type ShortestPathsTo map[rune][]string

func findShortestPaths(matrix [][]rune, from, to rune) []string {
	// r, c
	visit := ds.NewSet[[3]int]()
	h := ds.NewHeap()
	curCost := math.MaxInt
	paths := []string{}
	path := ""

out:
	for r := range len(matrix) {
		for c := range len(matrix[0]) {
			if matrix[r][c] == from {
				heap.Push(h, HItem{0, Coord{r, c}, "", -1})
				break out
			}
		}
	}

	for h.Len() > 0 {
		item, _ := heap.Pop(h).(HItem)
		visit.Add([3]int{item.dir, item.coord.r, item.coord.c})

		if matrix[item.coord.r][item.coord.c] == to {
			if item.Priority() <= curCost {
				curCost = item.Priority()
				path = item.path + string('A')
				paths = append(paths, path)
			}
		}

		for o, dir := range utils.Axes {
			nr, nc := item.coord.r+dir[0], item.coord.c+dir[1]
			ncoord := Coord{nr, nc}

			if !utils.IsOutOfBounds(matrix, nr, nc) && !visit.Contains([3]int{o, nr, nc}) && matrix[nr][nc] != ' ' {
				heap.Push(h, HItem{
					item.cost + 1,
					ncoord,
					item.path + string(utils.AxesArrows[o]),
					o,
				})
			}
		}
	}

	//fmt.Printf("Shortest paths from %c to %c: %v\n", from, to, paths)
	return paths
}

func getKeyPadPathMap(keys []rune, pad [][]rune) map[rune]ShortestPathsTo {
	keyToKeys := map[rune]ShortestPathsTo{}

	for r := range len(pad) {
		for c := range len(pad[0]) {
			if pad[r][c] == ' ' {
				continue
			}

			m, ok := keyToKeys[pad[r][c]]
			if !ok {
				m = map[rune][]string{}
			}

			for _, k := range keys {
				m[k] = append(m[k], findShortestPaths(pad, pad[r][c], k)...)
			}
			keyToKeys[pad[r][c]] = m
		}
	}

	return keyToKeys
}

func sliceProduct(args ...[]string) []string {
	pools := args
	npools := len(pools)
	indices := make([]int, npools)
	result := make([]string, npools)

	for i := range result {
		if len(pools[i]) == 0 {
			return nil
		}
		result[i] = pools[i][0]
	}

	results := [][]string{result}

	for {
		i := npools - 1
		for ; i >= 0; i -= 1 {
			pool := pools[i]
			indices[i] += 1

			if indices[i] == len(pool) {
				indices[i] = 0
				result[i] = pool[0]
			} else {
				result[i] = pool[indices[i]]
				break
			}

		}

		if i < 0 {
			res := []string{}
			for _, s := range results {
				res = append(res, strings.Join(s, ""))
			}
			return res
		}

		newresult := make([]string, npools)
		copy(newresult, result)
		results = append(results, newresult)
	}

	return nil
}

func getPossibleRobotPaths(spaths map[rune]ShortestPathsTo, code string) []string {
	pos := 'A'
	options := [][]string{}

	for _, digit := range code {
		options = append(options, spaths[pos][digit])
		pos = digit
	}

	// cartesian product of all (minimal) robot 1 options
	return sliceProduct(options...)
}

func filterShortest(paths []string) ([]string, int) {
	minLen := math.MaxInt
	for _, p := range paths {
		minLen = utils.Min(minLen, len(p))
	}

	return utils.Filter(paths, func(s string) bool {
		return len(s) == minLen
	}), minLen
}

func solve(codes [][]rune) (int, int) {
	p1, p2 := 0, 0

	keys := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A'}
	dirKeys := []rune{'^', '>', 'v', '<', 'A'}
	keyPad := [][]rune{
		{'7', '8', '9'},
		{'4', '5', '6'},
		{'1', '2', '3'},
		{' ', '0', 'A'},
	}

	dirKeyPad := [][]rune{
		{' ', '^', 'A'},
		{'<', 'v', '>'},
	}

	re := regexp.MustCompile(`\d+`)
	keyPaths := getKeyPadPathMap(keys, keyPad)
	dirPaths := getKeyPadPathMap(dirKeys, dirKeyPad)
	minLen := 0

	for _, code := range codes {
		// keypad bot
		allPaths := getPossibleRobotPaths(keyPaths, string(code))
		prevPaths, _ := filterShortest(allPaths)

		// dir pad bots
		for range 2 {
			nextPaths := []string{}
			for _, p := range prevPaths {
				nextPaths = append(nextPaths, getPossibleRobotPaths(dirPaths, p)...)
			}

			prevPaths, minLen = filterShortest(nextPaths)
		}

		// add to result
		snum := re.FindAllString(string(code), -1)
		num, err := strconv.Atoi(snum[0])
		utils.Check(err)

		p1 += num * minLen
	}

	return p1, p2
}

func main() {
	input, err := utils.GetFileContentsAsRuneMatrix(utils.GetPuzzleInputSrc())
	utils.Check(err)

	partOneRes, partTwoRes := solve(input)
	fmt.Println("Part one res", partOneRes)
	fmt.Println("Part two res", partTwoRes)
}
