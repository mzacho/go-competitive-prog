package solutions

import (
	"competitive_programming/util/convert"
	"competitive_programming/util/math"
	"competitive_programming/util/reader"
	"fmt"
	"strings"
)

func Run_codejam2017_2(reader reader.FileReader) {
	t, _ := reader.NextInt()
	for i := 0; i < t; i++ {
		l1, _ := reader.NextInt()
		fmt.Printf("Case #%v: %v\n", i+1, NonDecN(l1))
		// fmt.Printf("Case #%v:\n", i+1)
		// fmt.Println(MY_SOLUTION(l1))
	}
}

func NonDecN(n int) int {
	ns := strings.Split(convert.IntToString(n), "")
	var res = make([]int, len(ns)+1)
	for idx, ni := range ns {
		res[idx+1] = convert.UnsafeToIntBase10(ni)
	}
	res[0] = 0
	for i := 0; i < len(ns); {
		if res[i+1] >= res[i] {
			i++
		} else {
			res[i] -= 1
			for k := i+1; k < len(res); k++ {
				res[k] = 9
			}
			i = 0
		}
	}
	return math.Value(res[1:], 10)
}
