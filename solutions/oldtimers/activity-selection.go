package solutions

import (
	"competitive_programming/util/greedy"
	"competitive_programming/util/reader"
	"competitive_programming/util/types"
	"fmt"
)

func Run_activity_selection(reader reader.FileReader) {
	t, _ := reader.NextInt()
	for i := 0; i < t; i++ {
		n, _ := reader.NextInt()
		var input = make(types.IntervalSlice, n)
		for k := 0; k < n; k++ {
			line, _ := reader.SplitNextIntLine(" ")
			input[k] = types.Interval{
				Start:  line[0],
				Finish: line[1],
			}
		}
		fmt.Printf("Case #%v: %v\n", i+1, SelectMaxNoOfActivities(input))
		// fmt.Printf("Case #%v:\n", i+1)
		// fmt.Println(MY_SOLUTIONY(l1, l2))
	}
}

func SelectMaxNoOfActivities(i types.IntervalSlice) int {
	greedy.ActivityPartition(1, i)
	var res types.IntervalSlice
	for k := 0; k < len(i); k++ {
		if i[k].Partition == 1 {
			res = append(res, i[k])
		}
	}

	return len(res)
}
