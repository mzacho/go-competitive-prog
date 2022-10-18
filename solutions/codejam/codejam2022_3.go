package solutions

import (
	util "competitive_programming/util/reader"
	"fmt"
	"sort"
)

func Run_codejam2022_3(reader util.FileReader) {
	noOfTests, _ := reader.NextInt()
	for i := 0; i < noOfTests; i++ {
		n, _ := reader.NextInt()
		ds, _ := reader.SplitNextIntLine(" ")
		fmt.Printf("Case #%v: %v\n", i+1, findMaxDieLength(n, ds))
	}
}

func findMaxDieLength(n int, ds []int) int {
	sort.Ints(ds)
	best := 0
	for _, di := range ds {
		if di > best {
			best += 1
		}
	}
	return best
}
