package templates

import (
	"competitive_programming/util/reader"
	"fmt"
)

func Run_codejamYYYY_Y(reader reader.FileReader) {
	t, _ := reader.NextInt()
	for i := 0; i < t; i++ {
		l1, _ := reader.SplitNextIntLine(" ")
		l2, _ := reader.SplitNextIntLine(" ")
		fmt.Printf("Case #%v: %v\n", i+1, MY_SOLUTIONY(l1, l2))
		// fmt.Printf("Case #%v:\n", i+1)
		// fmt.Println(MY_SOLUTIONY(l1, l2))
	}
}

func MY_SOLUTIONY(l1 []int, l2 []int) string {
	return ""
}
