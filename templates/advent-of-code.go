package templates

import (
	"competitive_programming/util/reader"
	"fmt"
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

type firstLine []int

func Run_AoC_XX(reader reader.FileReader) {
	var tcs []testcase
	var fl firstLine

	// Some days the first line needs special attention ...
	if true {
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
	return 42
}