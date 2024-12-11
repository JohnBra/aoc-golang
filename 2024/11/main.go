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

func split(s string) (string, string) {
	first := s[:len(s)/2]
	second := strings.TrimLeft(s[len(s)/2:], "0")

	if second == "" {
		second = "0"
	}
	return first, second
}

// day11 part 1
// create doubly linked list
// add elements for every blink according to rules
// return length of list
func partOne(input []string) int {
	stones := list.New()

	for _, s := range input {
		stones.PushBack(s)
	}

	stones.PushFront("head") // dummy element
	stones.PushBack("tail")  // dummy element

	for i := 0; i < 25; i++ {

		cur := stones.Front().Next()

		for cur != stones.Back() {
			val := cur.Value.(string)

			if val == "0" {
				cur.Value = "1"
			} else if len(val)%2 == 0 {
				first, second := split(val)
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
	stones.Remove(stones.Front())
	stones.Remove(stones.Back())

	return stones.Len()
}

// day11 part 2
// create map for stones (stone -> count of stone)
// memoize splits
// memoize conversions from 0 -> 1 and num -> num * 2024
// iterate through stones map for each blink and calc
// new values
// return sum of final stones values
func partTwo(input []string) int {
	stones := map[string]int{}
	splits := map[string][2]string{}    // memo for number splits into two
	conv := map[string]string{"0": "1"} // memo for calculations

	for _, s := range input {
		stones[s] += 1
	}

	for i := 0; i < 75; i++ {
		nStones := map[string]int{}
		for stone, count := range stones {
			if stone != "0" && len(stone)%2 == 0 {
				var first, second string
				s, ok := splits[stone]
				if ok {
					first, second = s[0], s[1]
				} else {
					first, second = split(stone)
					splits[stone] = [2]string{first, second}
				}

				nStones[first] += count
				nStones[second] += count
			} else {
				s, ok := conv[stone]
				if !ok {
					m, err := times2024(stone)
					utils.Check(err)
					conv[stone] = m
					s = m
				}
				nStones[s] += count
			}
		}
		stones = nStones
	}

	res := 0
	for _, count := range stones {
		res += count
	}

	return res
}

func main() {
	nums, err := utils.GetFileContentsAsNumberMatrix(utils.GetPuzzleInputSrc())
	utils.Check(err)

	partOneRes := partOne(nums[0])
	fmt.Println("Number of stones (25 blinks)", partOneRes)

	partTwoRes := partTwo(nums[0])
	fmt.Println("Number of stones (75 blinks)", partTwoRes)
}
