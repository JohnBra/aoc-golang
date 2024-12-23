package utils

import (
	"fmt"
	"math"
)

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Pow(base, exponent int) int {
	return int(math.Pow(float64(base), float64(exponent)))
}

func Min(vals ...int) int {
	if len(vals) < 2 {
		panic(fmt.Errorf("need at least two values"))
	}

	res := vals[0]
	for _, v := range vals {
		if v < res {
			res = v
		}
	}

	return res
}

func Max(vals ...int) int {
	if len(vals) < 2 {
		panic(fmt.Errorf("need at least two values"))
	}

	res := vals[0]
	for _, v := range vals {
		if v > res {
			res = v
		}
	}

	return res
}
