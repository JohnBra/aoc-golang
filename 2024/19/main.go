package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/JohnBra/aoc-2024/internal/utils"
)

func dfs(patterns []string, dp map[string]int, target string) int {
	if res, ok := dp[target]; ok {
		return res
	}

	ans := 0
	if len(target) == 0 {
		ans = 1
	}

	for _, pattern := range patterns {
		if strings.HasPrefix(target, pattern) {
			ans += dfs(patterns, dp, target[len(pattern):])
		}
	}
	dp[target] = ans
	return ans
}

func solve(patterns, targets []string) (int, int) {
	valid, all := 0, 0
	dp := map[string]int{}

	for _, t := range targets {
		res := dfs(patterns, dp, t)
		if res > 1 {
			valid += 1
		}
		all += res
	}

	return valid, all
}

func parseInput(filepath string) ([]string, []string, error) {
	patterns, targets := []string{}, []string{}

	f, err := os.Open(filepath)
	if err != nil {
		return patterns, targets, err
	}
	defer f.Close()

	b := new(strings.Builder)
	io.Copy(b, f)

	input := strings.Split(b.String(), "\n\n")
	patterns = strings.Split(input[0], ", ")
	targets = strings.Split(strings.TrimRight(input[1], "\n"), "\n")

	return patterns, targets, nil
}

func main() {
	patterns, targets, err := parseInput(utils.GetPuzzleInputSrc())
	utils.Check(err)

	partOneRes, partTwoRes := solve(patterns, targets)
	fmt.Println("Part one res", partOneRes)
	fmt.Println("Part two res", partTwoRes)
}
