package array

import "competitive_programming/util/types"

func Max(a ...int) (idx int, elem int) {
	for i := 0; i < len(a); i++ {
		if i == 0 || a[i] > elem {
			elem = a[i]
			idx = i
		}
	}
	return
}

func Min(a ...int) (idx int, elem int) {
	for i := 0; i < len(a); i++ {
		if i == 0 || a[i] < elem {
			elem = a[i]
			idx = i
		}
	}
	return
}

func IndexOf[T comparable](t T, l []T) int {
	for i := 0; i < len(l); i++ {
		if l[i] == t {
			return i
		}
	}
	return -1
}

func Reverse[T any](l []T) []T {
	var res = make([]T, len(l))
	for idx, e := range l {
		res[len(res)-(idx+1)] = e
	}
	return res
}

func OfNats(n int) (res []int) {
	for i := 0; i < n; i++ {
		res = append(res, i)
	}
	return
}

func Zip[T, R any](a []T, b []R) []types.Pair[T, R] {
	if len(a) != len(b) {
		panic("zip: slices have different length")
	}
	var res = make([]types.Pair[T, R], len(a))
	for i := 0; i < len(a); i++ {
		res = append(res, types.Pair[T, R]{
			N1: a[i],
			N2: b[i],
		})
	}
	return res
}