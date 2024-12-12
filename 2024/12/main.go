package main

import (
	"fmt"

	ds "github.com/JohnBra/aoc-2024/internal/datastructures"
	"github.com/JohnBra/aoc-2024/internal/utils"
)

func dfs(matrix [][]rune, visit ds.Set[[2]int], r, c int, prev rune) (int, int, int) {
	if utils.IsOutOfBounds(matrix, r, c) || visit.Contains([2]int{r, c}) || matrix[r][c] != prev {
		return 0, 0, 0
	}

	visit.Add([2]int{r, c})
	area, perimeter, sides := 1, 0, 0

	for i := 0; i < len(utils.Directions); i += 2 {
		row, col := r+utils.Directions[i][0], c+utils.Directions[i][1]
		diagNeigh := utils.Directions[(i+1)%8]
		axisNeigh := utils.Directions[(i+2)%8]

		if utils.IsOutOfBounds(matrix, row, col) || matrix[r][c] != matrix[row][col] {
			perimeter += 1
		}

		// check angles pointing outside with horizontal/vertical neighbor of original field
		if (utils.IsOutOfBounds(matrix, row, col) || matrix[r][c] != matrix[row][col]) &&
			(utils.IsOutOfBounds(matrix, r+axisNeigh[0], c+axisNeigh[1]) || matrix[r][c] != matrix[r+axisNeigh[0]][c+axisNeigh[1]]) {
			sides += 1
		}

		// check angles pointing inside with diagonal and horizontal/vertical neighbor of original field
		if (!utils.IsOutOfBounds(matrix, row, col) && matrix[r][c] == matrix[row][col]) &&
			(!utils.IsOutOfBounds(matrix, r+diagNeigh[0], c+diagNeigh[1]) && matrix[r][c] != matrix[r+diagNeigh[0]][c+diagNeigh[1]]) &&
			(!utils.IsOutOfBounds(matrix, r+axisNeigh[0], c+axisNeigh[1]) && matrix[r][c] == matrix[r+axisNeigh[0]][c+axisNeigh[1]]) {
			sides += 1
		}

		a, p, s := dfs(matrix, visit, row, col, matrix[r][c])
		area += a
		perimeter += p
		sides += s
	}

	return area, perimeter, sides
}

func findFields(matrix [][]rune) (int, int) {
	partOneRes, partTwoRes := 0, 0
	rows, cols := len(matrix), len(matrix[0])
	visit := ds.NewSet[[2]int]()

	for r := range rows {
		for c := range cols {
			if !visit.Contains([2]int{r, c}) {
				area, perimeter, sides := dfs(matrix, visit, r, c, matrix[r][c])
				partOneRes += area * perimeter
				partTwoRes += area * sides
			}
		}
	}

	return partOneRes, partTwoRes
}

func parseLine(line string) ([]rune, error) {
	return []rune(line), nil
}

func main() {
	matrix, err := utils.GetSliceOfSlicesFromFile(utils.GetPuzzleInputSrc(), parseLine)
	utils.Check(err)

	partOneRes, partTwoRes := findFields(matrix)
	fmt.Println("Part one total price", partOneRes)
	fmt.Println("Part two total price", partTwoRes)
}
