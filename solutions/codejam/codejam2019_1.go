package solutions

import (
	"competitive_programming/util/convert"
	"competitive_programming/util/reader"
	"fmt"
	"math"
)

func Run_codejam2019_1(reader reader.FileReader) {
	t, _ := reader.NextInt()
	for i := 0; i < t; i++ {
		l1, _ := reader.NextInt()
		fmt.Printf("Case #%v: %v\n", i+1, SplitInteger(l1))
		// fmt.Printf("Case #%v:\n", i+1)
		// fmt.Println(MY_SOLUTION(l1))
	}
}

func SplitInteger(l1 int) string {
	var A float64
	var B float64
	var s = convert.IntToString(l1)
	for idx, c := range convert.IntToString(l1) {
		i := convert.UnsafeToIntBase10(string(c))
		if i == 4 {
			A += 2 * math.Pow(10, float64(len(s)-idx-1))
			B += 2 * math.Pow(10, float64(len(s)-idx-1))
		} else {
			A += float64(i) * math.Pow(10, float64(len(s)-idx-1))
		}
	}
	return fmt.Sprintf("%v %v", A, B)
}
