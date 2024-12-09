package utils

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	"golang.org/x/exp/constraints"
)

/* Personal notes:

Interesting read on bufio.Scanner vs os.ReadFile:
https://medium.com/golicious/comparing-ioutil-readfile-and-bufio-scanner-ddd8d6f18463

Gist: bufio.Scanner is less performant than os.ReadFile but can read line by line

To read a whole file it can be better to use os.ReadFile like in this SO answer:
https://stackoverflow.com/a/66804541
*/

// Types that can be added to each other respectively
type accumulable interface {
	constraints.Integer | constraints.Float | ~string
}

// Parses file line by line with parse func
// and adds result of type T to accumulative result of type T
//
// NOTE: returns error with lines longer than 65536 characters.
// configure buffer size if necessary: https://pkg.go.dev/bufio#Scanner.Buffer
func AccumulateLineResultFromFile[T accumulable](
	filepath string,
	parse func(line string) (T, error),
) (T, error) {
	var res T

	file, err := os.Open(filepath)

	if err != nil {
		return res, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineRes, err := parse(scanner.Text())

		if err != nil {
			return res, err
		}

		res += lineRes
	}

	if err := scanner.Err(); err != nil {
		return res, err
	}

	return res, nil
}

// Parses file line by line with parse func
// and appends parse result slice of type T to result slice
//
// NOTE: returns error with lines longer than 65536 characters.
// configure buffer size if necessary: https://pkg.go.dev/bufio#Scanner.Buffer
func GetSliceOfSlicesFromFile[T any](
	filepath string,
	parse func(line string) (T, error),
) ([]T, error) {
	var res []T

	file, err := os.Open(filepath)

	if err != nil {
		return res, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parsed, err := parse(scanner.Text())

		if err != nil {
			return res, err
		}

		res = append(res, parsed)
	}

	if err := scanner.Err(); err != nil {
		return res, err
	}

	return res, nil
}

func GetFileContentsAsIntMatrix(filepath string) ([][]int, error) {
	re := regexp.MustCompile(`\d+`)

	return GetSliceOfSlicesFromFile(filepath, func(line string) ([]int, error) {
		strings := re.FindAllString(line, -1)
		nums := make([]int, 0, len(strings))
		for _, v := range strings {
			num, err := strconv.Atoi(v)
			if err != nil {
				return nums, err
			}
			nums = append(nums, num)
		}
		return nums, nil
	})
}

// Returns slice of slices of positive integers but as strings
func GetFileContentsAsNumberMatrix(filepath string) ([][]string, error) {
	re := regexp.MustCompile(`\d+`)

	return GetSliceOfSlicesFromFile(filepath, func(line string) ([]string, error) {
		return re.FindAllString(line, -1), nil
	})
}
