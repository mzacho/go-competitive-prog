package templates

import (
	"competitive_programming/util/reader"
	"fmt"
)

func Run_codejamXXXX_XX(reader reader.FileReader) {
	t, _ := reader.NextInt()
	for i := 0; i < t; i++ {
		n, _ := reader.NextInt()
		var input = make([][]int, n)
		for k := 0; k < n; k++ {
			line, _ := reader.SplitNextIntLine(" ")
			input = append(input, line)
		}
		fmt.Printf("Case #%v: %v\n", i+1, MY_SOLUTION_VARXXX(input))
		// fmt.Printf("Case #%v:\n", i+1)
		// fmt.Println(MY_SOLUTIONY(l1, l2))
	}
}

func MY_SOLUTION_VARXXX(input [][]int) string {
	return ""
}