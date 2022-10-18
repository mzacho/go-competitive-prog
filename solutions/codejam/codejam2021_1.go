package solutions

import (
	"competitive_programming/util/array"
	"competitive_programming/util/reader"
	"fmt"
)

func Run_codejam2021_1(reader reader.FileReader) {
	t, _ := reader.NextInt()
	for i := 0; i < t; i++ {
		l1, _ := reader.NextInt()
		l2, _ := reader.SplitNextIntLine(" ")
		fmt.Printf("Case #%v: %v\n", i+1, reversort(l1, l2))
	}
}

func reversort(n int, l2 []int) int {
	if len(l2) == 0 {
		return 0
	}
	idx, _ := array.Min(l2...)
	res := array.Reverse(l2[:idx+1])

	return idx + 1 + reversort(n, append(res[1:], l2[idx+1:]...))
}
