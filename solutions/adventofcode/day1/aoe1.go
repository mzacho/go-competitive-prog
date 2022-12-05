package day1

import (
	"competitive_programming/util/reader"
	"fmt"
	"sort"
	"strconv"
)

func Run_AoC_1(reader reader.FileReader) {
	var lines [][]int
	lines = append(lines, nil)
	for {
		s, err := reader.NextLine()
		if err != nil {
			break
		}
		if s == "" {
			lines = append(lines, nil)
		} else {
			i, _ := strconv.ParseInt(s, 10, 0)
			lines[len(lines)-1] = append(lines[len(lines)-1], int(i))
		}
	}
	fmt.Println(max_sum(lines))
}

func max_sum(input [][]int) int {
	var sums sort.IntSlice
	for _, l := range input {
		sum := 0
		for _, i := range l {
			sum += i
		}
		sums = append(sums, sum)
	}
	sums.Sort()
	return sums[len(sums)-1] + sums[len(sums)-2] + sums[len(sums)-3]
}
