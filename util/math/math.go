package math

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
