package utils

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func GetFileContentsAsString(filepath string) (string, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var res string

	for scanner.Scan() {
		res += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return res, nil
}

func GetFileContentsAsRuneMatrix(filepath string) ([][]rune, error) {
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

func GetFileContentsAsIntMatrix(filepath string) ([][]int, error) {
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
