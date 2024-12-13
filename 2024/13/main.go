package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/JohnBra/aoc-2024/internal/utils"
)

func dfs(dp map[[2]int]int, g [6]int, a, b int) int {
	if a > 100 || b > 100 || a*g[0]+b*g[2] > g[4] || a*g[1]+b*g[3] > g[5] {
		return math.MaxInt
	}

	cur := [2]int{a, b}
	memo, ok := dp[cur]
	if ok {
		return memo
	}

	if a*g[0]+b*g[2] == g[4] && a*g[1]+b*g[3] == g[5] {
		return a*3 + b
	}

	one, two := dfs(dp, g, a+1, b), dfs(dp, g, a, b+1)

	if one < two {
		dp[cur] = one
	} else {
		dp[cur] = two
	}

	return dp[cur]
}

func partOne(games [][6]int) int {
	res := 0

	for _, game := range games {
		dp := map[[2]int]int{}
		cost := dfs(dp, game, 0, 0)
		if cost != math.MaxInt {
			res += cost
		}
	}

	return res
}

func getInput(filepath string) ([][6]int, error) {
	// slice of [AX, AY, BX, BY, PX, PY]
	var res [][6]int

	f, err := os.Open(filepath)
	if err != nil {
		return res, err
	}

	defer f.Close()
	b := new(strings.Builder)
	io.Copy(b, f)
	games := strings.Split(b.String(), "\n\n")
	re := regexp.MustCompile(`\d+`)

	for _, s := range games {
		strings := re.FindAllString(s, -1)
		var nums [6]int
		for i, num := range strings {
			n, err := strconv.Atoi(num)
			if err != nil {
				return res, err
			}
			nums[i] = n
		}
		res = append(res, nums)
	}

	return res, nil
}

func main() {
	input, err := getInput(utils.GetPuzzleInputSrc())
	utils.Check(err)

	partOneRes := partOne(input)
	fmt.Println("Fewest tokens to win prizes", partOneRes)
}
