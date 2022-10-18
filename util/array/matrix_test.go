package array_test

import (
	"competitive_programming/util/array"
	"testing"
)

func TestDim(t *testing.T) {
	l1 := []int{1, 2, 3, 4}
	l2 := []int{5, 6, 7, 8}
	M := array.Matrix[int]([][]int{l1, l2})
	n, m := M.Dim()
	if n != 2 || m != 4 {
		t.Errorf("M.Dim() = %v, %v; expected 2, 4", n, m)
	}
}

func TestTranspose(t *testing.T) {
	l1 := []int{1, 2, 3, 4}
	l2 := []int{5, 6, 7, 8}
	M := array.Matrix[int]([][]int{l1, l2})
	M.Transpose()
	n, m := M.Dim()
	if n != 4 || m != 2 {
		t.Errorf("M.Transpose().Dim() = %v, %v; expected 4, 2", n, m)
	}
}
