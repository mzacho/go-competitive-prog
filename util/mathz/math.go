package mathz

import (
	"competitive_programming/util/convert"
	"competitive_programming/util/functional"
	"math"
)

func Value(digits []int, base int) (res int) {
	res = 0
	for idx, d := range digits {
		res += d * int(math.Pow(float64(base), float64(len(digits)-idx-1)))
	}
	return
}

func StrValue(digits []string, base int) (res int) {
	return Value(functional.Map(digits, functional.Curry(convert.UnsafeToInt)(base)), base)
}

func AbsInt(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func MinInt(xs ...int) int {
	var min int
	set := false
	for _, x := range xs {
		if !set || min > x {
			min = x
			set = true
		}
	}
	return min
}

func MaxInt(xs ...int) int {
	var max int
	set := false
	for _, x := range xs {
		if !set || x > max {
			max = x
			set = true
		}
	}
	return max
}
