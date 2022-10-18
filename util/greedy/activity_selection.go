package greedy

import (
	"competitive_programming/util/types"
)

// Partitions activities into max size subset of mutually compatible activities
func ActivityPartition(i int, a types.IntervalSlice) {
	if len(a) == 0 {
		return
	}

	var k int
	for k = 0; k < len(a); k++ {
		if a[k].Partition == 0 {
			a[k].Partition = i
			break
		}
	}
	for m := 1; m < len(a); m++ {
		if a[m].Partition == 0 && a[m].Start >= a[k].Finish {
			a[m].Partition = i
			k = m
		}
	}
}

// Selects max size subset of mutually compatible activities with max overlap
// Assumes intervals are sorted on finish time ascending
func PartitionActivities(n int, a types.IntervalSlice) {
	for i := 0; i < n; i++ {
		ActivityPartition(i+1, a)
	}
}

func GetPartition(n int, i types.IntervalSlice) (res types.IntervalSlice) {
	for k := 0; k < len(i); k++ {
		if i[k].Partition == n {
			res = append(res, i[k])
		}
	}
	return
}

// // Partitions activities into max size subset of mutually compatible activities
// // Assumes intervals are
// // Precondition: 0 < idx[i] < len(a) for 0 < i < len(idx)
// func ActivityPartition(idx []int, a types.IntervalSlice) (cidx []int, icidx []int) {
// 	if len(idx) == 0 {
// 		return nil, nil
// 	}
// 	cidx = []int{idx[0]}
// 	// var incompatible types.IntervalSlice
// 	k := 0
// 	for m := 1; m < len(idx); m++ {
// 		if a[idx[m]].Start >= a[idx[k]].Finish {
// 			cidx = append(cidx, idx[m])
// 			k = m
// 		} else {
// 			icidx = append(icidx, idx[m])
// 		}
// 	}
// 	return
// }

// // Selects max size subset of mutually compatible activities with max overlap
// // Assumes intervals are sorted on finish time ascending
// func PartitionActivities(n int, a types.IntervalSlice) (res [][]int) {
// 	res = make([][]int, n)
// 	idx := array.OfNats(len(a))
// 	for i := 0; i < n; i++ {
// 		cidx, icidx := ActivityPartition(idx, a)
// 		res[i] = cidx
// 		idx = icidx
// 	}
// 	return
// }
