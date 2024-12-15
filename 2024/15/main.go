package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/JohnBra/aoc-2024/internal/utils"
)

var dir = map[rune][2]int{
	'^': {-1, 0},
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
}

func moveRobot(matrix [][]rune, r, c int, move rune) bool {
	// should never be out of bounds
	if utils.IsOutOfBounds(matrix, r, c) || matrix[r][c] == '#' {
		return false
	}

	if matrix[r][c] == '.' {
		return true
	}

	nextR, nextC := r+dir[move][0], c+dir[move][1]
	push := moveRobot(matrix, nextR, nextC, move)
	if push {
		matrix[nextR][nextC] = matrix[r][c]
	}
	return push
}

func getRobotPosition(matrix [][]rune) [2]int {
	for r := range len(matrix) {
		for c := range len(matrix[0]) {
			if matrix[r][c] == '@' {
				return [2]int{r, c}
			}
		}
	}
	return [2]int{}
}

func partOne(matrix [][]rune, moves []rune) int {
	res := 0
	robot := getRobotPosition(matrix)

	for _, m := range moves {
		nextR, nextC := robot[0]+dir[m][0], robot[1]+dir[m][1]
		if moveRobot(matrix, nextR, nextC, m) {
			matrix[robot[0]][robot[1]] = '.'
			robot[0], robot[1] = nextR, nextC
			matrix[robot[0]][robot[1]] = '@'
		}

	}

	for r := range len(matrix) {
		for c := range len(matrix[0]) {
			if matrix[r][c] == 'O' {
				res += r*100 + c
			}
		}
	}

	return res
}

func parseInput(filepath string) ([][]rune, []rune, error) {
	var matrix [][]rune
	var moves []rune

	file, err := os.Open(filepath)

	if err != nil {
		return matrix, moves, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	parseMap := true

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parseMap = false
		}

		if parseMap {
			matrix = append(matrix, []rune(line))
		} else {
			moves = append(moves, []rune(line)...)
		}
	}

	if err := scanner.Err(); err != nil {
		return matrix, moves, err
	}

	return matrix, moves, nil
}

func main() {
	matrix, moves, err := parseInput(utils.GetPuzzleInputSrc())
	utils.Check(err)

	partOneRes := partOne(matrix, moves)
	fmt.Println("Part one res:", partOneRes)
}
