package day6

import (
	"competitive_programming/util/reader"
	"competitive_programming/util/types"
	"fmt"
	"strings"
)

type line string
// type line []int
// type line [][]string

func parseLine(s string) line {
	return line(s)
	// return functional.Map(
	// 	strings.Split(s, " "),
	// 	convert.UnsafeToIntBase10,
	// )
	// return functional.Map(
	// 	strings.Split(s, " -> "),
	// 	functional.StrSplit(","),
	// )
}

type testcase struct {
	lines []line
}

func parseTestCase(lines []line) testcase {
	return testcase{lines}
}

type firstLine []string

func Run_AoC(reader reader.FileReader) {
	var tcs []testcase
	var fl firstLine

	// Some days the first line needs special attention ...
	if true {
		// s, _ := reader.NextLineSplit(",")
		s, _ := reader.NextLine()
		fl = firstLine(strings.Split(s, ""))
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
	fmt.Println(SOLVE2(fl, tcs))
}

func SOLVE(fl firstLine, tcs []testcase) int {
	var a string = fl[0]
	var b string = fl[1]
	var c string = fl[2]
	for i, x := range fl[3:] {
		if len(types.NewSet(a, b, c, x))==4 {
			return i+4
		} else {
			a = b
			b = c
			c = x
		}
	}
	return -1
}

func SOLVE2(fl firstLine, tcs []testcase) int {
	k := 14
	for i := range fl[k-1:] {
		x := types.NewSet(fl[i:i+k]...)
		if len(x)==k {
			return i+k
		}
	}
	return -1
}