package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/JohnBra/aoc-2024/internal/utils"
)

func getComboOperand(op, a, b, c int) int {
	switch op {
	case 4:
		return a
	case 5:
		return b
	case 6:
		return c
	}
	return op
}

func partOne(nums []int) string {
	// nums: a: 0; b: 1; c: 2; program: 3...n-1
	a, b, c := nums[0], nums[1], nums[2]
	pr := nums[3:]
	res := ""

	i := 0
	for i+1 < len(pr) {
		switch pr[i] {
		case 0: // adv
			a = int(a / utils.Pow(2, getComboOperand(pr[i+1], a, b, c)))
		case 1: // bxl
			b = b ^ pr[i+1]
		case 2: // bst
			b = getComboOperand(pr[i+1], a, b, c) % 8
		case 3: // jnc
			if a != 0 {
				i = pr[i+1]
				continue
			}
		case 4: // bxc
			b = b ^ c
		case 5: // out
			res += fmt.Sprintf("%d,", getComboOperand(pr[i+1], a, b, c)%8)
		case 6: // bdv
			b = int(a / utils.Pow(2, getComboOperand(pr[i+1], a, b, c)))
		case 7: // cdv
			c = int(a / utils.Pow(2, getComboOperand(pr[i+1], a, b, c)))
		}
		i += 2
	}

	if len(res) > 0 {
		res = res[:len(res)-1]
	}

	return res
}

func parseInput(filepath string) ([]int, error) {
	nums := []int{}

	f, err := os.Open(filepath)
	if err != nil {
		return nums, err
	}
	defer f.Close()

	b := new(strings.Builder)
	io.Copy(b, f)
	re := regexp.MustCompile(`\d+`)
	strings := re.FindAllString(b.String(), -1)

	for _, s := range strings {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nums, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

func main() {
	nums, err := parseInput(utils.GetPuzzleInputSrc())
	utils.Check(err)

	program := partOne(nums)
	fmt.Println("Part one res", program)
}
