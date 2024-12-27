package utils

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// shitty O(n) get index function
func IndexOf[T comparable](array []T, value T) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	return -1
}

func InsertItem[T any](array []T, value T, index int) []T {
	return append(array[:index], append([]T{value}, array[index:]...)...)
}

func RemoveItem[T any](array []T, index int) []T {
	return append(array[:index], array[index+1:]...)
}

func MoveItem[T any](array []T, srcIndex int, destIndex int) []T {
	value := array[srcIndex]
	return InsertItem(RemoveItem(array, srcIndex), value, destIndex)
}

// Returns the provided vertices in topological order
//
// Preserves the order of provided vertices if in degrees
// are equal (Stable Sort)
func SortListByOrder[T comparable](slice []T, order []T) []T {
	orderIndex := make(map[T]int)
	for i, num := range order {
		orderIndex[num] = i
	}

	/*
		TODO Optimize with this:

		slices.SortFunc(slice, func(a, b int) int {
			if n := strings.Compare(a.Name, b.Name); n != 0 {
				return n
			}
			// If names are equal, order by age
			return cmp.Compare(a.Age, b.Age)
		})
	*/

	sort.Slice(slice, func(i, j int) bool {
		indexI, foundI := orderIndex[slice[i]]
		indexJ, foundJ := orderIndex[slice[j]]

		// If not found, treat it as infinity
		if !foundI {
			indexI = math.MaxInt
		}
		if !foundJ {
			indexJ = math.MaxInt
		}

		return indexI < indexJ
	})

	return slice
}

// converts a list of tuples into separate lists
func ZipSplit[T any](tuples [][]T) ([][]T, error) {
	if len(tuples) == 0 {
		return nil, fmt.Errorf("can't perform Zip on empty list")
	}

	var res [][]T
	for i := 0; i < len(tuples[0]); i++ {
		res = append(res, make([]T, len(tuples)))
	}

	for i := 0; i < len(tuples); i++ {
		if len(tuples[i]) != len(tuples[0]) {
			return nil, fmt.Errorf("tuple lengths not equal. Error on index %d", i)
		}
		for j := 0; j < len(tuples[i]); j++ {
			res[j][i] = tuples[i][j]
		}
	}

	return res, nil
}

// converts two lists into pairs of tuples
func ZipMerge[T any](a, b []T) [][2]T {
	shortest := Min(len(a), len(b))
	tuples := make([][2]T, shortest)

	for i := 0; i < shortest; i++ {
		tuples[i] = [2]T{a[i], b[i]}
	}

	return tuples
}

func IntSliceToString(nums []int, sep string) string {
	str := make([]string, len(nums))
	for i, n := range nums {
		str[i] = strconv.Itoa(n)
	}
	return strings.Join(str, sep)
}

func Filter[T any](slice []T, keep func(item T) bool) []T {
	n := 0
	for _, x := range slice {
		if keep(x) {
			slice[n] = x
			n++
		}
	}
	return slice[:n]
}

func CartesianProduct[T any](args ...[]T) [][]T {
	pools := args
	npools := len(pools)
	indices := make([]int, npools)
	result := make([]T, npools)

	for i := range result {
		if len(pools[i]) == 0 {
			return nil
		}
		result[i] = pools[i][0]
	}

	results := [][]T{result}

	for {
		i := npools - 1
		for ; i >= 0; i -= 1 {
			pool := pools[i]
			indices[i] += 1

			if indices[i] == len(pool) {
				indices[i] = 0
				result[i] = pool[0]
			} else {
				result[i] = pool[indices[i]]
				break
			}

		}

		if i < 0 {
			return results
		}

		newresult := make([]T, npools)
		copy(newresult, result)
		results = append(results, newresult)
	}
}
