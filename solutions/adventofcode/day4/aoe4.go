package day4

import (
	"competitive_programming/util/convert"
	"competitive_programming/util/functional"
	"competitive_programming/util/reader"
	"competitive_programming/util/types"
	"fmt"
	"strings"
)

// type line string
// type line []int
type line [][]int

func parseLine(s string) line {
	//return line(s)
	// return functional.Map(
	// 	strings.Split(s, " "),
	// 	convert.UnsafeToIntBase10,
	// )
	return functional.Map(
		functional.Map(
			strings.Split(s, ","),
			functional.StrSplit("-"),
		),
		func(s []string) []int {
			return functional.Map(
				s,
				convert.UnsafeToIntBase10,
			)
		},
	)
}

type testcase struct {
	lines []line
}

func parseTestCase(lines []line) testcase {
	return testcase{lines}
}

type firstLine []int

func Run_AoC(reader reader.FileReader) {
	var tcs []testcase
	var fl firstLine

	// Some days the first line needs special attention ...
	if false {
		// s, _ := reader.NextLineSplit(",")
		s, _ := reader.SplitNextIntLine(",")
		fl = firstLine(s)
	}

	for {
		s, err := reader.NextLine()
		if err != nil {
			break
		}
		var t testcase
		var lines []line
		for s != "" && err == nil {
			lines = append(lines, parseLine(s))
			s, err = reader.NextLine()
		}
		if len(lines) > 0 {
			t = parseTestCase(lines)
			tcs = append(tcs, t)
		}
	}
	fmt.Println(SOLVE(fl, tcs))
}

func SOLVE(fl firstLine, tcs []testcase) int {
	sum := 0
	for _, l := range tcs[0].lines {
		i1 := types.Interval{
			Start:  l[0][0],
			Finish: l[0][1],
		}
		i2 := types.Interval{
			Start:  l[1][0],
			Finish: l[1][1],
		}
		if i1.Overlaps(i2) || i2.Overlaps(i1) {
			sum += 1
		}
	}
	return sum
}
