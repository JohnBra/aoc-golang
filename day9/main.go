package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func readFile(filepath string) []rune {
	file, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	b := new(strings.Builder)
	fi, _ := file.Stat()
	b.Grow(int(fi.Size()))
	io.Copy(b, file)

	return []rune(b.String())
}

// day9 part 1
// naive solution is to create empty slice of "file blocks"
// iterate through all runes and append elements to the new slice according
// to the alternating file block / empty space pattern in puzzle description
// then use two pointers on that new slice one at start (left) one at end (right)
// and move file ids from right pointer to empty spaces of left pointer
// increment left pointer and decrement right pointer accordingly
// iterate through file blocks slice, calculate checksum and return result
//
// # O(sum(n)) time and space complexity which could be quite big
//
// better:
// init int slice (blocks)
// two pointers one left, one right of input rune slice
//
// left pointer:
// fill int slice with file ids and create "virtual empty slots"
// if left pointer fulfils empty slot condition
//
//	if pointer idx % 2 == 0
//		fileblock with file id pointer idx / 2 of size int(rune - '0')
//	else
//		empty block of size int(rune - '0')
//
// if "virtual empty slots" count is 0 fill from left side
// else fill file blocks from right, subtracting 1 blocksize for each fileblock
// of right pointer
// iterate through resulting int slice, calculate checksum and return result
func partOne(runes []rune) int {
	res := 0
	l, r := 0, len(runes)-1
	blocks := []int{}
	emptyBlocks := 0

	for l < r {
		if emptyBlocks == 0 { // fill from left side

			if l%2 == 0 { // file block
				blocks = append(blocks, slices.Repeat([]int{l / 2}, int(runes[l]-'0'))...)
			} else {
				emptyBlocks += int(runes[l] - '0')
			}
			l += 1
		} else { // fill empty blocks with file blocks from right side
			// add file blocks to queue but only if not empty
			for emptyBlocks > 0 {
				blocksize := int(runes[r] - '0')
				if r%2 == 0 {
					if blocksize == 0 {
						r -= 1
					} else {
						blocks = append(blocks, r/2)
						emptyBlocks -= 1
						runes[r] = rune(blocksize - 1 + '0')
					}
				} else {
					r -= 1
				}
			}
		}
	}

	// append remaining right blocks if exist
	if int(runes[r]-'0') > 0 {
		blocks = append(blocks, slices.Repeat([]int{r / 2}, int(runes[r]-'0'))...)
	}

	for i, v := range blocks {
		res += i * v
	}

	return res
}

func main() {
	runes := readFile("9.in")

	partOneRes := partOne(runes)
	fmt.Println("Checksum", partOneRes)
}
