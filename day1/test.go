package main

import (
	"fmt"
	utils "github.com/JohnBra/aoc-2024/utils"
	"unicode"
)

func main() {
	fmt.Println("Hello")
	contents, err := utils.GetFileContentsAsRunes("./input.txt")
	utils.Check(err)

	for r := 0; r < len(contents); r++ {
		line := contents[r]
		var val = ""
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(line[i]) {
				val += string(line[i])
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(line[i]) {
				val += string(line[i])
				break
			}
		}
		fmt.Println(val)
	}
}
