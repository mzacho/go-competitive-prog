package functional

import "strings"

func Map[T, R any](a []T, t func(T) R) []R {
	var r []R
	for _, e := range a {
		r = append(r, t(e))
	}
	return r
}

func FoldL[T, R any](xs []T, y R, f func (T, R) R) R {
	for i := 0; i < len(xs); i++ {
		y = f(xs[i], y)
	}
	return y
}

func FoldR[T, R any](xs []T, y R, f func (T, R) R) R {
	for i := len(xs)-1; i >= 0; i-- {
		y = f(xs[i], y)
	}
	return y
}

func Curry[T, S, R any](f func(T, S) R) func(T) func(S) R {
	return func(t T) func(S) R {
		return func(s S) R {
			return f(t, s)
		}
	}
}

func Uncurry[T, S, R any](f func(T) func(S) R) func(T, S) R {
	return func(t T, s S) R {
		return f(t)(s)
	}
}

func StrSplit(sep string) func(string) []string {
	return func(s string) []string {
		return strings.Split(s, sep)
	}
}

func PlusInt(x, y int) int {
	return x + y
}