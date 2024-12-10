package main

import (
	"fmt"
	"strings"

	ds "github.com/JohnBra/aoc-2024/internal/datastructures"
	"github.com/JohnBra/aoc-2024/internal/utils"
)

type Coord struct {
	r, c int
}

func calcAntinode(a Coord, b Coord) Coord {
	rmod, cmod := a.r-b.r, a.c-b.c
	return Coord{a.r + rmod, a.c + cmod}
}

// day8 part 1
// add all antennas to a map: antenna -> []coord
// iterate through all keys (antennas) in map, calculate
// potential antinodes coords and add coords to set if in bounds of matrix
//
// day8 part 2
// add all antenna coords to antinode set AND
// when calculating antinodes also add new calculated antinode coord and closest
// coord used in previous antinode calculation to a queue if in bounds of matrix
// after initial antennas on map are done (part one) calculate antinodes
// for each antinode tuple in q and add new potential antinodes to q until
// calculated antinodes are out of bounds
// return length of antinode set
func getAntinodes(matrix [][]string) (int, int) {
	rows, cols := len(matrix), len(matrix[0])
	antennas := map[string]ds.Set[Coord]{}
	antinodesP1 := ds.NewSet[Coord]()
	antinodesP2 := ds.NewSet[Coord]()

	for r := range rows {
		for c := range cols {
			field := matrix[r][c]
			if field != "." {
				_, ok := antennas[field]
				if !ok {
					antennas[field] = ds.NewSet[Coord]()
				}

				antennas[field].Add(Coord{r, c})
				antinodesP2.Add(Coord{r, c})
			}
		}
	}

	for _, c := range antennas {
		if len(c) < 2 {
			continue
		}

		q := ds.NewDeque([][2]Coord{})
		coords := c.Members()
		// for part two add coord with rmod and cmod to a queue
		// pop from q as long items on it, use the modifier to
		// calculate a new coord for an antinode and put that coord
		// onto the q with the same modifier if not out of bounds

		for i := 0; i < len(coords)-1; i++ {
			for j := i + 1; j < len(coords); j++ {
				// a (8,8), b (9,9)
				// a => (8-9=-1, 8-9=-1), b => (9-8=1, 9-8=1)
				// a (3,8), b (5,3)
				// a=> (3-5=-2, 8-3=5), b => (8-5=3, 8-2=2)
				// antinode for i row and column
				// same for antinode j
				ai, aj := calcAntinode(coords[i], coords[j]), calcAntinode(coords[j], coords[i])
				if ai.c >= 0 && ai.c < cols && ai.r >= 0 && ai.r < rows {
					antinodesP1.Add(ai)
					antinodesP2.Add(ai)
					q.PushBack([2]Coord{ai, coords[i]})
				}

				if aj.c >= 0 && aj.c < cols && aj.r >= 0 && aj.r < rows {
					antinodesP1.Add(aj)
					antinodesP2.Add(aj)
					q.PushBack([2]Coord{aj, coords[j]})
				}
			}
		}

		// part two
		for q.Len() > 0 {
			coords, _ := q.PopFront()
			a := calcAntinode(coords[0], coords[1])
			if a.c >= 0 && a.c < cols && a.r >= 0 && a.r < rows {
				antinodesP2.Add(a)
				q.PushBack([2]Coord{a, coords[0]})
			}
		}

	}

	return len(antinodesP1), len(antinodesP2)
}

func parseLine(line string) ([]string, error) {
	return strings.Split(line, ``), nil
}

func main() {
	matrix, err := utils.GetSliceOfSlicesFromFile(utils.GetPuzzleInputSrc(), parseLine)
	utils.Check(err)

	partOneRes, partTwoRes := getAntinodes(matrix)
	fmt.Println("Part 1 antinode count", partOneRes)
	fmt.Println("Part 2 antinode count", partTwoRes)
}
