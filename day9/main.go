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
//
// approach 1:
// create empty slice of "file blocks"
// iterate through all runes and append elements to the new slice according
// to the alternating file block / empty space pattern in puzzle description
// then use two pointers on that new slice one at start (left) one at end (right)
// and move file ids from right pointer to empty spaces of left pointer
// increment left pointer and decrement right pointer accordingly
// iterate through file blocks slice, calculate checksum and return result
//
// # O(sum(n)) time and space complexity which could be quite big
//
// approach 2:
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

// day8 part 2
// create slice blocks of all file and empty blocks
// iterate through input from right side until all file blocks tested
// for every fileblock: search for window with empty blocks of fileblock size
// if found -> move there
// calculate checksum like in part one and return result
func partTwo(runes []rune) int {
	res := 0
	blocks := []int{}

	// first create blocks including empty spaces
	for i := 0; i < len(runes); i++ {
		if i%2 == 0 {
			blocks = append(blocks, slices.Repeat([]int{i / 2}, int(runes[i]-'0'))...)
		} else {
			blocks = append(blocks, slices.Repeat([]int{0}, int(runes[i]-'0'))...)
		}
	}

	// index for fileblock to move
	r := len(runes) - 1
	// lower and upper bound for search window in blocks
	lower, upper := int(runes[0]-'0'), len(blocks)

	for r > 0 {
		blocksize := int(runes[r] - '0')
		upper -= blocksize
		if r%2 == 0 {
			emptyBlocks := 0 // empty blocks in window
			left, right := lower, lower

			// search for window where empty blocks == blocksize
			for i := lower; i < upper; i++ {
				if blocks[right] == 0 {
					emptyBlocks += 1
				}

				// found window, move file there
				if blocksize == emptyBlocks {
					for j := 0; j < blocksize; j++ {
						blocks[left+j] = blocks[upper+j]
						blocks[upper+j] = 0
					}
					break
				}

				if right-left+1 >= blocksize {
					if blocks[left] == 0 {
						emptyBlocks -= 1
					}
					left += 1
				}
				right += 1
			}
		}

		r -= 1
	}

	for i, v := range blocks {
		res += i * v
	}

	return res
}

func main() {
	p1input := readFile("9.in")
	p2input := make([]rune, len(p1input))
	_ = copy(p2input, p1input)

	partOneRes := partOne(p1input)
	fmt.Println("Checksum part one:", partOneRes)

	partTwoRes := partTwo(p2input)
	fmt.Println("Checksum part two:", partTwoRes)
}
