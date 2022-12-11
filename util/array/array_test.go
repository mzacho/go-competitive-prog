package array_test

import (
	"competitive_programming/util/array"
	"testing"
)

func TestMax(t *testing.T) {
	a := []int{-2, 4, 0, 7, -3}
	idx, elem := array.Max(a...)
	if idx != 3 || elem != 7 {
		t.Errorf("max(%v) = %v, %v; want 3, 7", a, idx, elem)
	}
}

func TestReverse(t *testing.T) {
	a := []int{-2, 4, 0, 7, -3}
	r := array.Reverse(a)
	if len(a) != len(r) {
		t.Errorf("different lengths")
	}
	a = []int{-2, 4, 0, 7}
	r = array.Reverse(a)
	if len(a) != len(r) {
		t.Errorf("different lengths")
	}
}
