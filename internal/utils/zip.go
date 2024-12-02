package utils

import (
	"fmt"
)

// converts a list of tuples into separate lists
func Zip[T any](tuples [][]T) ([][]T, error) {
	if len(tuples) == 0 {
		return nil, fmt.Errorf("Can't perform Zip on empty list")
	}

	var res [][]T
	for i := 0; i < len(tuples[0]); i++ {
		res = append(res, make([]T, len(tuples)))
	}

	for i := 0; i < len(tuples); i++ {
		for j := 0; j < len(tuples[0]); j++ {
			res[j][i] = tuples[i][j]
		}
	}

	return res, nil
}
