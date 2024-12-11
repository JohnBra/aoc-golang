package main

import (
	"container/list"
	"fmt"
	"math/big"
	"strings"

	"github.com/JohnBra/aoc-2024/internal/utils"
)

/*
If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.
*/

func times2024(s string) (string, error) {
	n := new(big.Int)
	n, ok := n.SetString(s, 10)
	if !ok {
		return "", fmt.Errorf("couldn't set string %s to big int", s)
	}

	return big.NewInt(0).Mul(n, big.NewInt(2024)).String(), nil
}

func blink(input []string) int {
	stones := list.New()
	stones.PushBack("head") // dummy element

	for _, s := range input {
		stones.PushBack(s)
	}

	stones.PushBack("tail") // dummy element

	for i := 0; i < 25; i++ {
		cur := stones.Front().Next()
		for cur != stones.Back() {
			val := cur.Value.(string)

			if val == "0" {
				cur.Value = "1"
			} else if len(val)%2 == 0 {
				first := strings.TrimLeft(val[:len(val)/2], "0")
				second := strings.TrimLeft(val[len(val)/2:], "0")

				if first == "" {
					first = "0"
				}

				if second == "" {
					second = "0"
				}
				stones.InsertBefore(first, cur)
				c := stones.InsertAfter(second, cur)
				stones.Remove(cur)
				cur = c
			} else {
				m, err := times2024(cur.Value.(string))
				utils.Check(err)
				cur.Value = m
			}
			cur = cur.Next()
		}
	}

	tmp := stones.Front().Next()
	fmt.Print("Stones ")
	for tmp != stones.Back() {
		fmt.Print(" ", tmp.Value, " ")
		tmp = tmp.Next()
	}
	fmt.Println("")

	return stones.Len() - 2
}

func main() {
	nums, err := utils.GetFileContentsAsNumberMatrix(utils.GetPuzzleInputSrc())
	utils.Check(err)

	partOneRes := partOne(nums[0])
	fmt.Println("Number of stones (25 blinks)", partOneRes)
}
