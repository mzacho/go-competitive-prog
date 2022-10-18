package types

import "sort"

type Interval struct {
	Start     int
	Finish    int
	Id        int
	Partition int
}

type IntervalSlice []Interval

func (x *IntervalSlice) SortByFinishTimeAscending() {
	sort.SliceStable(*x, func(i, j int) bool {
		return (*x)[i].Finish < (*x)[j].Finish
	})
}

func (x *IntervalSlice) SortByFinishTimeDescending() {
	sort.SliceStable(*x, func(i, j int) bool {
		return (*x)[i].Finish > (*x)[j].Finish
	})
}

func (x *IntervalSlice) SortByIdAscending() {
	sort.SliceStable(*x, func(i, j int) bool {
		return (*x)[i].Id < (*x)[j].Id
	})
}

func (x *IntervalSlice) SortByIdDescending() {
	sort.SliceStable(*x, func(i, j int) bool {
		return (*x)[i].Id > (*x)[j].Id
	})
}

func (x *IntervalSlice) SortByPartitionAscending() {
	sort.SliceStable(*x, func(i, j int) bool {
		return (*x)[i].Partition < (*x)[j].Partition
	})
}

func (x *IntervalSlice) SortByPartitionDescending() {
	sort.SliceStable(*x, func(i, j int) bool {
		return (*x)[i].Partition > (*x)[j].Partition
	})
}
