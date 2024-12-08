package main

import (
	"fmt"
	"math"
	"strings"

	ds "github.com/JohnBra/aoc-2024/internal/datastructures"
	"github.com/JohnBra/aoc-2024/internal/utils"
)

type Coord struct {
	r, c int
}

// day8 part 1
// add all antennas to a map: antenna -> coord and
// all antenna coords to a set
// iterate through all keys (antennas) in map, calculate
// potential antinodes coords and add coords to set if in bounds of matrix or
// in antenna cord set
// return length of antinode set
func partOne(matrix [][]string) int {
	rows, cols := len(matrix), len(matrix[0])
	antennas := map[string]ds.Set[Coord]{}
	antinodes := ds.NewSet[Coord]()

	for r := range rows {
		for c := range cols {
			field := matrix[r][c]
			if field != "." {
				_, ok := antennas[field]
				if !ok {
					antennas[field] = ds.NewSet[Coord]()
				}

				antennas[field].Add(Coord{r, c})
			}
		}
	}

	for _, c := range antennas {
		if len(c) < 2 {
			continue
		}

		coords := c.Members()

		for i := 0; i < len(coords)-1; i++ {
			for j := i + 1; j < len(coords); j++ {
				rmod := int(math.Abs(float64(coords[i].r) - float64(coords[j].r)))
				cmod := int(math.Abs(float64(coords[i].c) - float64(coords[j].c)))

				// antinode for i row and column
				// same for antinode j
				var air, aic, ajr, ajc int

				if coords[i].r < coords[j].r {
					air = coords[i].r - rmod
					ajr = coords[j].r + rmod
				} else {
					air = coords[i].r + rmod
					ajr = coords[j].r - rmod
				}

				if coords[i].c < coords[j].c {
					aic = coords[i].c - cmod
					ajc = coords[j].c + cmod
				} else {
					aic = coords[i].c + cmod
					ajc = coords[j].c - cmod
				}

				ai, aj := Coord{air, aic}, Coord{ajr, ajc}
				if ai.c >= 0 && ai.c < cols && ai.r >= 0 && ai.r < rows {
					antinodes.Add(ai)
				}

				if aj.c >= 0 && aj.c < cols && aj.r >= 0 && aj.r < rows {
					antinodes.Add(aj)
				}
			}
		}
	}

	return len(antinodes)
}

func parseLine(line string) ([]string, error) {
	return strings.Split(line, ``), nil
}

func main() {
	matrix, err := utils.GetSliceOfSlicesFromFile("./8.in", parseLine)
	utils.Check(err)

	partOneRes := partOne(matrix)
	fmt.Println("Antinode count ", partOneRes)
}
