package utils

import (
	"fmt"
	"math"
	"sort"
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

func InsertItem[T comparable](array []T, value T, index int) []T {
	return append(array[:index], append([]T{value}, array[index:]...)...)
}

func RemoveItem[T comparable](array []T, index int) []T {
	return append(array[:index], array[index+1:]...)
}

func MoveItem[T comparable](array []T, srcIndex int, destIndex int) []T {
	value := array[srcIndex]
	return InsertItem(RemoveItem(array, srcIndex), value, destIndex)
}

func SortListByOrder[T comparable](lst []T, order []T) []T {
	orderIndex := make(map[T]int)
	for i, num := range order {
		orderIndex[num] = i
	}

	sort.Slice(lst, func(i, j int) bool {
		indexI, foundI := orderIndex[lst[i]]
		indexJ, foundJ := orderIndex[lst[j]]

		// If not found, treat it as infinity
		if !foundI {
			indexI = math.MaxInt
		}
		if !foundJ {
			indexJ = math.MaxInt
		}

		return indexI < indexJ
	})

	return lst
}

// converts a list of tuples into separate lists
func Zip[T any](tuples [][]T) ([][]T, error) {
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
