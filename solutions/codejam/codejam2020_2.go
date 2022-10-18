package solutions

import (
	"competitive_programming/util/convert"
	"competitive_programming/util/reader"
	"fmt"
)

func Run_codejam2020_2(reader reader.FileReader) {
	t, _ := reader.NextInt()
	for i := 0; i < t; i++ {
		l1, _ := reader.NextInt()
		fmt.Printf("Case #%v: %v\n", i+1, SplitInteger(l1))
		// fmt.Printf("Case #%v:\n", i+1)
		// fmt.Println(MY_SOLUTION(l1))
	}
}

func MY_SOLUTION(l1 int) string {
	l := 0
	res := ""
	for _, char := range convert.IntToString(l1) {
		n := convert.UnsafeToIntBase10(string(char))
		if n == l {
			res += string(char)
		} else if n > l {
			for k := 0; k < n-l; k++ {
				res += "("
			}
			res += string(char)
			l = n
		} else {
			for k := 0; k < l-n; k++ {
				res += ")"
			}
			res += string(char)
			l = n
		}
	}
	for k := 0; k < l; k++ {
		res += ")"
		l -= 1
	}
	return res
}
