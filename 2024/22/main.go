package main

import (
	"fmt"
	"strconv"

	ds "github.com/JohnBra/aoc-2024/internal/datastructures"
	"github.com/JohnBra/aoc-2024/internal/utils"
)

func mix(x, y int) int { return x ^ y }
func prune(x int) int  { return x & 0b111111111111111111111111 }

func solve(nums []int) (int, int) {
	const iter int = 2000
	p1, p2 := 0, 0

	buyers := map[[4]int]int{}
	for _, n := range nums {
		prices := make([]int, iter+1)
		prices[0] = n % 10
		seqs := map[[4]int]int{}

		for i := 1; i < iter+1; i++ {
			x := prune(mix(n<<6, n))
			x = prune(mix(x>>5, x))
			n = prune(mix(x<<11, x))
			prices[i] = n % 10
		}
		p1 += n

		seq := ds.NewDeque([]int{})

		for i := 1; i < len(prices); i++ {
			seq.PushBack(prices[i] - prices[i-1])

			if seq.Len() == 4 {
				key := [4]int(seq.Members())
				if _, ok := seqs[key]; !ok {
					seqs[key] = prices[i]
				}
				seq.PopFront()
			}
		}

		for seq, val := range seqs {
			buyers[seq] += val
		}
	}

	for _, val := range buyers {
		if val > p2 {
			p2 = val
		}
	}

	return p1, p2
}

func parseLine(line string) (int, error) {
	num, err := strconv.Atoi(line)
	if err != nil {
		return 0, nil
	}
	return num, nil
}

func main() {
	input, err := utils.GetSliceOfSlicesFromFile(utils.GetPuzzleInputSrc(), parseLine)
	utils.Check(err)

	partOneRes, partTwoRes := solve(input)
	fmt.Println("Part one res", partOneRes)
	fmt.Println("Part two res", partTwoRes)
}
