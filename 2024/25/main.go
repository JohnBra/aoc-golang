package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/JohnBra/aoc-2024/internal/utils"
)

func partOne(locks, keys [][5]int) int {
	res := 0

	for _, key := range keys {
		for _, lock := range locks {
			fits := true
			for i := range len(key) {
				if key[i]+lock[i] > 5 {
					fits = false
				}
			}

			if fits {
				res += 1
			}
		}
	}

	return res
}

func parseInput(filepath string) ([][5]int, [][5]int, error) {
	locks, keys := [][5]int{}, [][5]int{}

	f, err := os.Open(filepath)
	if err != nil {
		return locks, keys, err
	}
	defer f.Close()

	b := new(strings.Builder)
	io.Copy(b, f)

	strs := strings.Split(strings.TrimSuffix(b.String(), "\n"), "\n\n")
	for _, str := range strs {
		cnt := [5]int{}
		lock := true
		for i, line := range strings.Split(str, "\n") {
			if i == 0 && line == "#####" {
				// do nothing
			} else if i == 6 && line == "#####" {
				lock = false
			} else {
				for j, c := range line {
					if c == '#' {
						cnt[j] += 1
					}
				}
			}
		}
		if lock {
			locks = append(locks, cnt)
		} else {
			keys = append(keys, cnt)
		}
	}

	return locks, keys, nil
}

func main() {
	locks, keys, err := parseInput(utils.GetPuzzleInputSrc())
	utils.Check(err)

	partOneRes := partOne(locks, keys)
	fmt.Println("Part one res", partOneRes)
}
