package templates

import (
	"competitive_programming/util/reader"
	"fmt"
)

func Run_codejamXXXX_X(reader reader.FileReader) {
	t, _ := reader.NextInt()
	for i := 0; i < t; i++ {
		l1, _ := reader.SplitNextIntLine(" ")
		fmt.Printf("Case #%v: %v\n", i+1, MY_SOLUTION(l1))
		// fmt.Printf("Case #%v:\n", i+1)
		// fmt.Println(MY_SOLUTION(l1))
	}
}

func MY_SOLUTION(l1 []int) string {
	return ""
}
