package utils

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Stack struct {
	items []int
}

func (s *Stack) Clear() {
	s.items = nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Print() {
	for _, item := range s.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("Stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetFileContentsAsRunes(filepath string) ([][]rune, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var res [][]rune

	for scanner.Scan() {
		var line = []rune(scanner.Text())
		res = append(res, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func GetFileContentsAsInts(filepath string) ([][]int, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	re := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(file)
	var res [][]int

	for scanner.Scan() {
		var list = []int{}
		nums := re.FindAllString(scanner.Text(), -1)
		for _, i := range nums {
			val, err := strconv.Atoi(i)
			if err != nil {
				return nil, err
			}
			list = append(list, val)
		}
		res = append(res, list)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
