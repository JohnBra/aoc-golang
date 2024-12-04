package utils

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func GetFileContentsAsStringMatrix(filepath string, sep string) ([][]string, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var res [][]string

	for scanner.Scan() {
		var line = strings.Split(scanner.Text(), sep)
		res = append(res, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
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
		strings := re.FindAllString(scanner.Text(), -1)
		nums := make([]int, 0, len(strings))
		for _, v := range strings {
			num, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			nums = append(nums, num)
		}
		res = append(res, nums)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
