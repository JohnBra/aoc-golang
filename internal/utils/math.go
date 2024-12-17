package utils

import "math"

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Pow(base, exponent int) int {
	return int(math.Pow(float64(base), float64(exponent)))
}

func Min(a, b int, vals ...int) int {
	res := a
	if b < a {
		res = b
	}

	for _, v := range vals {
		if v < res {
			res = v
		}
	}

	return res
}

func Max(a, b int, vals ...int) int {
	res := a
	if b > a {
		res = b
	}

	for _, v := range vals {
		if v > res {
			res = v
		}
	}

	return res
}
