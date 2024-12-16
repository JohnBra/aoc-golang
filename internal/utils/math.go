package utils

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
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
