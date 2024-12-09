package main

import (
	"fmt"
	"strings"
	"time"

	ds "github.com/JohnBra/aoc-2024/internal/datastructures"
	"github.com/JohnBra/aoc-2024/internal/utils"
)

type Coord struct {
	r, c int
}

type Direction struct {
	mod  [2]int
	next string
	idx  int
}

type CoordWithDirection struct {
	r, c, d int
}

var direction map[string]Direction = map[string]Direction{
	"^": {[2]int{-1, 0}, ">", 0},
	">": {[2]int{0, 1}, "v", 1},
	"v": {[2]int{1, 0}, "<", 2},
	"<": {[2]int{0, -1}, "^", 3},
}

func partOne(board [][]string) (int, bool) {
	rows, cols := len(board), len(board[0])
	q := ds.NewDeque([]Coord{})
	visit := ds.NewSet[Coord]()
	cycle := ds.NewSet[CoordWithDirection]()
	var dir Direction

	// find caret
	// determine direction
	for r := range rows {
		for c := range cols {
			if board[r][c] != "." && board[r][c] != "#" {
				dir = direction[board[r][c]]
				q.PushBack(Coord{r, c})
				visit.Add(Coord{r, c})
				cycle.Add(CoordWithDirection{r, c, dir.idx})
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

		// Only add next field if can actually walk it
		q.PushBack(next)
		visit.Add(next)

		// For part two track cycle
		nextWithDir := CoordWithDirection{next.r, next.c, dir.idx}
		if cycle.Contains(nextWithDir) {
			return len(visit), true
		}
		cycle.Add(nextWithDir)
	}

	return len(visit), false
}

func partTwo(board [][]string) int {
	defer utils.TimeTrack(time.Now(), "Day 8 part 2")
	// brute force
	// put an obstacle on each field (one at a time; if not alreay one)
	// execute part one
	// if on field with same direction twice -> cycle
	rows, cols := len(board), len(board[0])
	res := 0
	for r := range rows {
		for c := range cols {
			if board[r][c] == "." {
				board[r][c] = "#"
				_, cycle := partOne(board)
				if cycle {
					res += 1
				}
				board[r][c] = "."
			}
		}
	}

	return res
}

func parseLine(line string) ([]string, error) {
	return strings.Split(line, ``), nil
}

func main() {
	contents, err := utils.GetSliceOfSlicesFromFile("./input.txt", parseLine)
	utils.Check(err)

	distinctFields, _ := partOne(contents)
	fmt.Println("Distinct fields: ", distinctFields)

	obstructions := partTwo(contents)
	fmt.Println("Possible obstructions: ", obstructions)
}
