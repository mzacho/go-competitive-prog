package array

import (
	"fmt"
)

type Matrix[T any] [][]T

func (M *Matrix[T]) Dim() (int, int) {
	if M == nil || len(*M) == 0 {
		return 0, 0
	}
	return len(*M), len((*M)[0])
}

func Trace[T, S any](M Matrix[T], s func(S, T) S) S {
	n, m := M.Dim()
	if n != m {
		panic(fmt.Errorf("Matrix dim: %v, %v; expected N x N", n, m))
	}
	var sum S
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			sum = s(sum, M[i][j])
		}
	}
	return sum
}

func (M *Matrix[T]) Transpose() {
	n, m := M.Dim()
	S := make([][]T, m)
	for j := 0; j < m; j++ {
		S[j] = make([]T, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			S[j][i] = (*M)[i][j]
		}
	}
	*M = S
}