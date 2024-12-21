package main

import (
	"fmt"
	"math"

	ds "github.com/JohnBra/aoc-2024/internal/datastructures"
	"github.com/JohnBra/aoc-2024/internal/utils"
)

type Coord struct {
	r, c int
}

type QItem struct {
	dist         int
	cstart, cend Coord
	ctime        int
	r, c         int
}

type SItem struct {
	r, c         int
	cstart, cend Coord
	ctime        int
}

func findCheat(matrix [][]rune, dist map[Coord]int, sr, sc, baseline, maxCheat int) int {
	const save int = 100
	res := ds.NewSet[[2]Coord]()
	// SItem: r, c, cheat start, cheat end, cheat time
	visit := ds.NewSet[SItem]()
	// QItem: distance, cheat start,cheat end,cheat time, r, c
	q := ds.NewDeque([]QItem{{
		dist:   0,
		cstart: Coord{-1, -1},
		cend:   Coord{-1, -1},
		ctime:  math.MinInt,
		r:      sr,
		c:      sc,
	}})

	for q.Len() > 0 {
		cur, _ := q.PopFront()

		if cur.dist >= baseline-save {
			continue
		}

		if matrix[cur.r][cur.c] == 'E' {
			if cur.cend.r == -1 && cur.cend.c == -1 {
				cur.cend = Coord{cur.r, cur.c}
			}

			validCheat := [2]Coord{cur.cstart, cur.cend}

			if cur.dist <= baseline-save {
				res.Add(validCheat)
			}
		}

		sitem := SItem{cur.r, cur.c, cur.cstart, cur.cend, cur.ctime}
		if visit.Contains(sitem) {
			continue
		}
		visit.Add(sitem)

		if cur.cstart.r == -1 && cur.cstart.c == -1 { // start cheat
			q.PushBack(QItem{
				cur.dist,
				Coord{cur.r, cur.c},
				Coord{-1, -1},
				maxCheat,
				cur.r,
				cur.c,
			})
		}

		if cur.ctime != math.MinInt && matrix[cur.r][cur.c] != '#' {
			if dist[Coord{cur.r, cur.c}] <= baseline-save-cur.dist {
				res.Add([2]Coord{cur.cstart, {cur.r, cur.c}})
			}
		}

		if cur.ctime == 0 {
			continue
		} else {
			for _, dir := range utils.Axes {
				nr, nc := cur.r+dir[0], cur.c+dir[1]
				if utils.IsOutOfBounds(matrix, nr, nc) {
					continue
				}

				if cur.ctime != math.MinInt {
					q.PushBack(QItem{
						cur.dist + 1,
						cur.cstart,
						Coord{-1, -1},
						cur.ctime - 1,
						nr,
						nc,
					})
				} else if matrix[nr][nc] != '#' {
					q.PushBack(QItem{
						cur.dist + 1,
						cur.cstart,
						cur.cend,
						cur.ctime,
						nr,
						nc,
					})
				}
			}
		}
	}

	return len(res)
}

func solve(matrix [][]rune) (int, int) {
	// res, start r/c, end r/c race time without cheat
	sr, sc, er, ec := 0, 0, 0, 0

	// find start/end and all inner walls
	for r := range len(matrix) {
		for c := range len(matrix[0]) {
			if matrix[r][c] == 'S' {
				sr, sc = r, c
			}

			if matrix[r][c] == 'E' {
				er, ec = r, c
			}

		}
	}

	q := ds.NewDeque([][3]int{{er, ec, 0}})
	dist := map[Coord]int{}

	for q.Len() > 0 {
		item, _ := q.PopFront()

		cur := Coord{item[0], item[1]}
		if _, ok := dist[cur]; ok {
			continue
		}

		dist[cur] = item[2]

		for _, dir := range utils.Axes {
			nr, nc := item[0]+dir[0], item[1]+dir[1]
			next := [3]int{nr, nc, item[2] + 1}

			if nr >= 0 && nr < len(matrix) &&
				nc >= 0 && nc < len(matrix[0]) &&
				matrix[nr][nc] != '#' {

				q.PushBack(next)
			}
		}
	}

	p1 := findCheat(matrix, dist, sr, sc, dist[Coord{sr, sc}], 2)
	p2 := findCheat(matrix, dist, sr, sc, dist[Coord{sr, sc}], 20)

	return p1, p2
}

func main() {
	matrix, err := utils.GetFileContentsAsRuneMatrix(utils.GetPuzzleInputSrc())
	utils.Check(err)

	partOneRes, partTwoRes := solve(matrix)
	fmt.Println("Part one res", partOneRes)
	fmt.Println("Part two res", partTwoRes)
}
