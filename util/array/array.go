package array

import (
	"competitive_programming/util/functional"
	"competitive_programming/util/types"
)

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

func Sum(a []int) int {
	return functional.FoldL(a, 0, functional.PlusInt)
}

func Contains[T comparable](a []T, t T) bool {
	return IndexOf(t, a) != -1
}

func ContainsAny[T comparable](a []T, ts ...T) bool {
	for i := 0; i < len(ts); i++ {
		if IndexOf(ts[i], a) != -1 {
			return true
		}
	}
	return false
}

func ContainsAll[T comparable](a []T, ts ...T) bool {
	for i := 0; i < len(ts); i++ {
		if IndexOf(ts[i], a) == -1 {
			return false
		}
	}
	return true
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

func Fill[T any](a []T, t T) {
	for i := 0; i < len(a); i++ {
		a[i] = t
	}
}

func FillFun[T any](a []T, f func (int) T) {
	for i := 0; i < len(a); i++ {
		a[i] = f(i)
	}
}

func Make[T any](t T, n int) []T {
	a := make([]T, n)
	Fill(a, t)
	return a
}

func MakeFun[T any](f func (int) T, n int) []T {
	a := make([]T, n)
	FillFun(a, f)
	return a
}