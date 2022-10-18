package templates

import (
	"competitive_programming/util/reader"
	"fmt"
)

func Run_codejamZZZZ_Z(reader reader.FileReader) {
	t, _ := reader.NextInt()
	for i := 0; i < t; i++ {
		l1, _ := reader.SplitNextIntLine(" ")
		l2, _ := reader.SplitNextIntLine(" ")
		l3, _ := reader.SplitNextIntLine(" ")
		fmt.Printf("Case #%v: %v\n", i+1, MY_SOLUTIONZ(l1, l2, l3))
		// fmt.Printf("Case #%v:\n", i+1)
		// fmt.Println(MY_SOLUTIONZ(l1, l2, l3))
	}
}

func MY_SOLUTIONZ(l1 []int, l2 []int, l3 []int) string {
	return ""
}
