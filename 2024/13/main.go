package main

import (
	"fmt"
	"io"
	"math"
	"math/big"
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

// day13 part 1
// recursively check combinations up until 100 or result reached
// memoize the results, return cache value if exist
// return minimum coin value
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

// day13 part 2
// Solves the linear equation by solving for a
// then solving for b and substituting
// modify each game result with +10000000000000 before calculating
// use big int, otherwise integer overflow
func partTwo(games [][6]int) string {
	res := big.NewInt(0)

	for _, g := range games {
		g[4] += 10000000000000
		g[5] += 10000000000000
		// 			t1			t2			t3			t4
		//a := (g[4]*g[3] - g[5]*g[2]) / (g[0]*g[3] - g[1]*g[2])
		//			t5			t6			t3			t4
		//b := (g[0]*g[5] - g[1]*g[4]) / (g[0]*g[3] - g[1]*g[2])

		t1 := big.NewInt(0).Mul(big.NewInt(int64(g[4])), big.NewInt(int64(g[3])))
		t2 := big.NewInt(0).Mul(big.NewInt(int64(g[5])), big.NewInt(int64(g[2])))
		t3 := big.NewInt(int64(g[0] * g[3]))
		t4 := big.NewInt(int64(g[1] * g[2]))
		t5 := big.NewInt(0).Mul(big.NewInt(int64(g[0])), big.NewInt(int64(g[5])))
		t6 := big.NewInt(0).Mul(big.NewInt(int64(g[1])), big.NewInt(int64(g[4])))
		t1Mt2 := big.NewInt(0).Sub(t1, t2)
		t3Mt4 := big.NewInt(0).Sub(t3, t4)
		t5Mt6 := big.NewInt(0).Sub(t5, t6)
		a, aRest := big.NewInt(0).DivMod(t1Mt2, t3Mt4, big.NewInt(1))
		b, bRest := big.NewInt(0).DivMod(t5Mt6, t3Mt4, big.NewInt(1))

		if aRest.String() == "0" && bRest.String() == "0" {
			a.Mul(a, big.NewInt(3))
			tokens := big.NewInt(0).Add(a, b)
			res.Add(res, tokens)
		}
	}

	return res.String()
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
	fmt.Println("Part one tokens", partOneRes)

	partTwoRes := partTwo(input)
	fmt.Println("Part two tokens", partTwoRes)
}
