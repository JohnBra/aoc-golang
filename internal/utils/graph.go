package utils

var Verticals = [4][2]int{
	{-1, 0}, // ⬆️
	{1, 0},  // ⬇️
}

var Horizontals = [4][2]int{
	{0, 1},  // ➡️
	{0, -1}, // ⬅️
}

var Axes = [4][2]int{
	{-1, 0}, // ⬆️
	{0, 1},  // ➡️
	{1, 0},  // ⬇️
	{0, -1}, // ⬅️
}

var Diagonals = [4][2]int{
	{1, 1},   // ↘
	{1, -1},  // ↙️
	{-1, -1}, // ↖
	{-1, 1},  // ↗
}

// Order:
//
// 0: ⬆️
// 1: ↗
// 2: ➡️
// 3: ↘
// 4: ⬇️
// 5: ↙️
// 6: ⬅
// 7: ↖
var Directions = [8][2]int{
	{-1, 0},  // ⬆️
	{-1, 1},  // ↗
	{0, 1},   // ➡️
	{1, 1},   // ↘
	{1, 0},   // ⬇️
	{1, -1},  // ↙️
	{0, -1},  // ⬅️
	{-1, -1}, // ↖
}

func IsOutOfBounds[T any](matrix [][]T, r, c int) bool {
	return r < 0 || c < 0 || r == len(matrix) || c == len(matrix[0])
}
