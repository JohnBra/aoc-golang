package datastructures

import "fmt"

type Map[T comparable, U any] map[T]U

// utility function to retrieve a value with an optional fallback
//
// can provide 0..1 fallback; panics on multiple fallbacks
func (m Map[T, U]) Get(key T, fallback ...U) U {
	val, ok := m[key]
	if !ok {
		if len(fallback) == 1 {
			return fallback[0]
		} else if len(fallback) > 1 {
			panic(fmt.Errorf("too many fallback values"))
		}
	}
	return val
}

func (m Map[T, U]) Keys() []T {
	keys := make([]T, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func (m Map[T, U]) Values() []U {
	values := make([]U, len(m))
	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}
