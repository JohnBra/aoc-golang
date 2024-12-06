package main

import (
	"fmt"

	ds "github.com/JohnBra/aoc-2024/internal/datastructures"
	"github.com/JohnBra/aoc-2024/internal/utils"
)

type Coord struct {
	r, c int
}

type Direction struct {
	mod  [2]int
	next string
}

var direction map[string]Direction = map[string]Direction{
	"^": {[2]int{-1, 0}, ">"},
	">": {[2]int{0, 1}, "v"},
	"v": {[2]int{1, 0}, "<"},
	"<": {[2]int{0, -1}, "^"},
}

var d map[int][3]int = map[int][3]int{
	0: {-1, 0, 1},
	1: {0, 1, 2},
	2: {1, 0, 3},
	3: {0, -1, 0},
}

func partOne(board [][]string) int {
	rows, cols := len(board), len(board[0])
	q := ds.NewDeque([]Coord{})
	set := ds.NewSet[Coord]()
	var dir Direction

	// find caret
	// determine direction
	for r := range rows {
		for c := range cols {
			if board[r][c] != "." && board[r][c] != "#" {
				dir = direction[board[r][c]]
				q.PushBack(Coord{r, c})
				set.Add(Coord{r, c})
			}
		}
	}

	for q.Len() > 0 {
		coord, _ := q.PopFront()
		next := Coord{coord.r + dir.mod[0], coord.c + dir.mod[1]}

		if next.r < 0 || next.r == rows || next.c < 0 || next.c == cols {
			break
		}

		if board[next.r][next.c] == "#" {
			dir = direction[dir.next]
			q.PushBack(coord)
			continue
		}

		q.PushBack(next)
		set.Add(next)
	}

	return len(set)
}

func partTwo(board [][]int) int {
	// brute force
	// put an obstacle on each field (one at a time; if not alreay one)
	// execute part one
	// if on field with same direction twice -> cycle

	return 0
}

func main() {
	contents, err := utils.GetFileContentsAsStringMatrix("test.txt", ``)
	utils.Check(err)

	distinctFields := partOne(contents)
	fmt.Println("Distinct fields: ", distinctFields)
}
