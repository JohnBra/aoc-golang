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

func partOne(a, b, c int, pr []int) string {
	// nums: a: 0; b: 1; c: 2; program: 3...n-1
	res := []int{}

	i := 0
	for i+1 < len(pr) {
		switch pr[i] {
		case 0: // adv
			a = a >> getComboOperand(pr[i+1], a, b, c)
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
			res = append(res, getComboOperand(pr[i+1], a, b, c)%8)
		case 6: // bdv
			b = a >> getComboOperand(pr[i+1], a, b, c)
		case 7: // cdv
			c = a >> getComboOperand(pr[i+1], a, b, c)
		}
		i += 2
	}

	return utils.IntSliceToString(res, ",")
}

func partTwo(program []int, res int) int {
	/*
		Program: 2,4, 1,1, 7,5, 1,5, 4,5, 0,3, 5,5, 3,0

		b = a % 8
		b = b ^ 1
		c = a >> b
		b = b ^ 5
		b = b ^ c
		a = a >> 3
		print b % 8
		a != 0 -> jump 0
	*/
	if len(program) == 0 {
		return res
	}

	for n := range 8 {
		a := (res << 3) | n
		b := n ^ 1
		c := a >> b
		b = b ^ 5
		b = b ^ c

		if b%8 == program[len(program)-1] {
			prev := partTwo(program[:len(program)-1], a)
			if prev == -1 {
				continue
			}
			return prev
		}
	}

	return -1
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

	program := partOne(nums[0], nums[1], nums[2], nums[3:])
	fmt.Println("Part one res", program)

	fmt.Println("\nPART TWO SOLUTION ONLY WORKS FOR THIS PROGRAM:")
	fmt.Println("2,4,1,1,7,5,1,5,4,5,0,3,5,5,3,0")

	val := partTwo(nums[3:], 0)
	fmt.Println("\nPart two res", val)
}
