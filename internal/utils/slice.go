package utils

import (
	"math"
	"sort"
)

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
