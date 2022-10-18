package solutions

import (
	"competitive_programming/util/greedy"
	"competitive_programming/util/reader"
	"competitive_programming/util/types"
	"fmt"
)

func Run_codejam2020_3(reader reader.FileReader) {
	t, _ := reader.NextInt()
	for i := 0; i < t; i++ {
		n, _ := reader.NextInt()
		var input = make(types.IntervalSlice, n)
		for k := 0; k < n; k++ {
			line, _ := reader.SplitNextIntLine(" ")
			input[k] = types.Interval{
				Start:  line[0],
				Finish: line[1],
				Id:     k,
			}
		}
		fmt.Printf("Case #%v: %v\n", i+1, ScheduleActivities(input))
		// fmt.Printf("Case #%v:\n", i+1)
		// fmt.Println(MY_SOLUTIONY(l1, l2))
	}
}

func ScheduleActivities(i types.IntervalSlice) (res string) {
	i.SortByFinishTimeAscending()
	greedy.PartitionActivities(2, i)
	i.SortByIdAscending()
	for n := 0; n < len(i); n++ {
		if i[n].Partition == 1 {
			res += "C"
		} else if i[n].Partition == 2 {
			res += "J"
		} else {
			res = "IMPOSSIBLE"
			return
		}
	}
	return
}
