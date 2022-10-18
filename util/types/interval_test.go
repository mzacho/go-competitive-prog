package types_test

import (
	"competitive_programming/util/types"
	"testing"
)

func TestSortIntervals(t *testing.T) {
	i1 := types.Interval{
		Start: 360,
		Finish: 480,
	}
	i2 := types.Interval{
		Start: 420,
		Finish: 640,
	}
	i3 := types.Interval{
		Start: 390,
		Finish: 800,
	}
	i := types.IntervalSlice{i2, i1, i3}
	i.SortByFinishTimeAscending()
	if i[0] != i1 {
		t.Errorf("expected i1, i2, i3; was %v", i)
	}
}

func TestPtrSlice(t *testing.T) {
	a := []int{42}
	foo(a)
	if a[0] == 43 {
		// slices are passed by reference
		return
	} else {
		t.Errorf("impossible")
	}

}

func foo(a []int) {
	a[0] = 43
}